package greeting

import "time"

type IndexPage struct {
	Date         time.Time
	MenuPagePath string
}

type SecretError struct {
	Date time.Time
}

func (err *SecretError) Error() string {
	return "" // unused
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
