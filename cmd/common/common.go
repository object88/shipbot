package common

import (
	cmdflags "github.com/object88/shipbot/cmd/flags"
	"github.com/object88/shipbot/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type CommonArgs struct {
	Logger *log.Log
}

func NewCommonArgs() *CommonArgs {
	return &CommonArgs{}
}

func (ca *CommonArgs) Setup(flags *pflag.FlagSet) {
	flags.BoolP(cmdflags.VerboseKey, "v", false, "Emit debug messages")
	viper.BindPFlag(cmdflags.VerboseKey, flags.Lookup(cmdflags.VerboseKey))
}

func (ca *CommonArgs) Evaluate() error {
	verbose := viper.GetBool(cmdflags.VerboseKey)
	if verbose {
		ca.Logger = log.Stderr()
	}

	return nil
}
