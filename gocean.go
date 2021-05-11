package main

import (
	"flag"
	"os"


	commands "github.com/nickmancari/gocean/cmd"
	token "github.com/nickmancari/gocean/env"
)

var tokenFlag = flag.String("token", "", "Creates Token File For DO API Connection\n'rm' Remove Token")
var createFlag = flag.String("create", "", "Create a Droplet")
var destroyFlag = flag.String("destroy", "", "Destroys Specified Droplet")
var getFlag = flag.String("droplet", "", "Get Info On Droplet\n'ls' Returns Info on All Droplets")
var rebootFlag = flag.String("reboot", "", "Reboot Specific Droplet")
var connectFlag = flag.String("connect", "", "Initiate a SSH Connection to Droplet")
var shutdownFlag = flag.String("shutdown", "", "Shutdown a Droplet")
var bootFlag = flag.String("boot", "", "Boot Up a Droplet")
var actionFlag = flag.String("action", "", "Dictate Action to A Droplet")

//func init() {
//	flag.StringVar(actionFlag, "cmd", "", "Action to Command")
//}

func main() {
	flag.Parse()
	token.CreateTokenFile(*tokenFlag)
	commands.CreateDroplet(*createFlag)
	commands.DestroyDroplet(*destroyFlag)
	commands.GetDroplet(*getFlag)
	commands.RebootDroplet(*rebootFlag)
	commands.Shell(*connectFlag)
	commands.ShutdownDroplet(*shutdownFlag)
	commands.BootDroplet(*bootFlag)
	commands.Action(*actionFlag, os.Args[3:])

}
