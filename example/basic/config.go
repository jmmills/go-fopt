package example

//go:generate go run go.foss.tools/generate/fopt go.foss.tools/generate/fopt/example Config -w options.go

// Config is a type we are going to generate functional options for
type Config struct {
	Name   string   `fopt:"Name"`
	Value  string   `fopt:"Value"`
	Names  []string `fopt:"Names"`
	Custom Custom   `fopt:"WithCustom"`
	Many   []Custom `fopt:"Many"`
}

// Custom is a custom type in our option struct.
type Custom struct {
	Value string
}

// Interface is an example.
type Interface interface{}

type private struct {
	cfg Config
}

func New(opts ...Option) Interface {
	return private{
		cfg: Options(opts).Apply(new(Config)),
	}
}
