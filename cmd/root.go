package cmd

import (
	"github.com/devuo/go-app/service"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// NewRootCommand creates the root app command
func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:              "app",
		Short:            "",
		Long:             "",
		PersistentPreRun: rootPersistentPreRunFunc,
	}

	// Build the cli flags from the list of service parameters
	flags := cmd.PersistentFlags()
	for _, param := range service.Parameters().List {
		flags.String(param.Key, param.Default, param.Description)
	}

	return cmd
}

// rootPersistentPreRunFunc ensures all application conf is loaded
// and the services are properly initialized before any command is
// run
func rootPersistentPreRunFunc(cmd *cobra.Command, args []string) {
	params := service.Parameters()

	// Set the parameter values from the CLI flags
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		params.Get(f.Name).Set(f.Value.String())
	})

	service.Init(params)
}
