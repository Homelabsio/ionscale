package cmd

import (
	"github.com/muesli/coral"
)

func Command() *coral.Command {
	rootCmd := rootCommand()
	rootCmd.AddCommand(serverCommand())
	rootCmd.AddCommand(versionCommand())
	rootCmd.AddCommand(tailnetCommand())
	rootCmd.AddCommand(authkeysCommand())
	rootCmd.AddCommand(machineCommands())

	return rootCmd
}

func Execute() error {
	return Command().Execute()
}

func rootCommand() *coral.Command {
	return &coral.Command{
		Use: "ionscale",
	}
}
