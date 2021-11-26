package command

import (
	"log"
	"net/http"

	"github.com/rizkazn/gorestfull/routers"
	"github.com/spf13/cobra"
)

var serveCommand = &cobra.Command{
	Use:   "serve",
	Short: "Start API Server",
	RunE:  serve,
}

func serve(cmd *cobra.Command, args []string) error {
	mainRutes := routers.New()

	log.Println("Serve API on http://localhost:8080")
	if err := http.ListenAndServe(":8080", mainRutes); err != nil {
		return err
	}

	return nil
}
