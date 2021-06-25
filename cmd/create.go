package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	connect "github.com/nickmancari/gocean/api"
	convert "github.com/nickmancari/gocean/data"
	color "github.com/nickmancari/gocean/pkg"
)

type Create struct {
	Distro  string
	Version string
}

func DistroSelection() *Create {

	fmt.Println("\nWhich Distro Do You Want?\n")
	fmt.Println("[" + color.Red + "1" + color.Reset + "]Ubuntu  [" + color.Red + "2" + color.Reset + "]CentOS  [" + color.Red + "3" + color.Reset + "]Debian  [" + color.Red + "4" + color.Reset + "]FreeBSD  [" + color.Red + "5" + color.Reset + "]Fedora  [" + color.Red + "6" + color.Reset + "]RancherOS")
	var distro string
	fmt.Scan(&distro)

	if distro == "1" {
		return &Create{Distro: "Ubuntu"}
	} else if distro == "2" {
		return &Create{Distro: "CentOS"}
	} else if distro == "3" {
		return &Create{Distro: "Debian"}
	} else if distro == "4" {
		return &Create{Distro: "FreeBSD"}
	} else if distro == "5" {
		return &Create{Distro: "Fedora"}
	} else if distro == "6" {
		return &Create{Distro: "RancherOS"}
	}
	fmt.Println("Input Not Recognized")
	os.Exit(1)

	return &Create{}

}

func DistroVersion(s *Create) (*Create, error) {

	var version string

	if string(s.Distro) == "Ubuntu" {
		fmt.Println("\nWhich Version of Ubuntu?\n")
		fmt.Println("[" + color.Red + "1" + color.Reset + "]20-04x64  [" + color.Red + "2" + color.Reset + "]18-04x64  [" + color.Red + "3" + color.Reset + "]21-04x64  [" + color.Red + "4" + color.Reset + "]20-10x64")
		fmt.Scan(&version)
		if version == "1" {
			d := &Create{Version: "ubuntu-20-04-x64"}
			return d, nil
		} else if version == "2" {
			d := &Create{Version: "ubuntu-18-04-x64"}
			return d, nil
		} else if version == "3" {
			d := &Create{Version: "ubuntu-21-04x64"}
			return d, nil
		} else if version == "4" {
			d := &Create{Version: "ubuntu-20-10x64"}
			return d, nil
		} else {
			fmt.Println("Input Not Recognized")
		}

	} else if string(s.Distro) == "CentOS" {
		fmt.Println("\nWhich Version of CentOS?\n")
		fmt.Println("[" + color.Red + "1" + color.Reset + "]7-x64  [" + color.Red + "2" + color.Reset + "]8-x64")
		fmt.Scan(&version)
		if version == "1" {
			d := &Create{Version: "centos-7-x64"}
			return d, nil
		} else if version == "2" {
			d := &Create{Version: "centos-8-x64"}
			return d, nil
		} else {
			fmt.Println("Input Not Recognized")
		}

	} else if string(s.Distro) == "Debian" {
		fmt.Println("\nWhich Version of Debian?")
		fmt.Println("[" + color.Red + "1" + color.Reset + "]9-x64  [" + color.Red + "2" + color.Reset + "]10-x64")
		fmt.Scan(&version)
		if version == "1" {
			d := &Create{Version: "debian-9-x64"}
			return d, nil
		} else if version == "2" {
			d := &Create{Version: "debian-10-x64"}
			return d, nil
		} else {
			fmt.Println("Input Not Recognized")
		}
	} else if string(s.Distro) == "Fedora" {
		fmt.Println("\nWhich Version of Fedora?\n")
		fmt.Println("[" + color.Red + "1" + color.Reset + "]34-x64  [" + color.Red + "2" + color.Reset + "]33-x64")
		fmt.Scan(&version)
		if version == "1" {
			d := &Create{Version: "fedora-34-x64"}
			return d, nil
		} else if version == "2" {
			d := &Create{Version: "fedora-33-x64"}
			return d, nil
		} else {
			fmt.Println("Input Not Recognized")
		}
	} else if string(s.Distro) == "FreeBSD" {
		fmt.Println("\nWhich Verison of FreeBSD?\n")
		fmt.Println("[" + color.Red + "1" + color.Reset + "]11-x64-zfs  [" + color.Red + "2" + color.Reset + "]11-x64-ufs  [" + color.Red + "3" + color.Reset + "]12-x64-zfs [" + color.Red + "4" + color.Reset + "]12-x64-ufs")
		fmt.Scan(&version)
		if version == "1" {
			d := &Create{Version: "freebsd-11-x64-zfs"}
			return d, nil
		} else if version == "2" {
			d := &Create{Version: "freebsd-11-x64-ufs"}
			return d, nil
		} else if version == "3" {
			d := &Create{Version: "freebsd-12-x64-zfs"}
			return d, nil
		} else if version == "4" {
			d := &Create{Version: "freebsd-12-x64-ufs"}
			return d, nil
		} else {
			fmt.Println("Input Not Recognized")
		}
	} else if string(s.Distro) == "RancherOS" {
		d := &Create{Version: "rancheros"}
		return d, nil
	} else if string(s.Distro) == "" {
		return &Create{}, errors.New("Distro Not Selected")
	}
	return &Create{}, nil
}

func DropletCreation(s *Create) (interface{}, error) {
	var name string
	fmt.Println("\nWhat Would You like to Name This Droplet?\n")
	fmt.Scan(&name)

	return s.NewDroplet(name)
}

func (v *Create) NewDroplet(name string) (interface{}, error) {

	jsonData := struct {
		Name   string
		Region string
		Size   string
		Image  *string
	}{
		Name:   name,
		Region: "nyc3",
		Size:   "s-1vcpu-1gb",
		Image:  &v.Version,
	}
	marsh, err := json.Marshal(jsonData)
	if err != nil {
		return fmt.Println(err)
	}

	r := connect.ConvertConnection("POST", apiAddress, bytes.NewBuffer(marsh))
	c, _ := convert.StructuredResponse(r)

	return c, nil
}
