package cmd

import (
	"strings"

	cmdflags "github.com/object88/shipbot/cmd/flags"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// InitializeCommands sets up the cobra commands
func InitializeCommands() *cobra.Command {
	rootCmd := createRootCommand()

	return rootCmd
}

func createRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "shipbot",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
		PersistentPostRunE: func(cmd *cobra.Command, _ []string) error {
			return nil
		},
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.SetEnvPrefix("SHIPBOT")

	flags := cmd.Flags()

	flags.String(cmdflags.SlackboyKey, "", "Slack token")
	viper.BindPFlag(cmdflags.SlackboyKey, flags.Lookup(cmdflags.SlackboyKey))

	flags.BoolP(cmdflags.VerboseKey, "v", false, "Emit debug messages")
	viper.BindPFlag(cmdflags.VerboseKey, flags.Lookup(cmdflags.VerboseKey))

	return cmd
}
