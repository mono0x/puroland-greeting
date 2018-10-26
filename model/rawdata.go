package model

type RawData struct {
	IndexPage      *IndexPage
	MenuPage       *MenuPage
	CharacterPages []*CharacterPage
}

func (r *RawData) ToRawGreetings() []*RawGreeting {
	var rawGreetings []*RawGreeting
	for _, characterPage := range r.CharacterPages {
		for _, item := range characterPage.Items {
			rawGreetings = append(rawGreetings, &RawGreeting{
				StartAt:   (*RawGreetingTime)(&item.StartAt),
				FinishAt:  (*RawGreetingTime)(&item.FinishAt),
				Place:     item.Place,
				Character: characterPage.Name,
			})
		}
	}
	return rawGreetings
}
