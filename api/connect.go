package connect

import (
	"fmt"
	"io/ioutil"
	"net/http"

	token "github.com/nickmancari/gocean/env"
)

var apiToken = token.ReadTokenFile(".token")

func Connection(format, address, info string) interface{} {

	request, err := http.NewRequest(format, address, info)
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
	r, err := fmt.Println("Response: ", string(body))
	if err != nil {
		fmt.Println(err)
	}
	return r

}
