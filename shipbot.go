package shipbot

import (
	"encoding/json"
	"strings"
)

// GitCommit contains the git SHA
var GitCommit = "unset"

// AppVersion contains the version
var AppVersion = "unset"

type Version struct{}

// AppVersion returns the binary version
func (Version) AppVersion() string {
	return AppVersion
}

// Git returns the commit SHA of the source
func (Version) Git() string {
	return GitCommit
}

func (v Version) String() string {
	var sb strings.Builder
	sb.WriteString("Version:    ")
	sb.WriteString(v.AppVersion())
	sb.WriteRune('\n')
	sb.WriteString("Git commit: ")
	sb.WriteString(v.Git())
	sb.WriteRune('\n')
	return sb.String()
}

func (v Version) MarshalJSON() ([]byte, error) {
	x := struct {
		Version string `json:"version"`
		SHA     string `json:"sha"`
	}{
		Version: AppVersion,
		SHA:     GitCommit,
	}
	return json.Marshal(x)
}
