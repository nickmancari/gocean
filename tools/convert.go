package convert

import (

)

func ToDropletID(f string) {
	request, err := http.NewRequest("GET", apiGetAddress, nil)
	if err != nil {
		fmt.Println(err)
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer ", +apiToken)
}
