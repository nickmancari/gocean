package term

import (
	"fmt"

	"golang.org/x/crypto/ssh/terminal"
)

func Center(s string) (interface{}, error) {
	w, _, err := terminal.GetSize(0)
	if err != nil {
		return fmt.Printf("Error: %s", err)
	}

	return fmt.Sprintf("%*s", -w, fmt.Sprintf("%s", (w+len(s))/2, s)), nil
}
