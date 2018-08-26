package main

import (
	"os"

	"github.com/YasushiKobayashi/countrobu/handler/cli_handler"
	"github.com/urfave/cli"
)

var Version string = "0.0.1"

func main() {
	app := cli.NewApp()
	app.Name = "countrobu"
	app.Usage = ""
	app.Author = "Yasushi Kobayashi"
	app.Email = "ptpadan@gmail.com"
	app.Version = Version
	app.Commands = cli_handler.Commands

	app.Run(os.Args)
}
