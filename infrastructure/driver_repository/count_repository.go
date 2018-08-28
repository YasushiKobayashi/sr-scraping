package driver_repository

import (
	"fmt"
	"log"
	"time"

	"github.com/YasushiKobayashi/countrobu/model"
	"github.com/pkg/errors"
)

func (d *DriverRepository) Count(user *model.User, target string) error {
	c, err := d.startChrome()
	if err != nil {
		return errors.Wrap(err, "startChromeDriver error")
	}
	defer c.Stop()

	if err = d.runCount(user, target); err != nil {
		return errors.Wrap(err, "driver.Run error")
	}

	log.Printf("end count %s", target)
	return nil
}

func (d *DriverRepository) runCount(user *model.User, target string) (err error) {
	log.Println(target)
	if err = d.P.Navigate(baseUrl + target); err != nil {
		return errors.Wrap(err, "Navigate error")
	}

	if err = d.Login(user); err != nil {
		return errors.Wrap(err, "Login error")
	}
	fmt.Println("target")
	fmt.Println(target)

	if err = d.count(target); err != nil {
		return errors.Wrap(err, "count error")
	}
	return nil
}

func (d *DriverRepository) count(target string) (err error) {
	input := "#js-chat-input-comment"
	btn := "//*[@id='js-room-comment']//*[@type='submit']"

	title, err := d.P.Title()
	if err != nil {
		return errors.Wrap(err, "Title error")
	}

	time.Sleep(5 * time.Second)
	for i := 1; i <= 50; i++ {
		time.Sleep(2 * time.Second)

		err = d.SendKey(input, fmt.Sprint(i))
		if err != nil {
			return errors.Wrap(err, "SendKeyById error")
		}

		for {
			err = d.ClickByXPath(btn)
			if err != nil {
				return errors.Wrap(err, "ClickByXPath error")
			}

			val, err := d.GetValue(input)
			if err != nil {
				return errors.Wrap(err, "GetValue error")
			}

			if val != "" {
				break
			}

			err = d.SendKey(input, "")
			if err != nil {
				return errors.Wrap(err, "SendKeyById error")
			}
			err = d.SendKey(input, fmt.Sprint(i))
			if err != nil {
				return errors.Wrap(err, "SendKeyById error")
			}
			time.Sleep(2 * time.Second)
		}
		log.Printf("count number %d for %s %s", i, target, title)
	}
	return nil
}
