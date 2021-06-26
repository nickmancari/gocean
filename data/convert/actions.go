package convert

import (
	"encoding/json"
	"fmt"
	"github.com/nickmancari/gocean/pkg/color"
)

type action struct {
	Action ActionDetail `json:"action"`
	Message	string		`json:"message,omitempty"`
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
	
	if len(actionStruct.Message) > 0 {
		return fmt.Printf("\n%s\n\n", actionStruct.Message)
	} else {
		return fmt.Printf("\nDroplet Action: "+color.Yellow+"%s %s\n"+color.Reset+"\n", actionStruct.Action.Type, actionStruct.Action.Status)
	}
}
