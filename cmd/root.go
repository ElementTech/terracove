package cmd

import (
	"fmt"

	"github.com/jatalocks/terracove/internal/types"
	"github.com/jatalocks/terracove/pkg/scan"
	"github.com/spf13/cobra"
)

var (
	OutputOptions    types.OutputOptions
	ValidateOptions  types.ValidateOptions
	RecursiveOptions types.RecursiveOptions
)

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

	cmd.Flags().BoolVarP(&OutputOptions.JSON, "json", "j", false, "Output JSON")
	// cmd.Flags().BoolVarP(&OutputOptions.Yaml, "yaml", "y", false, "Output YAML")
	cmd.Flags().BoolVarP(&OutputOptions.Junit, "junit", "x", false, "Output Junit XML")
	cmd.Flags().StringVar(&OutputOptions.JSONOutPath, "o-json", "terracove.json", "Output JSON")
	// cmd.Flags().StringVar(&OutputOptions.YamlOutPath, "o-yaml", "terracove.yaml", "Output YAML")
	cmd.Flags().StringVar(&OutputOptions.JunitOutPath, "o-junit", "terracove.xml", "Output Junit XML")
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
		}

		return fmt.Errorf("error executing root command: %w", err)
	}

	return nil
}

func run(cmd *cobra.Command, args []string) error {
	scan.TerraformModulesTerratest(args, OutputOptions, ValidateOptions, RecursiveOptions)

	return nil
}
