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
	preSess, err := d.GetSessID()
	if err != nil {
		return err
	}
	fmt.Println(preSess)

	if err = d.RunScript("showLoginDialog();"); err != nil {
		return err
	}

	err = d.SendKeyByJs("account_id", user.AccountId)
	if err != nil {
		return errors.Wrap(err, "SendKeyByJs account_id error")
	}

	err = d.SendKeyByJs("password", user.Password)
	if err != nil {
		return errors.Wrap(err, "SendKeyByJs password error")
	}

	for i := 0; ; {
		i++
		d.Click("#js-login-submit")

		sess, err := d.GetSessID()
		if err != nil {
			return err
		}

		if sess != preSess {
			break
		}

		time.Sleep(1 * time.Second)
		if i == 10 {
			return errors.New("login faile")
		}
	}

	return nil
}

func (d *DriverRepository) GetSessID() (res string, err error) {
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
