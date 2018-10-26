// +build wireinject

package command

import (
	"net/http"

	"github.com/mono0x/puroland-greeting/config"
	"github.com/mono0x/puroland-greeting/scraping"

	"github.com/mono0x/puroland-greeting/updater"

	"github.com/google/go-cloud/wire"
	"github.com/mono0x/puroland-greeting/server"
)

func injectHandler() (http.Handler, error) {
	wire.Build(
		server.NewHandler,
	)
	return nil, nil
}

func injectUpdater() (updater.Updater, error) {
	wire.Build(
		config.NewHttpClient,
		scraping.NewParser,
		config.NewFetcher,
		config.NewWalker,
		updater.NewUpdater,
	)
	return nil, nil
}
