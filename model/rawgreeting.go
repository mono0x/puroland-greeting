package model

import "time"

type RawGreetingTime time.Time

const rawGreetingTimeLayout = "2006-01-02 15:04:05 -0700"

func (r *RawGreetingTime) MarshalText() ([]byte, error) {
	return []byte((*time.Time)(r).Format(rawGreetingTimeLayout)), nil
}

func (r *RawGreetingTime) UnmarshalText(b []byte) error {
	t, err := time.Parse(rawGreetingTimeLayout, string(b))
	if err != nil {
		return err
	}
	*r = RawGreetingTime(t)
	return nil
}

type RawGreeting struct {
	Character string           `ltsv:"character"`
	Venue     string           `ltsv:"place"` // for compatibility
	StartTime *RawGreetingTime `ltsv:"start_at"`
	EndTime   *RawGreetingTime `ltsv:"end_at"` // for compatibility
}
