package main

import (
	"github.com/dionomusuko/cli/api"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()

	app.Action = func(context *cli.Context) error {
		str := context.Args().Get(0)
		api.Response(str)
		return nil
	}
	app.Run(os.Args)
}
