package convert

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	token "github.com/nickmancari/gocean/env"
)

var apiToken = token.ReadTokenFile("../.token")
var apiGetAddress = "https://api.digitalocean.com/v2/droplets"

type droplets struct {
	Droplets []info `json:"droplets"`
}

type info struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func ToID(s string) (int, error) {
	request, err := http.NewRequest("GET", apiGetAddress, nil)
	if err != nil {
		fmt.Println("HTTP Request Error: ", err)
		return
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+apiToken)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Client Error: ", err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("ReadAll Error: ", err)
		return
	}

	dropletStruct := droplets{}
	er := json.Unmarshal(body, &dropletStruct)
	if er != nil {
		fmt.Println("Unmarshall Error: ", err)
		return
	}

	var r []int
	search := s
	for _, id := range dropletStruct.Droplets {
		if strings.Contains(id.Name, search) {
			r = append(r, id.ID)
		}
	}
	return r, nil
}
