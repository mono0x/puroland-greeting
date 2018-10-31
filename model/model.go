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

type Greeting struct {
	Id          int64
	Date        time.Time
	StartAt     time.Time
	FinishAt    time.Time
	PlaceId     int64
	CharacterId int64
}

type Place struct {
	Id   int64
	Name string
}

type Character struct {
	Id   int64
	Name string
}

type CanonicalPlace struct {
	Id   int64
	Name string
}

type CanonicalCharacter struct {
	Id   int64
	Name string
}

type Style struct {
	Id                   int64
	CanonicalCharacterId int64
	Name                 string
}

type CharacterCanonicalization struct {
	CharacterId          int64
	CanonicalCharacterId int64
	StyleId              *int64
}

type PlaceCanonicalization struct {
	PlaceId          int64
	CanonicalPlaceId int64
}
