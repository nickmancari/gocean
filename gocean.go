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
	
	flag.Parse()
	token.CreateTokenFile(*tokenFlag)
	create.CreateDroplet(*createFlag)
	

	fmt.Println("test successful")	

}
