package driver_repository

import (
	"fmt"

	"github.com/sclevine/agouti"
)

func (d *DriverRepository) RunScript(js string) (err error) {
	err = d.P.RunScript(js, nil, nil)
	if err != nil {
		return err
	}
	return nil
}

// Find
// Find finds exactly one element by CSS selector.
func (d *DriverRepository) Find(sel string) *agouti.Selection {
	return d.P.Find(sel)
}

func (d *DriverRepository) SendKeyByJs(name string, val string) (err error) {
	js := fmt.Sprintf(`$("[name='%s']").val('%s');`, name, val)
	return d.RunScript(js)
}

func (d *DriverRepository) Click(sel string) error {
	err := d.Find(sel).Click()
	if err != nil {
		return err
	}

	return nil
}

func (d *DriverRepository) ClickByXPath(sel string) error {
	err := d.P.FindByXPath(sel).Click()
	if err != nil {
		return err
	}

	return nil
}

func (d *DriverRepository) SendKey(sel string, val string) error {
	err := d.Find(sel).Fill(val)
	if err != nil {
		return err
	}

	return nil
}

func (d *DriverRepository) GetById(sel string, val string) {
	d.Find(sel).Fill(val)
}

func (d *DriverRepository) GetValue(sel string) (val string, err error) {
	val, err = d.Find(sel).Attribute("value")
	if err != nil {
		return val, err
	}

	return val, nil
}

func (d *DriverRepository) GetText(sel string) (val string, err error) {
	val, err = d.Find(sel).Text()
	if err != nil {
		return val, err
	}
	return val, nil
}

func (d *DriverRepository) isOnLive() (bool, error) {
	// text, _ := d.P.HTML()
	// fmt.Println(text)
	// return false, errors.New("a")
	path := "//*[@id='icon-room-twitter-wrapper']"
	return d.P.FindByXPath(path).Visible()
}
