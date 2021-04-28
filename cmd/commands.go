package commands

import (
	"bytes"
	"encoding/json"
	"fmt"

	connect "github.com/nickmancari/gocean/api"
	shell "github.com/nickmancari/gocean/ssh"
	convert "github.com/nickmancari/gocean/data"
)

var apiAddress string = "https://api.digitalocean.com/v2/droplets"

type OceanJson struct {
	Name               string `json: name`
	Region             string `json: region`
	Size               string `json: size`
	Image              string `json: image`
	Backups            string `json: backups`
	Ipv6               string `json: ipv6`
	User_Data          string `json: user_data`
	Private_Networking string `json: private_networking`
}

func CreateDroplet(f string) (interface{}, error) {
	if f == "" {
		return "", nil
	} else {

		name := f

		jsonData, err := json.Marshal(OceanJson{Name: name, Region: "nyc3", Size: "s-1vcpu-1gb", Image: "ubuntu-16-04-x64", Backups: "false", Ipv6: "true", User_Data: "null", Private_Networking: "null"})
		if err != nil {
			return fmt.Println(err)
		}

		r := connect.ConvertConnection("POST", apiAddress, bytes.NewBuffer(jsonData))
		c, err := convert.StructuredResponse(r)

		return c, nil
	}
}

func DestroyDroplet(f string) (interface{}, error) {
	if f == "" {
		return "", nil
	} else {
		id, err := convert.ToID(f)
		if err != nil {
			return fmt.Println(err)
		}
		r := connect.Connection("DELETE", apiAddress+"/"+id, nil)
		if r > 0 {
			c, err := fmt.Println("-----------------\n|Droplet Deleted|\n-----------------\n")
			if err != nil {
				fmt.Println(err)
			}
			return c, nil
		}
		return "", nil
	}
}

func GetDroplet(f string) (interface{}, error) {
	if f == "" {
		return "", nil
	} else if f == "all" {
		r := connect.ConvertConnection("GET", apiAddress, nil)
		c, err := convert.AllDroplets(r)
		if err != nil {
			return fmt.Println(err)
		}
		return c, nil
	} else {
		id, err := convert.ToID(f)
		if err != nil {
			return fmt.Println(err)
		}
//testing	r := connect.Connection("GET", apiAddress+"/"+id, nil)
		r := connect.ConvertConnection("GET", apiAddress+"/"+id, nil)
		c, err := convert.StructuredResponse(r)
		return c, nil
	}
}

func RebootDroplet(f string) (interface{}, error) {
	if f == "" {
		return "", nil
	} else {
		jsonData := []byte(`{"type":"reboot"}`)
		id, err := convert.ToID(f)
		if err != nil {
			return fmt.Println("Program Error: ", err)
		}

		r := connect.Connection("POST", apiAddress+"/"+id+"/actions", bytes.NewBuffer(jsonData))

		return r, nil
	}

}

func Shell(f string) (interface{}, error) {
	if f == "" {
		return "", nil
	} else {
		r, err := shell.Session(f)
		if err != nil {
			return fmt.Println(err)
		}

		return r, nil
	}
}
