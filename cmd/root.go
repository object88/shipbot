package cmd

import (
	"strings"

	"github.com/object88/shipbot/cmd/common"
	"github.com/object88/shipbot/cmd/query"
	"github.com/object88/shipbot/cmd/serve"
	"github.com/object88/shipbot/cmd/traverse"
	"github.com/object88/shipbot/cmd/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// InitializeCommands sets up the cobra commands
func InitializeCommands() *cobra.Command {
	ca, rootCmd := createRootCommand()

	rootCmd.AddCommand(
		query.CreateCommand(ca),
		serve.CreateCommand(ca),
		version.CreateCommand(),
	)

	return rootCmd
}

func createRootCommand() (*common.CommonArgs, *cobra.Command) {
	ca := common.NewCommonArgs()

	cmd := &cobra.Command{
		Use: "shipbot",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			ca.Evaluate()
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

	flags := cmd.PersistentFlags()
	ca.Setup(flags)

	return ca, traverse.TraverseRunHooks(cmd)
}
