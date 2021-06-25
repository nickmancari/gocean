package convert

import (
	"encoding/json"
	"fmt"
	"github.com/nickmancari/gocean/pkg/color"
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
}

func StructuredResponse(b []byte) (interface{}, error) {
	dropletStruct := droplet{}
	er := json.Unmarshal(b, &dropletStruct)
	if er != nil {
		return fmt.Println("Unmarshal Error ", er)
	}

	v := dropletStruct.Droplet
		if v.Status == "active" {
			fmt.Printf("\n|"+color.Cyan+"%s"+color.Reset+"|\n-------------------------------------------------------\n|ID: %6d||Status: "+color.Green+"%6s"+color.Reset+"||Distro: %6s||Network: %6s|\n\n", v.Name, v.ID, v.Status, v.Image, v.Net.V4)
		}else if v.Status == "new" {
			fmt.Printf("\n|"+color.Cyan+"%s"+color.Reset+"|\n-------------------------------------------------------\n|ID: %6d||Status: "+color.Yellow+"%6s"+color.Reset+"||Distro: %6s||Network: %6s|\n\n", v.Name, v.ID, v.Status, v.Image, v.Net.V4)
		} else {
			fmt.Printf("\n|"+color.Cyan+"%s"+color.Reset+"|\n-------------------------------------------------------\n|ID: %6d||Status: "+color.Red+"%6s"+color.Reset+"||Distro: %6s||Network: %6s|\n\n", v.Name, v.ID, v.Status, v.Image, v.Net.V4)
		}


	return "", nil

}
