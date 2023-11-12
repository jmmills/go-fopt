// Code generated by "fopt go.foss.tools/generate/fopt/example/builder Config -w options.go --only-builder"
package example

// Option defines an functional option to populate Config
type Option func(s *Config)

// Options defines a plural set of Option
type Options []Option

// Apply will apply a set of options to a pointer to a given Config
func (opts Options) Apply(s *Config) Config {
	for _, opt := range opts {
		opt(s)
	}
	return *s
}

// Name will set the Name option for Config.
func (opts *Options) Name(value string) Options {
	*opts = append(*opts, Option(func(s *Config) {
		s.Name = value
	}))
	return *opts
}

// Value will set the Value option for Config.
func (opts *Options) Value(value string) Options {
	*opts = append(*opts, Option(func(s *Config) {
		s.Value = value
	}))
	return *opts
}

// Names will set the Names option for Config.
func (opts *Options) Names(values ...string) Options {
	*opts = append(*opts, Option(func(s *Config) {
		s.Names = append(s.Names, values...)
	}))
	return *opts
}

// WithCustom will set the Custom option for Config.
func (opts *Options) WithCustom(value Custom) Options {
	*opts = append(*opts, Option(func(s *Config) {
		s.Custom = value
	}))
	return *opts
}

// Many will set the Many option for Config.
func (opts *Options) Many(values ...Custom) Options {
	*opts = append(*opts, Option(func(s *Config) {
		s.Many = append(s.Many, values...)
	}))
	return *opts
}
