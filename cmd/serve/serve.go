package serve

import (
	"github.com/object88/shipbot"
	"github.com/object88/shipbot/cmd/common"
	"github.com/object88/shipbot/cmd/flags"
	"github.com/object88/shipbot/cmd/traverse"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type command struct {
	cobra.Command
	*common.CommonArgs

	b *shipbot.Bot
}

// CreateCommand returns the version command
func CreateCommand(ca *common.CommonArgs) *cobra.Command {
	var c *command
	c = &command{
		Command: cobra.Command{
			Use:  "serve",
			Args: cobra.NoArgs,
			PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
				return c.Preexecute(cmd, args)
			},
			RunE: func(cmd *cobra.Command, args []string) error {
				return c.Execute(cmd, args)
			},
		},
		CommonArgs: ca,
	}

	flgs := c.Flags()

	flgs.String(flags.SlackTokenKey, "", "Slack token")
	viper.BindPFlag(flags.SlackTokenKey, flgs.Lookup(flags.SlackTokenKey))
	viper.BindEnv(flags.SlackTokenKey)

	return traverse.TraverseRunHooks(&c.Command)
}

func (c *command) Preexecute(cmd *cobra.Command, args []string) error {
	token := viper.GetString(flags.SlackTokenKey)

	c.b = &shipbot.Bot{
		Slacktoken: token,
	}

	return nil
}

func (c *command) Execute(cmd *cobra.Command, args []string) error {
	c.b.Run()
	return nil
}
