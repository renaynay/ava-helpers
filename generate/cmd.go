package generate

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var (

	GenAccountsAndTxsCommand = cli.Command{
		Name:        "generate",
		Usage:       "Command to generate ava accounts and send txs between them",
		Description: `Creation of ava addresses / accounts`,
		Action:      generateAccounts,
		Flags:       []cli.Flag{},
	}

)

// curl -X POST --data '{
//    "jsonrpc":"2.0",
//    "id"     :3,
//    "method" :"avm.send",
//    "params" :{
//        "assetID" :"AVA",
//        "amount"  :10000,
//        "to"      :"X-xMrKg8uUECt5CS9RE9j5hizv2t2SWTbk",
//        "username":"userThatControlsAtLeast10000OfThisAsset",
//        "password":"myPassword"
//    }
//}' -H 'content-type:application/json;' 127.0.0.1:9650/ext/bc/X

func generateAccounts() error {
	ips := strings.Split(os.Args[2], ",")

	accounts, err := CreateAccounts(ips, Payload{
		Jsonrpc: "2.0",
		Method: "avm.createAddress",
		Params: Params{
			Username: "user",
			Password: "password",
		},
		ID: 1,
	})
	if err != nil {
		log.Error("error sending curl command to generate accounts: ", err)
	}

	err = generateTxs(ips, accounts)
	if err != nil {
		log.Error("error sending curl command to generate txs: ", err)
	}

	return nil
}

func generateTxs(ips []string, accounts []Account) error {
	// TODO figure out how to randomly send txs over and over again

	return nil
}
