package linter

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	require := require.New(t)

	config, err := LoadConfig("../openapi-linter.yaml")
	require.NoError(err)

	require.NotNil(config.Rules.NoEmptyDescriptions)
}
