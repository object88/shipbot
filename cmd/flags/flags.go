package flags

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	// OutputKey determines the output format
	OutputKey = "output"

	// SlackTokenKey specifies the slack token
	SlackTokenKey = "slack-token"

	// VerboseKey turns on verbose output to STDERR
	VerboseKey = "verbose"
)

// CreateOutputFlag adds the `--output` flag to the flagset
func CreateOutputFlag(flgs *pflag.FlagSet) {
	var def Output
	flgs.String(OutputKey, def.String(), Values())
	flg := flgs.Lookup(OutputKey)
	viper.BindPFlag(OutputKey, flg)
	viper.BindEnv(OutputKey)
}

// ReadOutputFlag gets the specified output setting, and verifies that it is a
// legitimate value
func ReadOutputFlag() (Output, error) {
	raw := viper.GetString(OutputKey)
	var o Output
	if err := o.UnmarshalText([]byte(raw)); err != nil {
		return Unknown, err
	}
	return o, nil
}
