package commands

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateDroplet(name string, token string) {
	apiAddress := "https://api.digitalocean.com/v2/droplets"

	jsonData := []byte(`{"name": name, "region": "nyc3", "size": "s-1vcpu-1gb", "image": "ubuntu-16-04-x64", "backups": "false", "ipv6": true, "user_data": null, "private_networking": null, "volumes": null}`)

	request, err := http.NewRequest("POST", apiAddress, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+token)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))
}

func DestroyDroplet() {

}
