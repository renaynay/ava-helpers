package generate

import (
	"os"
	"strings"

	"github.com/urfave/cli"
)

var (
	GenerateCommand = cli.Command{
		Name:        "generate",
		Usage:       "Command to generate ava accounts",
		Description: `Creation of ava addresses / accounts`,
		Action:      generateAccounts,
		Flags:       []cli.Flag{},
	}
)

func generateAccounts(ctx *cli.Context) error {
	ips := strings.Split(os.Args[2], ",")

	return Curl(ips)
}
