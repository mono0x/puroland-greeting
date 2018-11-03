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
	StartTime time.Time
	EndTime   time.Time
	Venue     string
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
	StartTime   time.Time
	EndTime     time.Time
	VenueId     int64
	CharacterId int64
}

type Venue struct {
	Id   int64
	Name string
}

type Character struct {
	Id   int64
	Name string
}

type CanonicalVenue struct {
	Id   int64
	Name string
}

type CanonicalCharacter struct {
	Id   int64
	Name string
}

type VenueCanonicalization struct {
	VenueId          int64
	CanonicalVenueId int64
}

type CharacterCanonicalization struct {
	CharacterId          int64
	CanonicalCharacterId int64
	Style                string
}
