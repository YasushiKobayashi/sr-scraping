package driver_repository

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
)

func newDocument(html string) (res *goquery.Document, err error) {
	r := strings.NewReader(html)
	document, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return res, errors.Wrap(err, "NewDocumentFromReader error")
	}

	return document, nil
}
