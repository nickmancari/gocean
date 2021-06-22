package commands

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	connect "github.com/nickmancari/gocean/api"
	convert "github.com/nickmancari/gocean/data"
	color "github.com/nickmancari/gocean/pkg"
	shell "github.com/nickmancari/gocean/ssh"
//	term "github.com/nickmancari/gocean/term"
)

var apiAddress string = "https://api.digitalocean.com/v2/droplets"

type Actions struct {
	Type		string	`json:"type"`
}

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
	} else if f == "-i" {
		d, _ := DistroVersion(DistroSelection())
		return DropletCreation(d)
	} else {

		name := f

		jsonData, err := json.Marshal(OceanJson{Name: name, Region: "nyc3", Size: "s-1vcpu-1gb", Image: "ubuntu-16-04-x64", Backups: "false", Ipv6: "true", User_Data: "null", Private_Networking: "null"})
		if err != nil {
			return fmt.Println(err)
		}

		r := connect.ConvertConnection("POST", apiAddress, bytes.NewBuffer(jsonData))
	//	c, err := convert.StructuredResponse(r)
		c, _ := fmt.Println(string(r))

		return c, nil
	}
}

func DestroyDroplet(f string) (interface{}, error) {
	if f == "" {
		return "", nil
	} else {
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Println("\nAre you sure you want to delete this droplet? [Yes/no]")
		scanner.Scan()
		text := scanner.Text()

		if text == "Yes" {
			id, err := convert.ToID(f)
			if err != nil {
				return fmt.Println(err)
			}
			r := connect.Connection("DELETE", apiAddress+"/"+id, nil)
			if r > 0 {
				c, err := fmt.Println("-----------------\n|" + color.Red + "Droplet Deleted" + color.Reset + "|\n-----------------\n")
				if err != nil {
					fmt.Println(err)
				}
				return c, nil
			}

		} else {
			return fmt.Println("Deleting Droplet Canceled")
		}
		var c string
		return c, nil
	}
}

func GetDroplet(f string) (interface{}, error) {
	if f == "" {
		return "", nil
	} else if f == "ls" {
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

/* The "reboot", "shutdown", and "boot" individual flags will be replaced with
the more manageable "action" flag. So these seperate functions will be removed 
after the droplet creation interface is implemented*/
/*
func RebootDroplet(f string) (interface{}, error) {
	if f == "" {
		return "", nil
	} else {
		jsonData := []byte(`{"type":"reboot"}`)
		id, err := convert.ToID(f)
		if err != nil {
			return fmt.Println(err)
		}

		r := connect.ConvertConnection("POST", apiAddress+"/"+id+"/actions", bytes.NewBuffer(jsonData))
		c, err := convert.Actions(r)

		return c, nil
	}

}


func ShutdownDroplet(f string) (interface{}, error) {
	if f == "" {
		return "", nil
	} else {
		jsonData := []byte(`{"type":"power_off"}`)
		id, err := convert.ToID(f)
		if err != nil {
			return fmt.Println(err)
		}

		r := connect.ConvertConnection("POST", apiAddress+"/"+id+"/actions", bytes.NewBuffer(jsonData))
		c, err := convert.Actions(r)

		return c, nil
	}

}


func BootDroplet(f string) (interface{}, error) {
	if f == "" {
		return "", nil
	} else {
		jsonData := []byte(`{"type":"power_on"}`)
		id, err := convert.ToID(f)
		if err != nil {
			return fmt.Println(err)
		}

		r := connect.ConvertConnection("POST", apiAddress+"/"+id+"/actions", bytes.NewBuffer(jsonData))
		c, err := convert.Actions(r)

		return c, nil
	}

}
*/
func Action(d string, a []string) (interface{}, error) {
	if d == "" && len(a) == 0 {
		return "", nil
	} else if len(d) > 0 && len(a) == 0 {
		return fmt.Println("\nNo action chosen.\n")
	} else {
		jsonData, err := json.Marshal(Actions{Type:a[0],})
		if err != nil {
			fmt.Println(err)
		}
//testing		fmt.Println(string(jsonData))

		id, err := convert.ToID(d)
		if err != nil {
			return fmt.Println(err)
		}

		r := connect.ConvertConnection("POST", apiAddress+"/"+id+"/actions", bytes.NewBuffer(jsonData))
//		c, err := fmt.Println(string(r))
		c, err := convert.Actions(r)

		return c, nil
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
