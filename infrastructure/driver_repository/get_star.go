package driver_repository

import (
	"fmt"
	"log"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/YasushiKobayashi/countrobu/model"
	"github.com/pkg/errors"
)

func (d *DriverRepository) GetStars(user *model.User) error {
	c, err := d.startChrome()
	if err = d.P.Navigate(baseUrl + "/onlive"); err != nil {
		return errors.Wrap(err, "Navigate error")
	}
	defer c.Stop()

	if err = d.Login(user); err != nil {
		return errors.Wrap(err, "Login error")
	}

	time.Sleep(5 * time.Second)
	if err = d.Click("#js-categorymenu-list > li:nth-of-type(3)"); err != nil {
		return errors.Wrap(err, "activate idol tab error")
	}

	providers, err := d.getProviders()
	if err != nil {
		return errors.Wrap(err, "getProviders error")
	}

	for k, v := range providers {
		err := d.RunScript(fmt.Sprintf(`window.open("%s");`, v))
		if err != nil {
			log.Printf("window.open err %+v", err)
			continue
		}

		err = d.RunScript(`this.$('#twitter-dialog').show()`)
		// // _, err = d.isOnLive()
		// if err != nil {
		// 	log.Printf("window.open err %+v", err)
		// 	continue
		// }
		// fmt.Println("end")
		// err = d.RunScript("window.close();")
		// if err != nil {
		// 	log.Printf("window.close err %+v", err)
		// 	continue
		// }
		fmt.Println(v)
		time.Sleep(5 * time.Second)
		if k == 30 {
			break
		}
	}
	time.Sleep(100 * time.Second)
	return nil
}

func (d *DriverRepository) getProviders() ([]string, error) {
	var res []string

	html, err := d.P.HTML()
	if err != nil {
		return res, errors.Wrap(err, "Html error")
	}

	doc, err := newDocument(html)
	doc.Find(".contentlist-list a.js-room-link").Each(func(_ int, s *goquery.Selection) {
		urlStr, exist := s.Attr("href")
		if exist {
			res = append(res, urlStr)
		}
	})
	return res, nil
}
