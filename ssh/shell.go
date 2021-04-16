package shell

import (
	"fmt"

	convert "github.com/nickmancari/gocean/tools"
)

func Session(d string) interface{} {

	ip, err := convert.ToIP(d)
	if err != nil {
		fmt.Println(err)
	}

	r, err := fmt.Println(ip)
	if err != nil {
		fmt.Println(err)
	}

	return r

}
