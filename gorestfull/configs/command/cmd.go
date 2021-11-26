package command

import (
	"github.com/rizkazn/gorestfull/configs/db"
	"github.com/spf13/cobra"
)

var initCommand = &cobra.Command{
	Use:   "gorestfull",
	Short: "POS App RestFull API with Golang",
}

func init() {
	initCommand.AddCommand(serveCommand)
	initCommand.AddCommand(db.MigrateUp)
	initCommand.AddCommand(db.MigrateDown)
}

func Run(args []string) error {
	initCommand.SetArgs(args)
	return initCommand.Execute()
}
