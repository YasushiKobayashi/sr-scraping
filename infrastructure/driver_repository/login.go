package driver_repository

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/YasushiKobayashi/countrobu/model"
	"github.com/pkg/errors"
)

// Login
func (d *DriverRepository) Login(user *model.User) (err error) {
	time.Sleep(5 * time.Second)
	preSess, err := d.getSessID()
	if err != nil {
		return err
	}

	maxloop := 10
	for i := 0; ; {
		i++
		err = d.login(user)
		if err != nil && i == maxloop {
			return errors.Wrap(err, "login error")
		}

		err = d.Click("#js-login-submit")
		if err != nil && i == maxloop {
			return errors.Wrap(err, "submit error")
		}

		isLoggedin, err := d.isLoggedin(preSess)
		if err != nil && i == maxloop {
			return errors.Wrap(err, "isLoggedin error")
		}
		if isLoggedin {
			break
		}

		time.Sleep(1 * time.Second)
		if i == maxloop {
			return errors.New("login faile")
		}
	}

	return nil
}

func (d *DriverRepository) login(user *model.User) (err error) {
	if err = d.RunScript("showLoginDialog();"); err != nil {
		return errors.Wrap(err, "RunScript showLoginDialog error")
	}

	err = d.SendKey("#js-login-form input[name='account_id']", user.AccountId)
	if err != nil {
		return errors.Wrap(err, "SendKeyByJs account_id error")
	}

	err = d.SendKey("#js-login-form input[name='password']", user.Password)
	if err != nil {
		return errors.Wrap(err, "SendKeyByJs password error")
	}
	return nil
}

func (d *DriverRepository) isLoggedin(preSess string) (res bool, err error) {
	selID := "#js-login-error"
	text, err := d.GetText(selID)
	if err != nil {
		return false, errors.Wrap(err, "GetText error")
	}

	fmt.Println(text)
	if text != "" {
		return true, nil
	}

	sess, err := d.getSessID()
	if err != nil {
		return false, errors.Wrap(err, "getSessID account_id error")
	}

	return (sess != preSess), nil
}

func (d *DriverRepository) getSessID() (res string, err error) {
	sessKey := "sr_id"

	var wg sync.WaitGroup
	cookies, err := d.P.GetCookies()
	if err != nil {
		return res, errors.Wrap(err, "SendKeyByJs account_id error")
	}
	for _, v := range cookies {
		wg.Add(1)
		go func(v *http.Cookie) {
			if v.Name == sessKey {
				res = v.Value
			}
			wg.Done()
		}(v)
	}
	wg.Wait()
	return res, nil
}
