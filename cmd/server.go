package cmd

import (
	"github.com/devuo/go-app/entity/indicator"
	"github.com/devuo/go-app/entity/player"
	"github.com/devuo/go-app/service"
	"github.com/spf13/cobra"
)

// NewServerCommand creates the web server command
func NewServerCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "server",
		Short: "Starts the application web server",
		Run:   serverCommandFunc,
	}
}

func serverCommandFunc(*cobra.Command, []string) {
	service.DB.Connect()

	service.Web.Route(
		indicator.Controller{},
		player.Controller{},
	)

	service.Web.Run()
}
