package testdata

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
