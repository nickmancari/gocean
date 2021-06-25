package main

import (
	"flag"
	"os"


	"github.com/nickmancari/gocean/cmd"
	token "github.com/nickmancari/gocean/env"
)

var tokenFlag = flag.String("token", "", "Creates Token File For DO API Connection\n'rm' Remove Token")
var createFlag = flag.String("create", "", "Create a Droplet")
var destroyFlag = flag.String("destroy", "", "Destroys Specified Droplet")
var getFlag = flag.String("droplet", "", "Get Info On Droplet\n'ls' Returns Info on All Droplets")
var connectFlag = flag.String("connect", "", "Initiate a SSH Connection to Droplet")
var actionFlag = flag.String("action", "", "Dictate Action to A Droplet")

//func init() {
//	flag.StringVar(actionFlag, "cmd", "", "Action to Command")
//}

func main() {
	flag.Parse()
	token.CreateTokenFile(*tokenFlag)
	cmd.CreateDroplet(*createFlag)
	cmd.DestroyDroplet(*destroyFlag)
	cmd.GetDroplet(*getFlag)
	cmd.Shell(*connectFlag)
	cmd.Action(*actionFlag, os.Args[3:])

}
