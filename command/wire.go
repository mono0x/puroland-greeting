// +build wireinject

package command

import (
	"net/http"

	"github.com/google/go-cloud/wire"
	"github.com/mono0x/puroland-greeting/server"
)

func injectHandler() (http.Handler, error) {
	wire.Build(
		server.NewHandler,
	)
	return nil, nil
}
