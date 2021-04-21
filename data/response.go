package convert

import (
	"encoding/json"
	"fmt"
)

type droplet struct {
	Droplet system `json:"droplet"`
}

type system struct {
	ID     	int    		`json:"id"`
	Name   	string 		`json:"name"`
	Status 	string 		`json:"status"`
	Image	imageInfo	`json:"image"`
	Net	network		`json:"networks"`
}

type imageInfo struct {
	Name	string		`json:"name"`
	Distro	string		`json:"distribution"`
}

type network struct {
	V4	[]netDetail	`json:"v4"`
}

type netDetail struct {
	IP	string		`json:"ip_address"`
	Netmask	string		`json:"netmask"`
	Gateway	string		`json:"gateway"`
}

func ToStructuredResponse(b []byte) (interface{}, error) {
	dropletStruct := droplet{}
	er := json.Unmarshal(b, &dropletStruct)
	if er != nil {
		fmt.Println("Unmarshal Error ", er)
	}
	r, _ := fmt.Printf("\n|%6s|\n\n|ID: %6d||Status: %6s||Distro: %6s||Ip Address: %6s|\n\n", dropletStruct.Droplet.Name, dropletStruct.Droplet.ID, dropletStruct.Droplet.Status, dropletStruct.Droplet.Image, dropletStruct.Droplet.Net.V4[0].IP)

	return r, nil

}
