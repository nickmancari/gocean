package main

import (
	"flag"

	commands "github.com/nickmancari/gocean/cmd"
	token "github.com/nickmancari/gocean/env"
)

var tokenFlag = flag.String("token", "", "Creates Token File For DO API Connection")
var createFlag = flag.String("create", "", "Create a Droplet")
var destroyFlag = flag.String("destroy", "", "Destroys Specified Droplet")
var getFlag = flag.String("droplet", "", "Get Info On Droplet")
var rebootFlag = flag.String("reboot", "", "Reboot Specific Droplet")
var connectFlag = flag.String("connect", "", "Initiate a SSH Connection to Droplet")

func init() {
	flag.StringVar(createFlag, "s", "", "Add")
}

func main() {
	flag.Parse()
	token.CreateTokenFile(*tokenFlag)
	commands.CreateDroplet(*createFlag)
	commands.DestroyDroplet(*destroyFlag)
	commands.GetDroplet(*getFlag)
	commands.RebootDroplet(*rebootFlag)
	commands.Shell(*connectFlag)

}
