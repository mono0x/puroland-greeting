// +build wireinject

package command

import (
	"net/http"

	"github.com/mono0x/puroland-greeting/config"
	"github.com/mono0x/puroland-greeting/importer"
	"github.com/mono0x/puroland-greeting/scraping"

	"github.com/mono0x/puroland-greeting/updater"

	"github.com/google/go-cloud/wire"
	"github.com/mono0x/puroland-greeting/server"
)

func injectHandler() (http.Handler, func(), error) {
	wire.Build(
		config.NewDB,
		server.NewHandler,
	)
	return nil, nil, nil
}

func injectImporter() (importer.Importer, func(), error) {
	wire.Build(
		config.NewDB,
		importer.NewImporter,
	)
	return nil, nil, nil
}

func injectUpdater() (updater.Updater, func(), error) {
	wire.Build(
		config.NewDB,
		config.NewFetcher,
		config.NewHttpClient,
		config.NewWalker,
		scraping.NewParser,
		updater.NewUpdater,
	)
	return nil, nil, nil
}
