package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	commands "github.com/nickmancari/gocean/cmd"
	token "github.com/nickmancari/gocean/env"
)

func main() {

	tokenFlag := flag.Bool("token", false, "Digital Ocean API Token")
//	createFlag := flag.String("create", "Droplet_Name", "Creates Droplet")
	//	destroyFlag := flag.String()

	flag.Parse()
	
	if *tokenFlag {
		token.CreateTokenFile("test123")
		readTokenFile, err := ioutil.ReadFile(".token")
		if err != nil {
			fmt.Println(err)
		}
		token := string(readTokenFile)

		commands.CreateDroplet("test-droplet", token)
		return
	}
	//	commands.DestroyDroplet(*destroyFlag)

	fmt.Println("test successful")

}
