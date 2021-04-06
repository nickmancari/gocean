package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	token "github.com/nickmancari/gocean/env"
)

const (
	apiAdress = "https://api.digitalocean.com/v2/droplets"
)

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

func CreateDroplet(flag string) int {
	if flag == "" {
		i, err := fmt.Println(1)
		if err != nil {
			fmt.Println(err)
		}
		return i
	} else {
		
		token := token.ReadTokenFile(".token")
		name := flag

		jsonData, err := json.Marshal(OceanJson{Name: name, Region: "nyc3", Size: "s-1vcpu-1gb", Image: "ubuntu-16-04-x64", Backups: "false", Ipv6: "true", User_Data: "null", Private_Networking: "null"})
		if err != nil {
			fmt.Println(err)
		}

		request, err := http.NewRequest("POST", apiAddress, bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Println(err)
		}

		request.Header.Add("Content-Type", "application/json")
		request.Header.Add("Authorization", "Bearer "+token)

		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			fmt.Println(err)
		}
		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		}
		r, err := fmt.Println("response Body:", string(body))
		if err != nil {
			fmt.Println(err)
		}

		return r
	}
}

func DestroyDroplet(flag string) int {
	if flag == "" {
		i, err := fmt.Println(1)
		if err != nil {
			fmt.Println(err)
		}
		return i
	} else {
		token := token.ReadTokenFile(".token")
		dropletName := flag

}

func GetDroplet() {

}
