package traverse

import (
	"github.com/spf13/cobra"
)

// TraverseRunHooks modifies c's PersistentPreRun* and PersistentPostRun*
// functions (when present) so that they will search c's command chain and
// invoke the corresponding hook of the first parent that provides a hook.
// When used on every command in the chain the invocation of hooks will be
// propagated up the chain to the root command.
//
// In the case of PersistentPreRun* hooks the parent hook is invoked before the
// child hook.  In the case of PersistentPostRun* the child hook is invoked
// first.
//
// Use it in place of &cobra.Command{}, e.g.
//     command := TraverseRunHooks(&cobra.Command{
//     	PersistentPreRun: func ...,
//     })
func TraverseRunHooks(c *cobra.Command) *cobra.Command {
	preRunE := c.PersistentPreRunE
	preRun := c.PersistentPreRun
	if preRunE != nil || preRun != nil {
		c.PersistentPreRun = nil
		c.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
			for p := c.Parent(); p != nil; p = p.Parent() {
				if p.PersistentPreRunE != nil {
					if err := p.PersistentPreRunE(cmd, args); err != nil {
						return err
					}
					break
				} else if p.PersistentPreRun != nil {
					p.PersistentPreRun(cmd, args)
					break
				}
			}

			if preRunE != nil {
				return preRunE(cmd, args)
			} else if preRun != nil {
				preRun(cmd, args)
			}

			return nil
		}
	}

	postRunE := c.PersistentPostRunE
	postRun := c.PersistentPostRun
	if postRunE != nil || postRun != nil {
		c.PersistentPostRun = nil
		c.PersistentPostRunE = func(cmd *cobra.Command, args []string) error {
			if postRunE != nil {
				if err := postRunE(cmd, args); err != nil {
					return err
				}
			} else if postRun != nil {
				postRun(cmd, args)
			}

			for p := c.Parent(); p != nil; p = p.Parent() {
				if p.PersistentPostRunE != nil {
					if err := p.PersistentPostRunE(cmd, args); err != nil {
						return err
					}
					break
				} else if p.PersistentPostRun != nil {
					p.PersistentPostRun(cmd, args)
					break
				}
			}

			return nil
		}
	}

	return c
}
