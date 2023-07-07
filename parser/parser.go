package parser

import (
	"fmt"
	"go/ast"
	"go/build"

	"github.com/wzshiming/gotype"
	"go.foss.tools/generate/fopt/generate"
)

const Tag = "fopt"

func Parse(path string) (gotype.Type, *build.Package, error) {
	importer := gotype.NewImporter()
	nodes, err := importer.Import(path, "")

	if err != nil {
		return nodes, nil, fmt.Errorf("importing package: %w", err)
	}

	build, err := importer.ImportBuild(path, "")
	if err != nil {
		return nodes, build, fmt.Errorf("importing build: %w", err)
	}

	return nodes, build, nil
}

func Populate(
	nodes gotype.Type,
	pack *build.Package,
	cfg *generate.Config,
) (generate.Config, error) {
	cfg.Package = pack.Name

	for i := 0; i < nodes.NumChild(); i++ {
		node := nodes.Child(i)
		if !ast.IsExported(node.Name()) {
			continue
		}

		if node.Kind() == gotype.Struct && node.Name() == cfg.Source {
			for f := 0; f < node.NumField(); f++ {
				field := node.Field(f)
				tag := field.Tag()
				name, ok := tag.Lookup(Tag)
				if !ok {
					continue
				}

				o := generate.Option{
					IsSlice: field.Elem().Kind() == gotype.Slice,
					Name:    name,
					Type:    field.Elem().String(),
					Field:   field.Name(),
				}

				if field.Elem().Kind() == gotype.Slice {
					o.Type = field.Elem().Elem().String()
				}

				cfg.Options = append(cfg.Options, o)
			}
		}
	}

	return *cfg, nil
}
