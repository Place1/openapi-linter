package linter

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	require := require.New(t)

	config, err := LoadConfig("../linter-config.yaml")
	require.NoError(err)

	require.NotNil(config.Rules.NoEmptyDescriptions)
}

func TestUsesPreset(t *testing.T) {
	require := require.New(t)

	config, err := LoadConfig("../linter-config.yaml")
	require.NoError(err)

	require.True(config.Rules.NoEmptyOperationID)
}
