package main

import (
	"flag"
	"fmt"

	commands "github.com/nickmancari/gocean/cmd"
	token "github.com/nickmancari/gocean/env"
)

var tokenFlag = flag.String("token", "", "Creates Token File For DO API Connection.")
var createFlag = flag.String("create", "", "Create a Droplet")

func init() {
	flag.StringVar(createFlag, "s", "", "Add")
}

func main() {
	flag.Parse()
	token.CreateTokenFile(*tokenFlag)
	r := commands.CreateDroplet(*createFlag)

	fmt.Println(r)
}
