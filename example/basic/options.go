// Code generated by "fopt go.foss.tools/generate/fopt/example/basic Config -w options.go"
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
func Name(value string) Option {
	return func(s *Config) {
		s.Name = value
	}
}

// Value will set the Value option for Config.
func Value(value string) Option {
	return func(s *Config) {
		s.Value = value
	}
}

// Names will set the Names option for Config.
func Names(values ...string) Option {
	return func(s *Config) {
		s.Names = append(s.Names, values...)
	}
}

// WithCustom will set the Custom option for Config.
func WithCustom(value Custom) Option {
	return func(s *Config) {
		s.Custom = value
	}
}

// Many will set the Many option for Config.
func Many(values ...Custom) Option {
	return func(s *Config) {
		s.Many = append(s.Many, values...)
	}
}
