package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	token "github.com/nickmancari/gocean/env"
)

var apiToken = token.ReadTokenFile(".token")

const (
	apiAddress = "https://api.digitalocean.com/v2/droplets"
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

func CreateDroplet(flag string) interface{} {
	if flag == "" {
		s, err := fmt.Println("")
		if err != nil {
			fmt.Println("Error: ", err)
		}
		return s
	} else {

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
		request.Header.Add("Authorization", "Bearer "+apiToken)

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

func DestroyDroplet(flag string) interface{} {
	if flag == "" {
		s, err := fmt.Println("")
		if err != nil {
			fmt.Println("Error: ", err)
		}
		return s
	} else {
		dropletName := flag
		r, err := fmt.Printf("Deleting: %s %s", dropletName, apiToken)
		if err != nil {
			fmt.Println(err)
		}
		return r
	}
}

func GetDroplet(flag string) interface{} {
	if flag == "" {
		s, err := fmt.Println("")
		if err != nil {
			fmt.Println(err)
		}
		return s
	} else {
		request, err := http.NewRequest("GET", apiAddress + "/" + flag, nil)
		if err != nil {
			fmt.Println(err)
		}

		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Authorization", "Bearer "+apiToken)

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
		r, err := fmt.Println("response Body: ", string(body))
		if err != nil {
			fmt.Println(err)
		}

		return r
	}
}

func RebootDroplet() {

}
