package cmd

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRootCmd(t *testing.T) {
	version := "1.0.0"
	cmd := newRootCmd(version)

	assert.Equal(t, "terracove [paths]...", cmd.Use)
	assert.Equal(t, "terracove tests a directory tree for terraform/terragrunt diffs", cmd.Short)
	assert.Contains(t, cmd.Long, "terracove provides a recursive way to test the health and validity")
	assert.Equal(t, version, cmd.Version)

	assert.False(t, OutputOptions.JSON)
	assert.False(t, OutputOptions.Junit)
	assert.Equal(t, "terracove.json", OutputOptions.JSONOutPath)
	assert.Equal(t, "terracove.xml", OutputOptions.JunitOutPath)
	assert.Equal(t, "main.tf", ValidateOptions.ValidateTerraformBy)
	assert.Equal(t, "terragrunt.hcl", ValidateOptions.ValidateTerragruntBy)
}

func TestRun(t *testing.T) {
	args := []string{"examples"}
	var stdout bytes.Buffer
	OutputOptions.JSON = true
	OutputOptions.JSONOutPath = "output.json"
	rootCmd := newRootCmd("1.0.0")
	rootCmd.SetOut(&stdout)

	err := run(rootCmd, args)

	assert.NoError(t, err)
}

func TestExecute(t *testing.T) {
	err := Execute("1.0.0", true)

	assert.NoError(t, err)
}
