package gocean

import (
	"fmt"
	"flag"
	"github.com/nickmancari/gocean/env"
)

func main() {

     tokenFlag := flag.String("token", "", "Digital Ocean API Token")
     flag.Parse()
     CreateTokenFile(*tokenFlag)

     fmt.Println("test successful")	

}
