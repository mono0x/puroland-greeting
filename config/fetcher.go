package config

import (
	"net/http"

	"github.com/mono0x/puroland-greeting/scraping"
)

func NewFetcher(client *http.Client, parser scraping.Parser) scraping.Fetcher {
	return scraping.NewFetcher(client, parser, scraping.BaseURL)
}
