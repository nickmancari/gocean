package api

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/nickmancari/gocean/env"
)

var Token = env.ReadTokenFile("/usr/local/bin/.token")

func Connection(method string, url string, info io.Reader) int {

	request, err := http.NewRequest(method, url, info)
	if err != nil {
		fmt.Println(err)
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+Token)

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
	r, err := fmt.Println(string(body))
	if err != nil {
		fmt.Println(err)
	}
	return r
}

func ConvertConnection(method string, url string, info io.Reader) []byte {

	request, err := http.NewRequest(method, url, info)
	if err != nil {
		fmt.Println(err)
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+Token)

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
	return body

}
