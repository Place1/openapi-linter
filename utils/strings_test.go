package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStripIndent(t *testing.T) {
	require := require.New(t)
	result := StripIndent(`
		Hello
			World
	`)
	require.Equal(result, "Hello\n\tWorld")
}
