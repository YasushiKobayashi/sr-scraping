package driver_repository

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/YasushiKobayashi/countrobu/model"
	"github.com/YasushiKobayashi/countrobu/utils"
	"github.com/pkg/errors"
)

func (d *DriverRepository) GetFollowsOnlive(user *model.User) ([]string, error) {
	var res []string
	c, err := d.startChrome()
	if err != nil {
		return res, errors.Wrap(err, "startChromeDriver error")
	}
	defer c.Stop()

	urlStr := baseUrl + "/follow"
	if err = d.P.Navigate(urlStr); err != nil {
		return res, errors.Wrap(err, "Navigate error")
	}

	if err = d.Login(user); err != nil {
		return res, errors.Wrap(err, "Login error")
	}

	html, err := d.P.HTML()
	if err != nil {
		return res, errors.Wrap(err, "Html error")
	}

	doc, err := newDocument(html)
	doc.Find("#js-genre-section-onlive a.room-url").Each(func(_ int, s *goquery.Selection) {
		urlStr, exist := s.Attr("href")
		if exist {
			res = append(res, urlStr)
		}
	})

	return utils.UniqStringArray(res), nil
}
