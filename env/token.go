package token

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func CreateTokenFile(t string) {

	if t == "" {
		return
	} else {
		file, err := os.Create(".token")
		if err != nil {
			log.Fatal(err)
		}

		file.Write([]byte(t))
		file.Close()
	}
}

func ReadTokenFile(path string) string {

	token, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Token File Not Found", err)
	}
	content := string(token)
	r := strings.TrimSuffix(content, "\n")
	return r
}
