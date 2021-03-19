package main

import (
	"encoding/json"
	"fmt"

	fcm "github.com/ozgur-soft/fcm/src"
)

func main() {
	api := &fcm.API{"api-key"}
	request := &fcm.Request{}
	request.To = ""           // Instance ID
	request.Priority = "high" // Priority (high,normal)
	request.Data = map[string]string{"msg": "Hello World"}
	request.Notification = fcm.Notification{Title: "Title", Body: "Body"}
	request.ContentAvailable = true
	response := api.Send(request)
	pretty, _ := json.MarshalIndent(response, " ", "\t")
	fmt.Println(string(pretty))
}
