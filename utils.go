package main

import (
	"bufio"
	"os"
	"runtime"
)

const WinHostsFile = ""
const LinuxHostFile = "/home/wangcheng/桌面/hoststmp"

func GetHostsLine() ([]string, error) {
	var lines []string
	var hostFile *os.File
	defer func(hostFile *os.File) {
		_ = hostFile.Close()
	}(hostFile)
	if runtime.GOOS == "linux" {
		tmp, err := os.Open(LinuxHostFile)
		if err != nil {
			return nil, err
		}
		hostFile = tmp
	} else if runtime.GOOS == "windows" {
		tmp, err := os.Open(WinHostsFile)
		if err != nil {
			return nil, err
		}
		hostFile = tmp
	}
	scanner := bufio.NewScanner(hostFile)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}
