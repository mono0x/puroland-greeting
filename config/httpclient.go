package config

import "net/http"

func NewHttpClient() *http.Client {
	return http.DefaultClient
}
