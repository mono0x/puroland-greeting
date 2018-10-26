package command

import "github.com/urfave/cli"

func NewUpdateCommand() cli.Command {
	return cli.Command{
		Name:   "update",
		Action: onUpdateCommand,
	}
}

func onUpdateCommand(c *cli.Context) error {
	updater, err := injectUpdater()
	if err != nil {
		return err
	}
	return updater.Update()
}
