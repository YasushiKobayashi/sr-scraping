package cli_handler

import (
	"strconv"

	"github.com/YasushiKobayashi/countrobu/infrastructure/driver_repository"
	"github.com/YasushiKobayashi/countrobu/model"
	"github.com/YasushiKobayashi/countrobu/usecase"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

var getStarsFlag []cli.Flag = []cli.Flag{
	cli.StringFlag{
		Name:  "line, l",
		Value: "1",
		Usage: "palallel line number",
	},
}

var getStars = cli.Command{
	Name:    "get-stars",
	Aliases: []string{"stars"},
	Usage:   "...",
	Description: `
`,
	Action: getStarsHandler,
	Flags:  margeBaseFlag(getStarsFlag),
}

type (
	GetStarsHandler struct {
		Interactor usecase.ShowRoomInteractor
	}
)

func NewGetStarsHandler(headless bool) *GetStarsHandler {
	return &GetStarsHandler{
		Interactor: usecase.ShowRoomInteractor{
			Repository: &driver_repository.DriverRepository{
				Headless: headless,
			},
		},
	}
}

func getStarsHandler(c *cli.Context) error {
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

	handler := NewGetStarsHandler(c.Bool("headless"))
	err = handler.Interactor.GetStars(user, paralleLine)
	if err != nil {
		return cli.NewExitError(err, 128)
	}
	return nil
}
