package convert

import (
	"encoding/json"
	"fmt"
	color "github.com/nickmancari/gocean/pkg"
)

type action struct {
	Action ActionDetail `json:"action"`
}

type ActionDetail struct {
	Status string `json:"status,omitempty"`
	Type   string `json:"type"`
}

func Actions(b []byte) (interface{}, error) {
	actionStruct := action{}
	er := json.Unmarshal(b, &actionStruct)
	if er != nil {
		fmt.Println("Unmarshal Error: ", er)
	}

	return fmt.Printf("\n"+color.Yellow+"%s %s\n"+color.Reset+"\n", actionStruct.Action.Type, actionStruct.Action.Status)
}
