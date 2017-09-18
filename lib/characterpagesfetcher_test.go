package greeting

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFetch(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/chara_gre/mobile/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "data/www.puroland.co.jp/chara_gre/mobile/index.asp")
	})
	mux.HandleFunc("/chara_gre/mobile/chara_sentaku.asp", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "data/www.puroland.co.jp/chara_gre/mobile/chara_sentaku.asp")
	})
	mux.HandleFunc("/chara_gre/mobile/chara_sche.asp", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "data/www.puroland.co.jp/chara_gre/mobile/chara_sche.asp")
	})
	ts := httptest.NewServer(mux)
	defer ts.Close()

	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		t.Fatal(err)
	}

	f := &CharacterPagesFetcher{
		PageURL:  ts.URL + "/chara_gre/mobile/",
		Duration: 0,
	}

	pages, err := f.Fetch()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 12, len(pages))
	assert.Equal(t, "キティ・ホワイト", pages[0].Name)
	assert.Equal(t, time.Date(2016, time.June, 15, 0, 0, 0, 0, loc), pages[0].Date)
	assert.Equal(t, 2, len(pages[0].Items))
	assert.Equal(t, time.Date(2016, time.June, 15, 11, 0, 0, 0, loc), pages[0].Items[0].StartAt)
	assert.Equal(t, time.Date(2016, time.June, 15, 11, 30, 0, 0, loc), pages[0].Items[0].EndAt)
}
