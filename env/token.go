package token

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func CreateTokenFile(tokenInput string) {

	if tokenInput == "" {
		return
	} else {
		file, err := os.Create(".token")
		if err != nil {
			log.Fatal(err)
		}

		file.Write([]byte(tokenInput))
		file.Close()
	}
}

func ReadTokenFile(path string) string {

	token, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	content := string(token)
	r := strings.TrimSuffix(content, "\n")
	return r
}
