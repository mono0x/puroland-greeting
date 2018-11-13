package updater

import (
	"fmt"

	"github.com/Songmu/go-ltsv"
	"github.com/jmoiron/sqlx"

	"github.com/mono0x/puroland-greeting/scraping"
)

type Updater interface {
	Update() error
}

type updaterImpl struct {
	walker scraping.Walker
	db     *sqlx.DB
}

func NewUpdater(walker scraping.Walker, db *sqlx.DB) Updater {
	return &updaterImpl{
		walker: walker,
		db:     db,
	}
}

func (u *updaterImpl) Update() error {
	rawData, err := u.walker.Walk()
	if err != nil {
		return err
	}

	for _, rawGreeting := range rawData.ToRawGreetings() {
		dump, err := ltsv.Marshal(rawGreeting)
		if err != nil {
			return err
		}
		fmt.Println(string(dump))
	}
	return nil
}
