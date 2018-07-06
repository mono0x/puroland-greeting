package greeting

import (
	"net/http"
	"time"
)

type CharacterPagesFetcher struct {
	PageURL  string
	Client   *http.Client
	Duration time.Duration
}

func (f *CharacterPagesFetcher) pageURL() string {
	if f.PageURL == "" {
		return IndexPageURL
	}
	return f.PageURL
}

func (f *CharacterPagesFetcher) client() *http.Client {
	if f.Client == nil {
		return http.DefaultClient
	}
	return f.Client
}

func (f *CharacterPagesFetcher) sleep() {
	if f.Duration > 0 {
		time.Sleep(f.Duration)
	}
}

func (f *CharacterPagesFetcher) Fetch() ([]*CharacterPage, error) {
	indexPage, err := f.fetchIndexPage()
	if err != nil {
		return nil, err
	}

	f.sleep()

	menuPage, err := f.fetchMenuPage(indexPage.MenuPageURL)
	if err != nil {
		return nil, err
	}

	characterPages := make([]*CharacterPage, 0, len(menuPage.Items))
	for _, item := range menuPage.Items {
		f.sleep()

		characterPage, err := f.fetchCharacterPage(item.CharacterPageURL)
		if err != nil {
			return nil, err
		}
		characterPages = append(characterPages, characterPage)
	}
	return characterPages, nil
}

func (f *CharacterPagesFetcher) fetchIndexPage() (*IndexPage, error) {
	client := f.client()

	res, err := client.Get(f.pageURL())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	parser := &Parser{
		IndexPageURL: f.pageURL(),
	}

	page, err := parser.ParseIndexPage(res.Body)
	if err != nil {
		if err, ok := err.(*SecretError); ok {
			f.sleep()

			res, err := client.Get(parser.GetSecretIndexPageURL(err.Date))
			if err != nil {
				return nil, err
			}
			defer res.Body.Close()

			page, err := parser.ParseIndexPage(res.Body)
			if err != nil {
				return nil, err
			}
			return page, nil
		}
		return nil, err
	}
	return page, nil
}

func (f *CharacterPagesFetcher) fetchMenuPage(menuPageURL string) (*MenuPage, error) {
	client := f.client()

	res, err := client.Get(menuPageURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	parser := &Parser{
		IndexPageURL: f.pageURL(),
	}

	page, err := parser.ParseMenuPage(res.Body)
	if err != nil {
		return nil, err
	}
	return page, nil
}

func (f *CharacterPagesFetcher) fetchCharacterPage(characterPageURL string) (*CharacterPage, error) {
	client := f.client()

	res, err := client.Get(characterPageURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	parser := &Parser{
		IndexPageURL: f.pageURL(),
	}

	page, err := parser.ParseCharacterPage(res.Body)
	if err != nil {
		return nil, err
	}
	return page, nil
}
