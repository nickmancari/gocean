package convert

import (
	"encoding/json"
	"fmt"
	"errors"
	"strings"

	connect "github.com/nickmancari/gocean/api"
	color "github.com/nickmancari/gocean/pkg"
)

var apiGetAddress = "https://api.digitalocean.com/v2/droplets"

type droplets struct {
	Droplets []info `json:"droplets"`
}

type info struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Status	 string	  `json:"status"`
	Image	 OsInfo	  `json:"image"`
	Networks networks `json:"networks"`
}

type OsInfo struct {
	Name	string	  `json:"name"`
	Distro	string	  `json:"distribution"`
}

type networks struct {
	V4 	[]netInfo `json:"v4"`
}

type netInfo struct {
	IP      string `json:"ip_address"`
//	Netmask string `json:"netmask"`
//	Gateway string `json:"gateway"`
//	Type    string `json:"type"`
}

// Takes the input of droplet name and converts it
// into the droplet ID to use with the DO API

func ToID(s string) (string, error) {
	body := connect.ConvertConnection("GET", apiGetAddress, nil)

	dropletStruct := droplets{}
	er := json.Unmarshal(body, &dropletStruct)
	if er != nil {
		fmt.Println("Unmarshal Error: ", er)
	}

	search := s
	for _, id := range dropletStruct.Droplets {
		if strings.Contains(id.Name, search) {
			r := fmt.Sprint(id.ID)
			return r, nil
		}
	}
	err := errors.New("\nDroplet Not Found\n")
	return "", err
}

// Takes the input of droplet name and converts it
// into the droplet IP address for network manipulation

func ToIP(s string) (string, error) {
	body := connect.ConvertConnection("GET", apiGetAddress, nil)

	dropletStruct := droplets{}
	er := json.Unmarshal(body, &dropletStruct)
	if er != nil {
		fmt.Println("Unmarshal Error: ", er)
	}

	search := s
	for _, id := range dropletStruct.Droplets {
		if strings.Contains(id.Name, search) {
			if len(id.Networks.V4) > 1 {
				r := fmt.Sprint(id.Networks.V4[1].IP)
				return r, nil
			} else {
				r := fmt.Sprint(id.Networks.V4[0].IP)
				return r, nil
			}
		}
	}
	err := errors.New("\nDroplet Not Found\n")
	return "", err

}

func AllDroplets(b []byte) (interface{}, error) {
	dropletStruct := droplets{}
	er := json.Unmarshal(b, &dropletStruct)
	if er != nil {
		return fmt.Println("Unmarsahl Error: ", er)
	}

	for _, v := range dropletStruct.Droplets {
		if strings.Contains(v.Status, "active") {
			fmt.Printf("\n|"+color.Cyan+"%s"+color.Reset+"|\n-------------------------------------------------------\n|ID: %6d||Status: "+color.Green+"%6s"+color.Reset+"||Distro: %6s||Network: "+color.Blue+"%6s"+color.Reset+"|\n\n", v.Name, v.ID, v.Status, v.Image, v.Networks.V4) 
		} else {
			fmt.Printf("\n|"+color.Cyan+"%s"+color.Reset+"|\n-------------------------------------------------------\n|ID: %6d||Status: "+color.Red+"%s"+color.Reset+"||Distro: %6s||Network: "+color.Blue+"%6s"+color.Reset+"|\n\n", v.Name, v.ID, v.Status, v.Image, v.Networks.V4)
		}
	}
	
	return "", nil


}
