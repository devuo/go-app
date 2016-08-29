package cmd

import (
	"github.com/devuo/go-app/service"
	"github.com/spf13/cobra"
)

// NewParametersCommand dumps the current configuration parameters
func NewParametersCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "parameters",
		Short: "Dumps the current configuration parameters",
		Run:   parametersCommandFunc,
	}
}

func parametersCommandFunc(cmd *cobra.Command, args []string) {
	for _, p := range service.ParameterContainer.List {
		var flag string

		if p.Value == p.Default {
			flag = "[default]"
		}

		cmd.Printf("%s - %s %s\n", p.Key, p.Value, flag)
	}
}
