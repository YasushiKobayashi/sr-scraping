package cli_handler

import (
	"strconv"

	"github.com/YasushiKobayashi/countrobu/infrastructure/driver_repository"
	"github.com/YasushiKobayashi/countrobu/model"
	"github.com/YasushiKobayashi/countrobu/usecase"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

var countAllfollowFlag []cli.Flag = []cli.Flag{
	cli.StringFlag{
		Name:  "line, l",
		Value: "1",
		Usage: "palallel line number",
	},
}

var countAllfollow = cli.Command{
	Name:    "count-all-follow",
	Aliases: []string{"all"},
	Usage:   "...",
	Description: `
`,
	Action: countAllfollowHandler,
	Flags:  margeBaseFlag(countAllfollowFlag),
}

type (
	CountHandler struct {
		Interactor usecase.CountInteractor
	}
)

func NewCountHandler(headless bool) *CountHandler {
	return &CountHandler{
		Interactor: usecase.CountInteractor{
			Repository: &driver_repository.DriverRepository{
				Headless: headless,
			},
		},
	}
}

func countAllfollowHandler(c *cli.Context) error {
	accountId := c.String("account")
	password := c.String("password")
	line := c.String("line")
	paralleLine, err := strconv.Atoi(line)
	if err != nil {
		err = errors.Wrap(err, "parallel must number.")
		return cli.NewExitError(err, 128)
	}

	user, err := model.SetUser(accountId, password)
	if password == "" {
		return cli.NewExitError(err, 128)
	}

	handler := NewCountHandler(c.Bool("headless"))
	err = handler.Interactor.Count(user, paralleLine)
	if err != nil {
		return cli.NewExitError(err, 128)
	}
	return nil
}
