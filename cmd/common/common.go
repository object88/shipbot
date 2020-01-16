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
	flags.Bool(cmdflags.DebugKey, false, "Emit debug messages; implies verbiose")
	viper.BindPFlag(cmdflags.DebugKey, flags.Lookup(cmdflags.DebugKey))

	flags.Bool(cmdflags.VerboseKey, false, "Emit verbose messages")
	viper.BindPFlag(cmdflags.VerboseKey, flags.Lookup(cmdflags.VerboseKey))
}

func (ca *CommonArgs) Evaluate() error {
	ca.Logger = log.Stderr()
	if viper.GetBool(cmdflags.VerboseKey) {
		ca.Logger.SetLevel(log.Verbose)
	}
	if viper.GetBool(cmdflags.DebugKey) {
		ca.Logger.SetLevel(log.Debug)
	}

	return nil
}
