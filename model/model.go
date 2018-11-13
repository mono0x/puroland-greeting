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
	StartTime      time.Time
	EndTime        time.Time
	RawVenueId     int64
	RawCharacterId int64
}

type RawVenue struct {
	Id   int64
	Name string
}

type RawCharacter struct {
	Id   int64
	Name string
}

type Venue struct {
	Id   int64
	Name string
}

type Character struct {
	Id   int64
	Name string
}

type CharacterStyle struct {
	Id          int64
	CharacterId int64
	Name        string
}

type VenueCanonicalization struct {
	VenueId    int64
	RawVenueId int64
}

type CharacterCanonicalization struct {
	CharacterId      int64
	CharacterStyleId int64
	RawCharacterId   int64
}
