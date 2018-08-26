package usecase

import (
	"sync"

	"github.com/YasushiKobayashi/countrobu/model"
	"github.com/pkg/errors"
)

type (
	CountInteractor struct {
		Repository CountDriverRepository
	}

	CountDriverRepository interface {
		Count(*model.User, string) error
		GetFollowsOnlive(*model.User) ([]string, error)
	}
)

func (i *CountInteractor) Count(user *model.User, paralleLine int) error {
	actors, err := i.getFollowsOnlive(user)
	if err != nil {
		return errors.Wrap(err, "GetFollowsOnlive error")
	}
	err = i.countOnLive(actors, user, paralleLine)
	return nil
}

func (i *CountInteractor) getFollowsOnlive(user *model.User) ([]string, error) {
	actors, err := i.Repository.GetFollowsOnlive(user)
	if err != nil {
		return actors, errors.Wrap(err, "GetFollowsOnlive error")
	}
	return actors, nil
}

func (i *CountInteractor) countOnLive(actors []string, user *model.User, paralleLine int) error {
	ch := make(chan bool, paralleLine)
	var wg sync.WaitGroup
	for _, v := range actors {
		wg.Add(1)
		ch <- true
		go func(v string) {
			defer func() { <-ch }()
			i.Repository.Count(user, v)

			wg.Done()
		}(v)
	}

	wg.Wait()
	return nil
}
