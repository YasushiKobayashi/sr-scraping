package driver_repository

import (
	"github.com/YasushiKobayashi/countrobu/app_error"
	"github.com/pkg/errors"
	"github.com/sclevine/agouti"
)

type (
	DriverRepository struct {
		P        *agouti.Page
		Headless bool
	}
)

const baseUrl = "https://www.showroom-live.com"

func (d *DriverRepository) startChrome() (res *agouti.WebDriver, err error) {
	res, err = d.startChromeDriver()
	if err != nil {
		return res, errors.Wrap(err, "start ChromeDriver error")
	}

	err = d.startChromeBrowser(res)
	if err != nil {
		return res, errors.Wrap(err, "start chromeBrowser error")
	}
	return res, nil
}

func (d *DriverRepository) startChromeDriver() (res *agouti.WebDriver, err error) {
	args := []string{
		"--window-size=1280,800",
	}
	if d.Headless {
		args = append(args, "--headless")
	}
	res = agouti.ChromeDriver(
		agouti.ChromeOptions("args", args),
		agouti.Debug,
	)
	if err := res.Start(); err != nil {
		err = errors.Wrap(err, "start chrome error")
		return res, app_error.NewErr(err)
	}
	return res, nil
}

func (d *DriverRepository) startChromeBrowser(w *agouti.WebDriver) error {
	page, err := w.NewPage(agouti.Browser("chrome"))
	if err != nil {
		err = errors.Wrap(err, "start chrome error")
		return app_error.NewErr(err)
	}

	d.P = page
	return nil
}
