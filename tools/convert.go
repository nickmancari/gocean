package convert

import (
	
	"gitub.com/nickmancari/gocean/env"

)

var apiToken = token.ReadTokenFile(".token")
var apiGetAddress = "https://api.digitalocean.com/v2/droplets?page=1&per_page=1"

func ToDropletID(f string) {
	request, err := http.NewRequest("GET", apiGetAddress, nil)
	if err != nil {
		fmt.Println(err)
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer ", +apiToken)


}
