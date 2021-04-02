package commands

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func CreateDroplet(token string) (int, error) {

	apiAddress := "https://api.digitalocean.com/v2/droplets"

	jsonData := []byte(`{"name": "example", "region": "nyc3", "size": "s-1vcpu-1gb", "image": "ubuntu-16-04-x64", "backups": "false", "ipv6": true, "user_data": null, "private_networking": null, "volumes": null}`)

	request, err := http.NewRequest("POST", apiAddress, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("http error: ", err)
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+token)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Client error: ", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Read Error: ", err)
	}
	r, err := fmt.Println("response Body:", string(body))
	return r, err
}

func DestroyDroplet() {

}
