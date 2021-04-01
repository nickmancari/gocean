package token

import (
	"log"
	"os"
)


func CreateTokenFile(tokenInput string) {

	file, err := os.Create(".token")
	if err != nil {
		log.Fatal(err)
	}

	file.Write([]byte(tokenInput))
	file.Close()
}

