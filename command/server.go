package command

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mono0x/puroland-greeting/server"

	"github.com/pkg/errors"

	"github.com/lestrrat/go-server-starter/listener"

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

	go func() {
		if err := s.Serve(l); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM, os.Interrupt)
	<-signalChan

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.Shutdown(ctx)
}
