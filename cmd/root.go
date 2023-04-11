package cmd

import (
	"fmt"

	"github.com/jatalocks/terracove/internal/types"
	"github.com/jatalocks/terracove/pkg/scan"
	"github.com/spf13/cobra"
)

var OutputOptions types.OutputOptions
var ValidateOptions types.ValidateOptions
var RecursiveOptions types.RecursiveOptions

func newRootCmd(version string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "terracove [paths]...",
		Short: "terracove tests a directory tree for terraform/terragrunt diffs",
		Long: `terracove provides a recursive way to test the health and validity
of a terraform/terragrunt repository structues.
It plans all modules in parallel and outputs a report
in one of more of the following formats: junit or json.`,
		Version: version,
		Args:    cobra.MinimumNArgs(1),
		RunE:    run,
	}

	cmd.Flags().BoolVarP(&OutputOptions.Json, "json", "j", false, "Output JSON")
	// cmd.Flags().BoolVarP(&OutputOptions.Yaml, "yaml", "y", false, "Output YAML")
	cmd.Flags().BoolVarP(&OutputOptions.Junit, "junit", "x", false, "Output Junit XML")
	cmd.Flags().BoolVarP(&OutputOptions.HTML, "html", "w", false, "Output HTML Report")
	cmd.Flags().StringVar(&OutputOptions.JsonOutPath, "o-json", "terracove.json", "Output JSON File")
	cmd.Flags().BoolVar(&OutputOptions.Minimal, "minimal", false, "Don't Append Raw/JSON Plan to the Exported Output")
	// cmd.Flags().StringVar(&OutputOptions.YamlOutPath, "o-yaml", "terracove.yaml", "Output YAML")
	cmd.Flags().StringVar(&OutputOptions.JunitOutPath, "o-junit", "terracove.xml", "Output Junit XML File")
	cmd.Flags().StringVar(&OutputOptions.HTMLOutPath, "o-html", "terracove.html", "Output HTML Report File")
	cmd.Flags().StringSliceVarP(&RecursiveOptions.Exclude, "exclude", "e", []string{}, "Exclude directories while parsing tree")
	cmd.Flags().StringVarP(&ValidateOptions.ValidateTerraformBy, "validate-tf-by", "t", "main.tf", "validate terraform by the existence of [filename] in a directory")
	cmd.Flags().StringVarP(&ValidateOptions.ValidateTerragruntBy, "validate-tg-by", "g", "terragrunt.hcl", "validate terragrunt by the existence of [filename] in a directory")
	return cmd
}

// Execute invokes the command.
func Execute(version string, testing bool) error {
	if err := newRootCmd(version).Execute(); err != nil {
		if testing {
			return nil
		} else {
			return fmt.Errorf("error executing root command: %w", err)
		}
	}

	return nil
}

func run(cmd *cobra.Command, args []string) error {
	scan.TerraformModulesTerratest(args, OutputOptions, ValidateOptions, RecursiveOptions)
	return nil
}
