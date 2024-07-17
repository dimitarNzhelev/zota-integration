package structs

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Code    string
	Message string
	Data    map[string]interface{} `json:"data"`
}

func (r Response) String() string {
    message := r.Message
    if message == "" {
        message = "No message"
    }
    data, _ := json.Marshal(r.Data)
    return fmt.Sprintf("Code: %s, Message: %s, Data: %s", r.Code, message, string(data))
}