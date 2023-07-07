package parser

import (
	"github.com/stretchr/testify/require"
	"go.foss.tools/generate/fopt/generate"
	"testing"
)

func TestParseAndPopulate(t *testing.T) {
	typ, p, err := Parse("go.foss.tools/generate/fopt/parser/testdata")
	require.NoError(t, err)
	require.NotEmpty(t, p)
	require.NotEmpty(t, typ)

	input := &generate.Config{
		Source: "Config",
	}

	cfg, err := Populate(typ, p, input)
	require.NoError(t, err)
	require.NotEmpty(t, cfg)
	require.Equal(t, cfg, *input)

	require.NotEmpty(t, cfg.Options)
	require.Len(t, cfg.Options, 5)
}
