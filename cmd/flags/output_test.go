package flags

import "testing"

func Test_Output_Default(t *testing.T) {
	var o Output
	if o != Text {
		t.Fatalf("Uninitialized value is not text")
	}
}

func Test_Output_Unmarshal(t *testing.T) {
	tcs := []struct {
		name     string
		input    string
		expected Output
	}{
		{
			name:     "json",
			input:    "json",
			expected: JSON,
		},
		{
			name:     "compact json",
			input:    "json-compact",
			expected: JSONCompact,
		},
		{
			name:     "invalid",
			input:    "foo",
			expected: Unknown,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			var o Output
			err := o.UnmarshalText([]byte(tc.input))
			if err != nil && tc.expected != Unknown {
				// We have an err returned, but we were expecting a valid result.
				t.Errorf("Unexpected error from UnmarshalText:\n%s", err.Error())
			} else if o != tc.expected {
				// We did not get the expected result; either
				// 1) there was an error AND o was set to something other than Unknown
				// 2) the decoding resulted in the wrong value
				t.Errorf("Incorrect decoding; expected '%s', actual '%s'", tc.expected, o)
			}
		})
	}
}
