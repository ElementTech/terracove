package scan

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/jatalocks/terracove/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestGetAllDirectories(t *testing.T) {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path) // for example /home/user
	// Test getAllDirectories function
	testDir := "../../examples"
	validateOptions := types.ValidateOptions{ValidateTerraformBy: "main.tf", ValidateTerragruntBy: "terragrunt.hcl"}
	recursiveOptions := types.RecursiveOptions{Exclude: []string{"error"}}
	subpaths := getAllDirectories([]string{testDir}, validateOptions, recursiveOptions)

	expectedResult := map[string][]string{
		testDir: {filepath.ToSlash(testDir + "/terraform/success"), filepath.ToSlash(testDir + "/terraform/tfstate-diff"), filepath.ToSlash(testDir + "/terragrunt/no-resources")},
	}

	assert.Equal(t, expectedResult, subpaths)

	assert.Nil(t, err)
}
