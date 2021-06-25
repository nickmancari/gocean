package ssh

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	convert "github.com/nickmancari/gocean/data"
)

type SessionConfig struct {
	IP	string
	User	string

}

func Start() *SessionConfig {

	fmt.Println("\nWhich Droplet Would You Like to Connect to?\n")
	var droplet string
	fmt.Scan(&droplet)

	ip, err := convert.ToIP(droplet)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nUsername:\n")
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
