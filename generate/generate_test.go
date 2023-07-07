package generate

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenerate(t *testing.T) {
	table := []struct {
		Name        string
		Config      Config
		ExpectError error
		Contains    string
	}{
		{
			Name:     "slice",
			Contains: "func WithNames(values ...string) Option",
			Config: Config{
				Package:      "test",
				Singular:     "Option",
				Plural:       "Options",
				Source:       "Config",
				OptionPrefix: "With",
				Options: []Option{
					{
						IsSlice: true,
						Name:    "Names",
						Type:    "string",
						Field:   "Name",
					},
				},
			},
		},
		{
			Name:     "primitive",
			Contains: "func WithName(value string) Option",
			Config: Config{
				Package:      "test",
				Singular:     "Option",
				Plural:       "Options",
				Source:       "Config",
				OptionPrefix: "With",
				Options: []Option{
					{
						IsSlice: false,
						Name:    "Name",
						Type:    "string",
						Field:   "Name",
					},
				},
			},
		},
	}

	for _, row := range table {
		t.Run(row.Name, func(t *testing.T) {
			b, err := Generate(row.Config)
			if row.ExpectError != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, row.ExpectError)
				return
			}
			require.NoError(t, err)
			require.NotEmpty(t, b)
			require.Contains(t, string(b), row.Contains)
		})
	}
}
