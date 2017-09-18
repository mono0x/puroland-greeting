package greeting

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetSecretIndexPageURL(t *testing.T) {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		t.Fatal(err)
	}

	date := time.Date(2016, time.June, 15, 0, 0, 0, 0, loc)
	parser := &Parser{}
	assert.Equal(t, "http://www.puroland.co.jp/chara_gre/mobile/?para=20160615", parser.GetSecretIndexPageURL(date))
}

func TestParseIndexPage(t *testing.T) {
	f, err := os.Open("data/index.asp")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		t.Fatal(err)
	}

	parser := &Parser{}

	page, err := parser.ParseIndexPage(f)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "http://www.puroland.co.jp/chara_gre/mobile/chara_sentaku.asp?TCHK=2016156&lang=", page.MenuPageURL)
	assert.Equal(t, time.Date(2016, time.June, 15, 0, 0, 0, 0, loc), page.Date)
}

func TestParseMenuPage(t *testing.T) {
	f, err := os.Open("data/chara_sentaku.asp")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	parser := &Parser{}

	page, err := parser.ParseMenuPage(f)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 12, len(page.Items))
	assert.Equal(t, "キティ・ホワイト", page.Items[0].CharacterName)
	assert.Equal(t, "http://www.puroland.co.jp/chara_gre/mobile/chara_sche.asp?TCHK=2016156&C_KEY=1", page.Items[0].CharacterPageURL)
}

func TestParseCharacterPage(t *testing.T) {
	f, err := os.Open("data/chara_sche.asp")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		t.Fatal(err)
	}

	parser := &Parser{}

	page, err := parser.ParseCharacterPage(f)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, time.Date(2016, time.June, 15, 0, 0, 0, 0, loc), page.Date)
	assert.Equal(t, "キティ・ホワイト", page.Name)
	assert.Equal(t, 2, len(page.Items))
	assert.Equal(t, time.Date(2016, time.June, 15, 11, 0, 0, 0, loc), page.Items[0].StartAt)
	assert.Equal(t, time.Date(2016, time.June, 15, 11, 30, 0, 0, loc), page.Items[0].EndAt)
}

func TestParseCharacterListPage(t *testing.T) {
	f, err := os.Open("data/www.puroland.jp/greeting/schedule/index.html")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		t.Fatal(err)
	}

	parser := &Parser{}

	page, err := parser.ParseCharacterListPage(f)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, time.Date(2017, time.August, 4, 0, 0, 0, 0, loc), page.Date)
	assert.Equal(t, 21, len(page.Items))
	assert.Equal(t, "シナモン", page.Items[0].Name)
}
