package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	commands "github.com/nickmancari/gocean/cmd"
	token "github.com/nickmancari/gocean/env"
)

func main() {

	tokenFlag := flag.String("token", "", "Digital Ocean API Token")
	createFlag := flag.String("create", "", "Creates Droplet")

	flag.Parse()

	if *tokenFlag == "" {
		return

	} else {
		token.CreateTokenFile(*tokenFlag)
		readTokenFile, err := ioutil.ReadFile(".token")
		if err != nil {
			fmt.Println(err)
		}
		token := string(readTokenFile)
		fmt.Println("Token Accepted: ", token)

		return
	}

	if *createFlag == "" {
		return
	} else {
		readToken, err := ioutil.ReadFile(".token")
		if err != nil {
			fmt.Println("Token File Not Found", err)
		}
		createToken := string(readToken)
		commands.CreateDroplet(*createFlag, createToken)
		fmt.Println("Droplet Created")

		return
	}

}
