package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	token "github.com/nickmancari/gocean/env"
)

var apiToken = token.ReadTokenFile("../../.token")
var apiGetAddress = "https://api.digitalocean.com/v2/droplets"

type droplets struct {
	Droplets []info `json:"droplets"`
}

type info struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Networks	networks	`json:"networks"`
}

type networks struct {
	V4	[]netInfo	`json:"v4"`
}

type netInfo struct {
	IP	string	`json:"ip_address"`
	Netmask	string	`json:"netmask"`
	Gateway string	`json:"gateway"`
	Type	string	`json:"type"`
}

func main() {
	request, err := http.NewRequest("GET", apiGetAddress, nil)
	if err != nil {
		fmt.Println("HTTP Request Error: ", err)
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+apiToken)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("Client Error: %v", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("ReadAll Error: %v", err)
	}

	dropletStruct := droplets{}
	er := json.Unmarshal(body, &dropletStruct)
	if er != nil {
		fmt.Println("Unmarshall Error: %s", err)
	}

	//var r int
	search := "Kitchen"
	for _, id := range dropletStruct.Droplets {
		if strings.Contains(id.Name, search) {
			fmt.Println(id.Networks.V4[0].IP)
		}
	}
	//fmt.Println(r)

	//replace := strings.NewReplacer("[", "\"", "]", "\"")
	//conversion, err := fmt.Println(replace.Replace(c))
	//if err != nil {
	//	fmt.Printf("Conversion Error: %v", err)
	//}
	//fmt.Println(conversion)
}
