package main

import (
	"fmt"
	"time"
)

func main() {

	timer := time.NewTimer(30 * time.Second)
	for {
		<-timer.C

		ipOfGithubCom, err := GetIpOfGithubCom()
		if err != nil {
			fmt.Printf(err.Error())
		}
		ipOfGithubSsl, err := GetIpOfGithubSsl()
		if err != nil {
			fmt.Printf(err.Error())
		}
		err = SetHostsLine(ipOfGithubSsl, ipOfGithubCom)
		if err != nil {
			fmt.Printf(err.Error())
		}

		timer.Reset(10 * time.Second)
	}

}
