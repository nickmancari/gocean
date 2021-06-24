package ssh

import (
	"fmt"

	convert "github.com/nickmancari/gocean/data"
//	commands "github.com/nickmancari/gocean/cmd"
)

type SessionConfig struct {
	IP	string

}

func Session(d string) (interface{}, error) {

	ip, err := convert.ToIP(d)
	if err != nil {
		return fmt.Println(err)
	}

	r, err := fmt.Println(ip)
	if err != nil {
		return fmt.Println(err)
	}

	return r, nil

}

func Start() *SessionConfig {

//	list, err := commands.GetDroplet("ls")
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(list)
	fmt.Println("\nWhich Droplet Would You Like to Connect to?\n")
	var droplet string
	fmt.Scan(&droplet)

	ip, err := convert.ToIP(droplet)
	if err != nil {
		fmt.Println(err)
	}
	return &SessionConfig{IP: ip}

}

func Run() {

}
