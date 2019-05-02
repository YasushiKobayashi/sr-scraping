package usecase

import (
	"log"
	"sync"

	"github.com/YasushiKobayashi/countrobu/model"
	"github.com/pkg/errors"
)

type (
	ShowRoomInteractor struct {
		Repository ShowRoomRepository
	}

	ShowRoomRepository interface {
		Count(*model.User, string) error
		GetFollowsOnlive(*model.User) ([]string, error)
		GetStars(*model.User) error
	}
)

func (i *ShowRoomInteractor) Count(user *model.User, paralleLine int) error {
	actors, err := i.getFollowsOnlive(user)
	if err != nil {
		return errors.Wrap(err, "GetFollowsOnlive error")
	}
	err = i.countOnLive(actors, user, paralleLine)
	return nil
}

func (i *ShowRoomInteractor) getFollowsOnlive(user *model.User) ([]string, error) {
	actors, err := i.Repository.GetFollowsOnlive(user)
	if err != nil {
		return actors, errors.Wrap(err, "GetFollowsOnlive error")
	}
	return actors, nil
}

func (i *ShowRoomInteractor) countOnLive(actors []string, user *model.User, paralleLine int) error {
	ch := make(chan bool, paralleLine)
	var wg sync.WaitGroup
	for _, v := range actors {
		wg.Add(1)
		ch <- true
		go func(v string) {
			defer func() { <-ch }()
			err := i.Repository.Count(user, v)
			if err != nil {
				log.Printf("count %+v", err)
			}

			wg.Done()
		}(v)
	}

	wg.Wait()
	return nil
}

func (i *ShowRoomInteractor) GetStars(user *model.User, paralleLine int) error {
	err := i.Repository.GetStars(user)
	if err != nil {
		return errors.Wrap(err, "GetFollowsOnlive error")
	}
	return nil
}
