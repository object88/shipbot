package query

import (
	"github.com/object88/shipbot/client"
	"github.com/object88/shipbot/cmd/common"
	"github.com/object88/shipbot/cmd/traverse"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type command struct {
	cobra.Command
	*common.CommonArgs

	// cflags *genericclioptions.ConfigFlags

	c *client.Client

	// actionConfig *action.Configuration
}

// CreateCommand returns the version command
func CreateCommand(ca *common.CommonArgs) *cobra.Command {
	var c *command
	c = &command{
		Command: cobra.Command{
			Use:   "query CHART",
			Short: "request information of the specified deployment",
			Args:  cobra.ExactArgs(1),
			PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
				return c.Preexecute(cmd, args)
			},
			RunE: func(cmd *cobra.Command, args []string) error {
				return c.Execute(cmd, args)
			},
		},
		CommonArgs: ca,
	}

	// flgs := c.Flags()

	// c.cflags = genericclioptions.NewConfigFlags(false)
	// c.cflags.AddFlags(flgs)

	return traverse.TraverseRunHooks(&c.Command)
}

func (c *command) Preexecute(cmd *cobra.Command, args []string) error {
	opts := []client.Option{
		client.SetLogger(c.Logger),
	}

	var err error
	c.c, err = client.New(opts...)
	if err != nil {
		return errors.Wrapf(err, "Failed to instantiate client")
	}

	// helmDriver := os.Getenv("HELM_DRIVER")
	// settings := cli.New()

	// c.actionConfig = &action.Configuration{}
	// err = c.actionConfig.Init(settings.RESTClientGetter(), settings.Namespace(), helmDriver, c.Logger.Infof)
	// if err != nil {
	// 	return errors.Wrapf(err, "Failed to init actionConfig")
	// }

	return nil
}

func (c *command) Execute(cmd *cobra.Command, args []string) error {
	c.c.Foo()

	// name := args[0]
	// releases, err := c.actionConfig.Releases.History(name)
	// if err != nil {
	// 	return errors.Wrapf(err, "Failed to get chart foo")
	// }

	// if len(releases) == 0 {
	// 	return nil
	// }

	// r := releases[0]
	// fmt.Printf("Chart: %s  %s", r.Chart.Metadata.Name, r.Chart.Metadata.Version)
	// // for _, d := range r.Chart.Dependencies() {

	// // }

	return nil
}
