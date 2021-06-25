package ssh

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	convert "github.com/nickmancari/gocean/data"
//	commands "github.com/nickmancari/gocean/cmd"
)

type SessionConfig struct {
	IP	string
	User	string

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

	fmt.Println("\nWhich Droplet Would You Like to Connect to?\n")
	var droplet string
	fmt.Scan(&droplet)

	ip, err := convert.ToIP(droplet)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nWhich User?\n")
	var user string
	fmt.Scan(&user)

	return &SessionConfig{IP: ip, User: user}

}

func (s SessionConfig) Run() {
	connect, err := exec.LookPath("ssh")
	if err != nil {
		panic(err)
	}
	// +build linux,386 darwin,!cgo
	syscall.Exec(connect, []string{"ssh", s.User+"@"+s.IP}, os.Environ())
}
