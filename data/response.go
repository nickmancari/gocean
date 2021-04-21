package convert

import (
	"encoding/json"
	"fmt"
)

type droplet struct {
	Droplet system `json:"droplet"`
}

type system struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

func ToStructuredResponse(b []byte) (interface{}, error) {
	dropletStruct := droplet{}
	er := json.Unmarshal(b, &dropletStruct)
	if er != nil {
		fmt.Println("Unmarshal Error ", er)
	}
	r, _ := fmt.Printf("\n|%6s|\n\n|ID: %6d||Status: %6s|\n\n", dropletStruct.Droplet.Name, dropletStruct.Droplet.ID, dropletStruct.Droplet.Status)

	return r, nil

}
