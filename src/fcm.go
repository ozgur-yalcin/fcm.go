package fcm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const EndPoint = "https://fcm.googleapis.com/fcm/send"

type API struct {
	Key string
}

type (
	Request struct {
		Data                  interface{}  `json:"data,omitempty"`
		Notification          Notification `json:"notification,omitempty"`
		RegistrationIDs       []string     `json:"registration_ids,omitempty"`
		To                    string       `json:"to,omitempty"`
		CollapseKey           string       `json:"collapse_key,omitempty"`
		Priority              string       `json:"priority,omitempty"`
		RestrictedPackageName string       `json:"restricted_package_name,omitempty"`
		Condition             string       `json:"condition,omitempty"`
		TimeToLive            int          `json:"time_to_live,omitempty"`
		DryRun                bool         `json:"dry_run,omitempty"`
		ContentAvailable      bool         `json:"content_available,omitempty"`
		DelayWhileIdle        bool         `json:"delay_while_idle,omitempty"`
		MutableContent        bool         `json:"mutable_content,omitempty"`
	}
	Response struct {
		Ok           bool
		StatusCode   int
		MulticastID  int    `json:"multicast_id"`
		Success      int    `json:"success"`
		Failure      int    `json:"failure"`
		CanonicalIDs int    `json:"canonical_ids"`
		MsgID        int    `json:"Request_id,omitempty"`
		Err          string `json:"error,omitempty"`
		Results      Result `json:"results,omitempty"`
	}
	Result []struct {
		MessageID      string `json:"message_id"`
		RequestID      string `json:"request_id"`
		RegistrationID string `json:"registration_id"`
		Error          string `json:"error"`
	}
	Notification struct {
		Title        string `json:"title,omitempty"`
		Body         string `json:"body,omitempty"`
		Sound        string `json:"sound,omitempty"`
		Badge        string `json:"badge,omitempty"`
		Icon         string `json:"icon,omitempty"`
		Tag          string `json:"tag,omitempty"`
		Color        string `json:"color,omitempty"`
		ClickAction  string `json:"click_action,omitempty"`
		BodyLocKey   string `json:"body_loc_key,omitempty"`
		BodyLocArgs  string `json:"body_loc_args,omitempty"`
		TitleLocKey  string `json:"title_loc_key,omitempty"`
		TitleLocArgs string `json:"title_loc_args,omitempty"`
	}
)

func (api *API) Send(Request *Request) (response Response) {
	cli := new(http.Client)
	postdata, _ := json.Marshal(Request)
	req, err := http.NewRequest("POST", EndPoint, bytes.NewReader(postdata))
	if err != nil {
		fmt.Println(err)
		return response
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "key="+api.Key)
	res, err := cli.Do(req)
	if err != nil {
		fmt.Println(err)
		return response
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	decoder.UseNumber()
	decoder.Decode(&response)
	return response
}
