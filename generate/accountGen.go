package generate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type Payload struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  Params `json:"params"`
	ID      int    `json:"id"`
}
type Params struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Curl(hosts []string) error {
	data := Payload{
		Jsonrpc: "2.0",
		Method: "avm.createAddress",
		Params: Params{
			Username: "user",
			Password: "password",
		},
		ID: 1,
	}

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		log.Error("error marshaling payload: ", err)
	}
	body := bytes.NewReader(payloadBytes)

	for _, host := range hosts {
		req, err := http.NewRequest("POST", fmt.Sprintf("http://%s:9655/ext/bc/X", host), body)
		if err != nil {
			log.Error("error creating new http request: ", err)
		}
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Error("error sending request: ", err)
		}
		defer resp.Body.Close()
	}

	return nil
}