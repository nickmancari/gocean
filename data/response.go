package convert

import (
	"encoding/json"
	"fmt"
	color "github.com/nickmancari/gocean/pkg"
)

type droplet struct {
	Droplet system `json:"droplet"`
}

type system struct {
	ID     int       `json:"id"`
	Name   string    `json:"name"`
	Status string    `json:"status"`
	Image  imageInfo `json:"image"`
	Net    network   `json:"networks,omitempty"`
}

type imageInfo struct {
	Name   string `json:"name"`
	Distro string `json:"distribution"`
}

type network struct {
	V4 []netDetail `json:"v4,omitempty"`
}

type netDetail struct {
	IP      string `json:"ip_address,omitempty"`
//	Netmask string `json:"netmask"`
//	Gateway string `json:"gateway"`
}

func StructuredResponse(b []byte) (interface{}, error) {
	dropletStruct := droplet{}
	er := json.Unmarshal(b, &dropletStruct)
	if er != nil {
		return fmt.Println("Unmarshal Error ", er)
	}

	r, err := fmt.Printf("\n|"+color.Cyan+"%s"+color.Reset+"|\n-------------------------------------------------------\n|ID: %6d||Status: "+color.Green+"%6s"+color.Reset+"||Distro: %6s||Network: %6s|\n\n", dropletStruct.Droplet.Name, dropletStruct.Droplet.ID, dropletStruct.Droplet.Status, dropletStruct.Droplet.Image, dropletStruct.Droplet.Net.V4)
	if err != nil {
		return fmt.Println(err)
	}

	return r, nil

}
