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

func CreateDroplet(f string) interface{} {
	if f == "" {
		return ""
	} else {

		name := f

		jsonData, err := json.Marshal(OceanJson{Name: name, Region: "nyc3", Size: "s-1vcpu-1gb", Image: "ubuntu-16-04-x64", Backups: "false", Ipv6: "true", User_Data: "null", Private_Networking: "null"})
		if err != nil {
			fmt.Println(err)
		}

		r := connect.Connection("POST", apiAddress, bytes.NewBuffer(jsonData))

		return r
	}
}

func DestroyDroplet(f string) interface{} {
	if f == "" {
		return ""
	} else {
		id, err := convert.ToID(f)
		if err != nil {
			fmt.Println(err)
		}
		r := connect.Connection("DELETE", apiAddress+"/"+id, nil)
		return r
	}
}

func GetDroplet(f string) interface{} {
	if f == "" {
		return ""
	} else {
		id, err := convert.ToID(f)
		if err != nil {
			fmt.Println(err)
		}
//testing	r := connect.Connection("GET", apiAddress+"/"+id, nil)
		r := connect.ConvertConnection("GET", apiAddress+"/"+id, nil)
		c, err := convert.StructuredResponse(r)
		return c
	}
}

func RebootDroplet(f string) interface{} {
	if f == "" {
		return ""
	} else {
		jsonData := []byte(`{"type":"reboot"}`)
		id, err := convert.ToID(f)
		if err != nil {
			fmt.Println("Conversion Error: ", err)
		}

		r := connect.Connection("POST", apiAddress+"/"+id+"/actions", bytes.NewBuffer(jsonData))

		return r
	}

}

func Shell(f string) interface{} {
	if f == "" {
		return ""
	} else {
		r := shell.Session(f)

		return r
	}
}
