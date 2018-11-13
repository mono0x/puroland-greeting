package command

import (
	"os"

	"github.com/urfave/cli"
)

func NewImportCommand() cli.Command {
	return cli.Command{
		Name:   "import",
		Action: onImportCommand,
	}
}

func onImportCommand(c *cli.Context) error {
	importer, cleanup, err := injectImporter()
	if err != nil {
		return err
	}
	defer cleanup()
	return importer.Import(os.Stdin)
}
