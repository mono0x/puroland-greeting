package config

import (
	"time"

	"github.com/mono0x/puroland-greeting/scraping"
)

func NewWalker(fetcher scraping.Fetcher) scraping.Walker {
	return scraping.NewWalker(fetcher, 500*time.Millisecond)
}
