package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	token "github.com/nickmancari/gocean/env"
)

func main() {

	tokenFlag := flag.String("token", "", "Digital Ocean API Token")

	flag.Parse()

	if *tokenFlag == "" {
		fmt.Println("Missing API Token")
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

	fmt.Println("test successful")

}
