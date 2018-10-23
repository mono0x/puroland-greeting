package command

import (
	"net"
	"net/http"

	"github.com/pkg/errors"

	"github.com/lestrrat/go-server-starter/listener"

	"github.com/mono0x/puroland-greeting/server"
	"github.com/urfave/cli"
)

func NewServerCommand() cli.Command {
	return cli.Command{
		Name: "server",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "addr"},
		},
		Action: onServerCommand,
	}
}

func onServerCommand(c *cli.Context) error {
	var l net.Listener
	if addr := c.String("addr"); addr != "" {
		var err error
		l, err = net.Listen("tcp", addr)
		if err != nil {
			return err
		}
	} else {
		listeners, err := listener.ListenAll()
		if err != nil {
			return err
		}
		if len(listeners) == 0 {
			return errors.New("address is not specified")
		}
		l = listeners[0]
	}

	s := http.Server{Handler: server.New()}

	if err := s.Serve(l); err != nil {
		return err
	}
	return nil
}
