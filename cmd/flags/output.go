package flags

import (
	"strings"

	"github.com/pkg/errors"
)

// Output represents the output formatting
type Output int

const (
	// Order is important: Text is the default, and so must come after Unknown
	// (iota-1)

	// Unknown is an unspecified output format
	Unknown Output = iota - 1

	// Text is a human readable output format
	Text

	// JSON is JSON with spacing, indentation, etc.
	JSON

	// JSONCompact is JSON without extra whitespace
	JSONCompact

	// Yaml is YAML
	Yaml
)

const (
	jsonText        = "json"
	jsonCompactText = "json-compact"
	textText        = "text"
	unknownText     = "unknown"
	yamlText        = "yaml"
)

// Values returns a human-readable list of values for the Output type
func Values() string {
	return strings.Join([]string{jsonText, jsonCompactText, textText, yamlText}, ", ")
}

// UnmarshalText satisfies encoding.TextUnmarshaler
func (o *Output) UnmarshalText(text []byte) error {
	*o = Unknown
	switch string(text) {
	case jsonText:
		*o = JSON
	case jsonCompactText:
		*o = JSONCompact
	case textText:
		*o = Text
	case yamlText:
		*o = Yaml
	default:
		return errors.Errorf("Value '%s' is not a valid Output", text)
	}
	return nil
}

// String satisfies fmt.Stringer
func (o Output) String() string {
	switch o {
	case JSON:
		return jsonText
	case JSONCompact:
		return jsonCompactText
	case Text:
		return textText
	case Yaml:
		return yamlText
	default:
		return unknownText
	}
}
