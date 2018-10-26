package model

import (
	"time"
)

type IndexPage struct {
	Date         time.Time
	MenuPagePath string
}

type MenuPage struct {
	Items []*MenuPageItem
}

type MenuPageItem struct {
	CharacterName     string
	CharacterPagePath string
}

type CharacterPage struct {
	Name  string
	Date  time.Time
	Items []*CharacterPageItem
}

type CharacterPageItem struct {
	StartAt  time.Time
	FinishAt time.Time
	Place    string
}

type CharacterListPage struct {
	Date  time.Time
	Items []*CharacterListPageItem
}

type CharacterListPageItem struct {
	Name string
}

type RawGreetingTime time.Time

const rawGreetingTimeLayout = "2006-01-02 15:04:05 -0700"

func (r *RawGreetingTime) MarshalText() ([]byte, error) {
	return []byte((*time.Time)(r).Format(rawGreetingTimeLayout)), nil
}

func (r *RawGreetingTime) UnmarshalText(b []byte) error {
	t, err := time.Parse(rawGreetingTimeLayout, string(b))
	if err != nil {
		return err
	}
	*r = RawGreetingTime(t)
	return nil
}

type RawGreeting struct {
	Character string           `ltsv:"character"`
	Place     string           `ltsv:"place"`
	StartAt   *RawGreetingTime `ltsv:"start_at"`
	FinishAt  *RawGreetingTime `ltsv:"end_at"` // for compatibility
}

type Greeting struct {
	Id          int64
	Date        time.Time
	StartAt     time.Time
	FinishAt    time.Time
	PlaceId     int64
	CharacterId int64
	CostumeId   *int64
}

type PreNotice struct {
	Id          int64
	Date        time.Time
	CharacterId int64
	CostumeId   *int64
}

type Place struct {
	Id   int64
	Name string
}

type Character struct {
	Id   int64
	Name string
}

type Costume struct {
	Id          int64
	CharacterId int64
	Name        string
}
