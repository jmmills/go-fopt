/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"io"
	"os"
	"path"
	"strings"

	pluralize "github.com/gertd/go-pluralize"
	"github.com/spf13/cobra"

	"go.foss.tools/generate/fopt/generate"
	"go.foss.tools/generate/fopt/parser"
)

const (
	DefaultOptionTypeName = "Option"
)

var (
	rootCmd = &cobra.Command{
		Use:   "fopt <package> <type name>",
		Short: "Generate functional options",
		Args:  cobra.ExactArgs(2),
		RunE:  RunE,
	}
	optionName   string = DefaultOptionTypeName
	optionPrefix string
	outputFile   string
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&optionName, "option", "o", DefaultOptionTypeName, "defines the name of the option type")
	rootCmd.Flags().StringVarP(&optionPrefix, "prefix", "p", "", "defines a prefix to prepend to option function names")
	rootCmd.Flags().StringVarP(&outputFile, "write", "w", "", "defines file to write rather than stdout")
}

func RunE(cmd *cobra.Command, args []string) error {
	nodes, pack, err := parser.Parse(args[0])
	if err != nil {
		return err
	}

	var runCmd string

	if len(os.Args) > 0 {
		runCmd = strings.Join(append([]string{path.Base(os.Args[0])}, os.Args[1:]...), " ")
	}

	cfg := generate.Config{
		Source:       args[1],
		Singular:     optionName,
		OptionPrefix: optionPrefix,
		GenerateCmd:  runCmd,
		Plural:       pluralize.NewClient().Plural(optionName),
	}

	cfg, err = parser.Populate(nodes, pack, &cfg)
	if err != nil {
		return err
	}

	b, err := generate.Generate(cfg)
	if err != nil {
		return err
	}

	var writer io.WriteCloser = os.Stdout

	if outputFile != "" {
		f, err := os.OpenFile(outputFile, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0660)
		if err != nil {
			return err
		}
		writer = f
	}

	_, err = writer.Write(b)
	if err != nil {
		return err
	}

	return writer.Close()
}
