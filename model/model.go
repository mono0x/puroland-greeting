package model

import "time"

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
	StartAt time.Time
	EndAt   time.Time
	Place   string
}

type CharacterListPage struct {
	Date  time.Time
	Items []*CharacterListPageItem
}

type CharacterListPageItem struct {
	Name string
}

type RawData struct {
	IndexPage      *IndexPage
	MenuPage       *MenuPage
	CharacterPages []*CharacterPage
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
