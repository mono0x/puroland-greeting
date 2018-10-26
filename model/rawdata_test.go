package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestToRawGreetings(t *testing.T) {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		t.Fatal(err)
	}

	rawData := &RawData{
		CharacterPages: []*CharacterPage{
			{
				Name: "キティ・ホワイト",
				Items: []*CharacterPageItem{
					{
						StartAt:  time.Date(2018, time.October, 26, 10, 0, 0, 0, loc),
						FinishAt: time.Date(2018, time.October, 26, 10, 30, 0, 0, loc),
						Place:    "エンターテイメントホール(1F)",
					},
					{
						StartAt:  time.Date(2018, time.October, 26, 15, 0, 0, 0, loc),
						FinishAt: time.Date(2018, time.October, 26, 15, 30, 0, 0, loc),
						Place:    "館外(他)",
					},
				},
			},
			{
				Name: "ダニエル・スター",
				Items: []*CharacterPageItem{
					{
						StartAt:  time.Date(2018, time.October, 26, 15, 0, 0, 0, loc),
						FinishAt: time.Date(2018, time.October, 26, 15, 30, 0, 0, loc),
						Place:    "館外(他)",
					},
				},
			},
		},
	}

	rawGreetings := rawData.ToRawGreetings()

	assert.Equal(t, 3, len(rawGreetings))
	assert.Equal(t, "キティ・ホワイト", rawGreetings[0].Character)
	assert.Equal(t, "エンターテイメントホール(1F)", rawGreetings[0].Place)
	assert.Equal(t, RawGreetingTime(time.Date(2018, time.October, 26, 10, 0, 0, 0, loc)), *rawGreetings[0].StartAt)
	assert.Equal(t, RawGreetingTime(time.Date(2018, time.October, 26, 10, 30, 0, 0, loc)), *rawGreetings[0].FinishAt)
	assert.Equal(t, "キティ・ホワイト", rawGreetings[1].Character)
	assert.Equal(t, "館外(他)", rawGreetings[1].Place)
	assert.Equal(t, "ダニエル・スター", rawGreetings[2].Character)
	assert.Equal(t, "館外(他)", rawGreetings[2].Place)
}
