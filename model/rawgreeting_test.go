package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/Songmu/go-ltsv"
)

func TestLtsvMarshal(t *testing.T) {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		t.Fatal(err)
	}

	startAt := RawGreetingTime(time.Date(2018, time.October, 26, 15, 0, 0, 0, loc))
	finishAt := RawGreetingTime(time.Date(2018, time.October, 26, 15, 30, 0, 0, loc))
	rawGreeting := &RawGreeting{
		Character: "キティ・ホワイト",
		Place:     "館外(他)",
		StartAt:   &startAt,
		FinishAt:  &finishAt,
	}

	ltsv, err := ltsv.Marshal(rawGreeting)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "character:キティ・ホワイト\tplace:館外(他)\tstart_at:2018-10-26 15:00:00 +0900\tend_at:2018-10-26 15:30:00 +0900", string(ltsv))
}

func TestLtsvUnmarshal(t *testing.T) {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		t.Fatal(err)
	}

	const data = "character:キティ・ホワイト\tplace:館外(他)\tstart_at:2018-10-26 15:00:00 +0900\tend_at:2018-10-26 15:30:00 +0900"

	var rawGreeting RawGreeting
	if err := ltsv.Unmarshal([]byte(data), &rawGreeting); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "キティ・ホワイト", rawGreeting.Character)
	assert.Equal(t, "館外(他)", rawGreeting.Place)
	assert.WithinDuration(t, time.Date(2018, time.October, 26, 15, 0, 0, 0, loc), time.Time(*rawGreeting.StartAt), 0)
	assert.WithinDuration(t, time.Date(2018, time.October, 26, 15, 30, 0, 0, loc), time.Time(*rawGreeting.FinishAt), 0)
}
