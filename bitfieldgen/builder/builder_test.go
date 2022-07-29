package builder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWhiteSpace(t *testing.T) {
	for _, test := range []struct {
		name     string
		width    int
		tabs     int
		expected string
	}{
		{
			name:     "2 spaces, no tabs",
			width:    2,
			tabs:     0,
			expected: "",
		},
		{
			name:     "4 spaces, 2 tabs",
			width:    4,
			tabs:     2,
			expected: "        ",
		},
		{
			name:     "2 spaces, 3 tabs",
			width:    2,
			tabs:     3,
			expected: "      ",
		},
	} {
		assert.Equal(t, test.expected, WhiteSpace(test.width, test.tabs))
	}
}
