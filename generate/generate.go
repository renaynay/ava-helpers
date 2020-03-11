package generate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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
	AssetId  string	`json:"assetID"`
	Amount	 int `json:"amount"`
	To       string `json:"to"`
}

type Account string

func CreateAccounts(hosts []string, payload Payload) ([]Account, error) {
	payloadBytes, err := json.Marshal(payload)

	if err != nil {
		log.Error("error marshaling payload: ", err)
	}
	body := bytes.NewReader(payloadBytes)

	accounts := make([]Account, len(hosts))

	for i, host := range hosts {
		resp, err := Curl(host, body)
		if err != nil {
			return []Account{}, err
		}

		acct, err := readBody(resp.Body)
		if err != nil {
			log.Error("could not read body of response: ", err)
		}

		accounts[i] = acct

		defer resp.Body.Close()
	}

	return accounts, nil
}

func readBody(rawBody io.ReadCloser) (Account, error) {
	rawAcct := make([]byte, 36)

	bytesRead, err := rawBody.Read(rawAcct)
	if bytesRead != 36 || err != nil {
		return Account(""), err
	}

	return Account(string(rawAcct)), nil
}

func SendTxs(hosts []string, payload Payload) error {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Error("error marshaling payload: ", err)
		return err
	}
	body := bytes.NewReader(payloadBytes)

	for _, host := range hosts {
		resp, err := Curl(host, body)
		if err != nil {
			return err
		}

		defer resp.Body.Close()
	}

	return nil
}

func Curl(host string, body *bytes.Reader) (*http.Response, error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("http://%s:9655/ext/bc/X", host), body)
	if err != nil {
		log.Error("error creating new http request: ", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error("error sending request: ", err)
	}

	return resp, nil
}