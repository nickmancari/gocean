package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	token "github.com/nickmancari/gocean/env"
)

func main() {

	tokenFlag := flag.String("token", "", "Digital Ocean API Token")
//	createFlag := flagString("create", "", "Creates Droplet")

	flag.Parse()

	if *tokenFlag == "" {
		fmt.Println("No Token")
		os.Exit(1)
	
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
	

}
