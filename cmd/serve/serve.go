package serve

import (
	"sync"

	"github.com/object88/shipbot"
	"github.com/object88/shipbot/cmd/common"
	"github.com/object88/shipbot/cmd/flags"
	"github.com/object88/shipbot/cmd/traverse"
	"github.com/object88/shipbot/serve"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type command struct {
	cobra.Command
	*common.CommonArgs

	b *shipbot.Bot

	s *serve.Server
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

	flgs.String(flags.SlackTokenKey, "", "If provided, start a connection to Slack")
	viper.BindPFlag(flags.SlackTokenKey, flgs.Lookup(flags.SlackTokenKey))
	viper.BindEnv(flags.SlackTokenKey)

	return traverse.TraverseRunHooks(&c.Command)
}

func (c *command) Preexecute(cmd *cobra.Command, args []string) error {
	c.b = &shipbot.Bot{
		Log:        c.Logger,
		Slacktoken: viper.GetString(flags.SlackTokenKey),
	}

	c.s = serve.New()

	return nil
}

func (c *command) Execute(cmd *cobra.Command, args []string) error {
	var wg sync.WaitGroup
	wg.Add(2)

	c.Logger.Infof("Starting\n")

	go func() {
		defer wg.Done()

		c.Logger.Infof("gRPC listener: started\n")
		c.s.Run()
		c.Logger.Infof("gRPC listener: complete\n")
	}()

	go func() {
		defer wg.Done()

		c.Logger.Infof("Slackbot listener: started\n")
		c.b.Run()
		c.Logger.Infof("Slackbot listener: complete\n")
	}()

	c.Logger.Infof("Started\n")

	wg.Wait()

	return nil
}
