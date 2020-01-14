package query

import (
	"fmt"
	"os"

	"github.com/object88/shipbot/cmd/common"
	"github.com/object88/shipbot/cmd/traverse"
	"github.com/object88/shipbot/query"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

type command struct {
	cobra.Command
	*common.CommonArgs

	cflags *genericclioptions.ConfigFlags

	q query.Querier

	actionConfig *action.Configuration
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

	flgs := c.Flags()

	c.cflags = genericclioptions.NewConfigFlags(false)
	c.cflags.AddFlags(flgs)

	return traverse.TraverseRunHooks(&c.Command)
}

func (c *command) Preexecute(cmd *cobra.Command, args []string) error {
	opts := []query.Option{
		query.SetLogger(c.Logger),
	}

	var err error
	c.q, err = query.New(opts...)
	if err != nil {
		return errors.Wrapf(err, "Failed to instantiate querier")
	}

	helmDriver := os.Getenv("HELM_DRIVER")
	settings := cli.New()

	c.actionConfig = &action.Configuration{}
	err = c.actionConfig.Init(settings.RESTClientGetter(), settings.Namespace(), helmDriver, c.Logger.Infof)
	if err != nil {
		return errors.Wrapf(err, "Failed to init actionConfig")
	}

	return nil
}

func (c *command) Execute(cmd *cobra.Command, args []string) error {
	name := args[0]
	releases, err := c.actionConfig.Releases.History(name)
	if err != nil {
		return errors.Wrapf(err, "Failed to get chart foo")
	}

	if len(releases) == 0 {
		return nil
	}

	r := releases[0]
	fmt.Printf("Chart: %s  %s", r.Chart.Metadata.Name, r.Chart.Metadata.Version)
	// for _, d := range r.Chart.Dependencies() {

	// }

	return nil
}
