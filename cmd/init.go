package cmd

// Init builds and executes the application cli
func Init() {
	root := NewRootCommand()
	root.AddCommand(
		NewServerCommand(),
		NewParametersCommand(),
	)
	root.Execute()
}
