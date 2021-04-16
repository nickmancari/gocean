package convert

import (
	"encoding/json"
	"fmt"
	"strings"

	connect "github.com/nickmancari/gocean/api"
)

var apiGetAddress = "https://api.digitalocean.com/v2/droplets"

type droplets struct {
	Droplets []info `json:"droplets"`
}

type info struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Networks networks `json:"networks"`
}

type networks struct {
	V4 []netInfo `json:"v4"`
}

type netInfo struct {
	IP      string `json:"ip_address"`
	Netmask string `json:"netmask"`
	Gateway string `json:"gateway"`
	Type    string `json:"type"`
}

func ToID(s string) (string, error) {
	body := connect.ConvertConnection("GET", apiGetAddress, nil)

	dropletStruct := droplets{}
	er := json.Unmarshal(body, &dropletStruct)
	if er != nil {
		fmt.Println("Unmarshal Error: ", er)
	}

	var r int
	search := s
	for _, id := range dropletStruct.Droplets {
		if strings.Contains(id.Name, search) {
			r = id.ID
		}
	}
	c := fmt.Sprint(r)

	return c, nil
}

func ToIP(s string) (string, error) {
	body := connect.ConvertConnection("GET", apiGetAddress, nil)

	dropletStruct := droplets{}
	er := json.Unmarshal(body, &dropletStruct)
	if er != nil {
		fmt.Println("Unmarshal Error: ", er)
	}
	var r string
	search := s
	for _, id := range dropletStruct.Droplets {
		if strings.Contains(id.Name, search) {
			r = id.Networks.V4[0].IP
		}
	}
	c := fmt.Sprint(r)

	return c, nil

}
