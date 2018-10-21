package scraping

import (
	"time"

	"github.com/mono0x/puroland-greeting/model"
)

type FetcherMock struct {
	MockFetchIndexPage       func() (*model.IndexPage, error)
	MockFetchSecretIndexPage func(date time.Time) (*model.IndexPage, error)
	MockFetchMenuPage        func(path string) (*model.MenuPage, error)
	MockFetchCharacterPage   func(path string) (*model.CharacterPage, error)
}

func (f *FetcherMock) FetchIndexPage() (*model.IndexPage, error) {
	return f.MockFetchIndexPage()
}

func (f *FetcherMock) FetchSecretIndexPage(date time.Time) (*model.IndexPage, error) {
	return f.MockFetchSecretIndexPage(date)
}

func (f *FetcherMock) FetchMenuPage(path string) (*model.MenuPage, error) {
	return f.MockFetchMenuPage(path)
}

func (f *FetcherMock) FetchCharacterPage(path string) (*model.CharacterPage, error) {
	return f.MockFetchCharacterPage(path)
}
