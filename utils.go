package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
)

const WinHostsFile = ""
const LinuxHostFile = "/home/wangcheng/桌面/hoststmp"

//SetHostsLine ssl is ip of "github.global.ssl.fastly.net", github is ip of "github.com"
func SetHostsLine(ssl, github string) error {
	var lines []string
	var hostFile *os.File
	if runtime.GOOS == "linux" {
		tmp, err := os.OpenFile(LinuxHostFile, os.O_APPEND, 666)
		if err != nil {
			return err
		}
		hostFile = tmp
	} else if runtime.GOOS == "windows" {
		tmp, err := os.OpenFile(WinHostsFile, os.O_APPEND, 666)
		if err != nil {
			return err
		}
		hostFile = tmp
	}
	scanner := bufio.NewScanner(hostFile)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	_ = hostFile.Close()
	// if hostFile will be changed, flag will be set to true
	flag := false
	githubExist := false
	sslExist := false

	for index, line := range lines {
		real := strings.TrimSpace(line)
		// ignore if is a comment line
		if strings.HasPrefix(line, "#") {
			continue
		}
		ipAndUrl := strings.Split(real, " ")
		// ignore if not a valid line
		if len(ipAndUrl) != 2 {
			continue
		}
		ip := strings.TrimSpace(ipAndUrl[0])
		url := strings.TrimSpace(ipAndUrl[1])
		if strings.EqualFold(url, GithubCom) {
			githubExist = true
			if !strings.EqualFold(ip, github) {
				fmt.Printf("%s changed, old is %s, new is %s\n", GithubCom, ip, github)
				lines[index] = github + " " + GithubCom
				flag = true
			}
		}
		if strings.EqualFold(url, GithubSsl) {
			sslExist = true
			if !strings.EqualFold(ip, ssl) {
				fmt.Printf("%s changed, old is %s, new is %s\n", GithubSsl, ip, ssl)

				lines[index] = ssl + " " + GithubSsl
				flag = true
			}
		}
	}
	if !githubExist {
		fmt.Printf("%s not exist, will be set to %s\n", GithubCom, github)
		lines = append(lines, github+" "+GithubCom)
		flag = true
	}
	if !sslExist {
		fmt.Printf("%s not exist, will be set to %s\n", GithubSsl, ssl)
		lines = append(lines, ssl+" "+GithubSsl)
		flag = true
	}

	if flag {
		if runtime.GOOS == "linux" {
			tmp, err := os.OpenFile(LinuxHostFile, os.O_APPEND|os.O_RDWR|os.O_TRUNC, os.ModeAppend)
			if err != nil {
				return err
			}
			hostFile = tmp
		} else if runtime.GOOS == "windows" {
			tmp, err := os.OpenFile(WinHostsFile, os.O_APPEND|os.O_RDWR|os.O_TRUNC, os.ModeAppend)
			if err != nil {
				return err
			}
			hostFile = tmp
		}

		for _, line := range lines {
			_, err := io.WriteString(hostFile, line+"\n")
			if err != nil {
				fmt.Printf(err.Error() + "\n")
			}
		}

		_ = hostFile.Close()

		source()
	} else {
		fmt.Printf("not any host changed, noting to do\n")
	}

	return nil
}

//source source host file
func source() {

}
