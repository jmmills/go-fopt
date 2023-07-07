// Code generated by "{{.GenerateCmd}}"
package {{.Package}}

// {{.Singular}} defines an functional option to populate {{.Source}}
type {{.Singular}} func(s *{{.Source}})

// {{.Plural}} defines a plural set of {{.Singular}}
type {{.Plural}} []{{.Singular}}

func (opts {{.Plural}}) apply(s *{{.Source}}) {{.Source}} {
    for _, opt := range opts {
        opt(s)
    }
    return *s
}

{{- range .Options }}
{{if .IsSlice}}
func {{$.OptionPrefix}}{{.Name}}(values ...{{.Type}}) {{$.Singular}} {
    return func(s *{{$.Source}}) {
        s.{{.Field}} = append(s.{{.Field}}, values...)
    }
}
{{else}}
func {{$.OptionPrefix}}{{.Name}}(value {{.Type}}) {{$.Singular}} {
    return func(s *{{$.Source}}) {
        s.{{.Field}} = value
    }
}
{{end}}
{{- end}}

