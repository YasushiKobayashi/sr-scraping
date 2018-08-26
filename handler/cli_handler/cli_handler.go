package cli_handler

import "github.com/urfave/cli"

type (
	headless string
)

func margeBaseFlag(f []cli.Flag) []cli.Flag {
	var baseFlag []cli.Flag = []cli.Flag{
		cli.StringFlag{
			Name:   "account, a",
			Value:  "",
			EnvVar: "account_id",
			Usage:  "your showroom acount_id",
		},
		cli.StringFlag{
			Name:   "password, p",
			Value:  "",
			EnvVar: "password",
			Usage:  "your showroom password",
		},
		cli.BoolFlag{
			Name:  "headless",
			Usage: "use headless chrome",
		},
	}
	return append(baseFlag, f...)
}

var Commands = []cli.Command{
	countAllfollow,
}
