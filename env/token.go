package token

import (
	"flag"
	"fmt"
	"log"
	"os"
)

\\func main() {
\\	tokenFlag := flag.String("token", "", "Digital Ocean API Token")
\\	flag.Parse()
\\	CreateTokenFile(*tokenFlag)
\\
\\	fmt.Println("test successful")
\\}

func CreateTokenFile(tokenInput string) {

	file, err := os.Create(".token")
	if err != nil {
		log.Fatal(err)
	}

	file.Write([]byte(tokenInput + "\n"))
	file.Close()
}

