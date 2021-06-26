package cmd

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/nickmancari/gocean/api"
	convert "github.com/nickmancari/gocean/data"
	"github.com/nickmancari/gocean/pkg/color"
	"github.com/nickmancari/gocean/ssh"
)

var apiAddress string = "https://api.digitalocean.com/v2/droplets"

type Actions struct {
	Type string `json:"type"`
}

func CreateDroplet(f string) (interface{}, error) {
	if f == "" {
		return "", nil
	} else if f == "-i" {
		d, _ := DistroVersion(DistroSelection())
		return DropletCreation(d)
	} else {
		return fmt.Println("Input Not Recognized")
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
			r := api.Connection("DELETE", apiAddress+"/"+id, nil)
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
		r := api.ConvertConnection("GET", apiAddress, nil)
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

		r := api.ConvertConnection("GET", apiAddress+"/"+id, nil)
		c, err := convert.StructuredResponse(r)
		return c, nil
	}
}

func Action(d string, a []string) (interface{}, error) {
	if d == "" && len(a) == 0 {
		return "", nil
	} else if len(d) > 0 && len(a) == 0 {
		return fmt.Println("\nNo action chosen.\n")
	} else {
		jsonData, err := json.Marshal(Actions{Type: a[0]})
		if err != nil {
			fmt.Println(err)
		}

		id, err := convert.ToID(d)
		if err != nil {
			return fmt.Println(err)
		}

		r := api.ConvertConnection("POST", apiAddress+"/"+id+"/actions", bytes.NewBuffer(jsonData))
		c, err := convert.Actions(r)

		return c, nil
	}
}

//placeholder
func Shell(f string) (interface{}, error) {
	if f == "" {
		return "", nil
	} else if f == "-i" {

		list, err := GetDroplet("ls")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(list)
		ssh.Start().Run()
		return "", nil
	}
	return "", nil
}
