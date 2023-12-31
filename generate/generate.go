package generate

import (
	"bytes"
	"embed"
	"go/format"
	"text/template"
)

//go:embed templates/types.tpl
var templates embed.FS
var types *template.Template = template.Must(
	template.ParseFS(templates, "templates/types.tpl"),
)

// Config defines template configuration.
type Config struct {
	GenerateCmd  string
	Package      string
	Singular     string
	Plural       string
	Source       string
	OptionPrefix string
	Options      []Option
	WithBuilder  bool
	NoPkgOptions bool
}

// Option defines a functional option to generate.
type Option struct {
	IsSlice bool
	Name    string
	Type    string
	Field   string
}

// Generate executes the option generation template
// and return sthe resulting bytes or an error.
func Generate(cfg Config) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := types.Execute(buf, cfg)
	if err != nil {
		return buf.Bytes(), err
	}

	return format.Source(buf.Bytes())
}
