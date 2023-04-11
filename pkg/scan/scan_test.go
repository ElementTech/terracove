package scan

import (
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"testing"

	"github.com/jatalocks/terracove/internal/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

	assert.Empty(t, nil)
}

func TestCheckModuleType(t *testing.T) {
	// Create temporary directory for testing
	dir := t.TempDir()

	// Create test files
	terrFile := filepath.Join(dir, "terragrunt.hcl")
	terraformFile := filepath.Join(dir, "main.tf")
	os.Create(terrFile)
	os.Create(terraformFile)

	// Define ValidateOptions
	opts := types.ValidateOptions{
		ValidateTerragruntBy: "terragrunt.hcl",
		ValidateTerraformBy:  "main.tf",
	}

	// Test for Terragrunt module
	if moduleType := checkModuleType(dir, opts); moduleType != "terragrunt" {
		t.Errorf("Expected module type 'terragrunt', but got '%s'", moduleType)
	}

	// Test for Terraform module
	os.Remove(terrFile)
	if moduleType := checkModuleType(dir, opts); moduleType != "terraform" {
		t.Errorf("Expected module type 'terraform', but got '%s'", moduleType)
	}

	// Test for non-Terraform/Terragrunt module
	os.Remove(terraformFile)
	if moduleType := checkModuleType(dir, opts); moduleType != "" {
		t.Errorf("Expected empty module type, but got '%s'", moduleType)
	}
}

func TestFlatten(t *testing.T) {
	// Test case 1
	input1 := [][]int{{1, 2, 3}, {4, 5}, {6, 7, 8, 9}}
	expected1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if res := Flatten(input1); !equal(res, expected1) {
		t.Errorf("Flatten(%v) = %v; expected %v", input1, res, expected1)
	}

	// Test case 2
	input2 := [][]string{{"foo", "bar"}, {"baz", "qux", "quux"}}
	expected2 := []string{"foo", "bar", "baz", "qux", "quux"}
	if res := Flatten(input2); !equal(res, expected2) {
		t.Errorf("Flatten(%v) = %v; expected %v", input2, res, expected2)
	}
}

// Helper function to check if two slices are equal
func equal[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestPercentage(t *testing.T) {
	// Test case 1: denominator is zero
	res1 := percentage(50.0, 0.0)
	require.Equal(t, float64(100), res1)

	// Test case 2: numerator is zero
	res2 := percentage(0.0, 10.0)
	require.Equal(t, float64(0), res2)

	// Test case 3: normal case
	res3 := percentage(6.0, 10.0)
	require.Equal(t, float64(60), res3)

	// Test case 4: rounding to two decimal places
	res4 := percentage(5.0, 6.0)
	require.Equal(t, float64(83.33), math.Round(res4*100)/100)
}

func TestAveragePercentage(t *testing.T) {
	// Test case 1: empty slice
	res1 := averagePercentage([]types.Result{})
	require.Equal(t, float64(0), res1)

	// Test case 2: normal case
	results := []types.Result{
		{Coverage: 50},
		{Coverage: 75},
		{Coverage: 80},
	}
	res2 := averagePercentage(results)
	require.Equal(t, float64(68.33), math.Round(res2*100)/100)
}

func TestTerraformModulesTerratest(t *testing.T) {
	// Define some example input values
	paths := []string{"../../examples"}
	outputOptions := types.OutputOptions{Junit: true, HTML: false, JunitOutPath: "../../test.xml", Json: true, JsonOutPath: "../../test.json", HTMLOutPath: "../../test.html"}
	validateOptions := types.ValidateOptions{
		ValidateTerragruntBy: "terragrunt.hcl",
		ValidateTerraformBy:  "main.tf",
	}

	recursiveOptions := types.RecursiveOptions{Exclude: []string{}}

	// Call the function being tested
	result := TerraformModulesTerratest(paths, outputOptions, validateOptions, recursiveOptions)
	os.Remove("../../test.xml")
	os.Remove("../../test.json")
	// os.Remove("../../test.html")
	// Check that the actual result matches the expected output
	assert.Nil(t, result)
}
