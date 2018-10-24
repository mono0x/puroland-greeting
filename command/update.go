package command

import "github.com/urfave/cli"

func NewUpdateCommand() cli.Command {
	return cli.Command{
		Name:   "update",
		Action: onUpdateCommand,
	}
}

func onUpdateCommand(c *cli.Context) error {
	return nil
}
