package version

import (
	"encoding/json"
	"os"

	"github.com/object88/shipbot"
	"github.com/object88/shipbot/cmd/flags"
	"github.com/object88/shipbot/cmd/traverse"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type command struct {
	cobra.Command

	output flags.Output
}

// CreateCommand returns the version command
func CreateCommand() *cobra.Command {
	var c *command
	c = &command{
		Command: cobra.Command{
			Use:   "version",
			Short: "report the version of the tool",
			Args:  cobra.NoArgs,
			PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
				return c.Preexecute(cmd, args)
			},
			RunE: func(cmd *cobra.Command, args []string) error {
				return c.Execute(cmd, args)
			},
		},
	}

	flgs := c.Flags()

	flags.CreateOutputFlag(flgs)

	return traverse.TraverseRunHooks(&c.Command)
}

func (c *command) Preexecute(cmd *cobra.Command, args []string) error {
	var err error
	c.output, err = flags.ReadOutputFlag()
	if err != nil {
		return err
	}

	return nil
}

func (c *command) Execute(cmd *cobra.Command, args []string) error {
	var v shipbot.Version

	switch c.output {
	case flags.Text:
		os.Stdout.WriteString(v.String())
	case flags.JSON:
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		err := enc.Encode(v)
		if err != nil {
			return errors.Wrapf(err, "internal error: failed to encode version")
		}
	case flags.JSONCompact:
		enc := json.NewEncoder(os.Stdout)
		err := enc.Encode(v)
		if err != nil {
			return errors.Wrapf(err, "internal error: failed to encode version")
		}
	}

	return nil
}
