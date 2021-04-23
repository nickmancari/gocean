package shell

import (
	"fmt"

	convert "github.com/nickmancari/gocean/data"
)

func Session(d string) (interface{}, error) {

	ip, err := convert.ToIP(d)
	if err != nil {
		return fmt.Println(err)
	}

	r, err := fmt.Println(ip)
	if err != nil {
		return fmt.Println(err)
	}

	return r, nil

}
