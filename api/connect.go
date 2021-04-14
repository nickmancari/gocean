package connect

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	token "github.com/nickmancari/gocean/env"
)

var apiToken = token.ReadTokenFile(".token")

func Connection(method string, url string, info io.Reader) interface{} {

	request, err := http.NewRequest(method, url, info)
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
