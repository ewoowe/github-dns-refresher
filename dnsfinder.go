package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

const ipaddress = "https://ipaddress.com/website/"
const github_com = "github.com"
const github_ssl = "github.global.ssl.fastly.net"
const pattern = "IP Address</th><td><ul class=\"comma-separated\"><li>"

func httpGet(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func getIpOf(url string) (string, error) {
	rsp, err := httpGet(ipaddress + url)
	if err != nil {
		return "", err
	}
	index := strings.LastIndex(rsp, pattern)
	if index != -1 {
		i := len(pattern)
		var ip string
		for {
			if rsp[index+i] == '<' {
				break
			}
			ip += string(rsp[index+i])
			i++
		}
		return ip, nil
	}
	return "", errors.New("unknown error")
}

func GetIpOfGithubCom() (string, error) {
	return getIpOf(github_com)
}

func GetIpOfGithubSsl() (string, error) {
	return getIpOf(github_ssl)
}
