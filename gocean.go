package main

import (
	"fmt"
	"flag"
	"github.com/nickmancari/gocean/env"
	"github.com/nickmancari/gocean/cmd"
)

func main() {

	tokenFlag := flag.String("token", "", "Digital Ocean API Token")
	createFlag := flag.String("create", "Droplet_Name", "Creates Droplet")
//	destroyFlag := flag.String()
	
	flag.Parse()

	token.CreateTokenFile(*tokenFlag)

	commands.CreateDroplet(*createFlag)
//	commands.DestroyDroplet(*destroyFlag)
	

	fmt.Println("test successful")	

}
