package types

import (
	"time"

	tfjson "github.com/hashicorp/terraform-json"
)

type OutputOptions struct {
	Json bool
	// Yaml         bool
	Junit       bool
	JsonOutPath string
	// YamlOutPath  string
	JunitOutPath string
}

type ValidateOptions struct {
	ValidateTerraformBy  string
	ValidateTerragruntBy string
}

type TestReport struct {
	Name      string
	Tests     int
	Failures  int
	Skipped   int
	Errors    int
	TestCases []TestCase
}

type TestCase struct {
	Name    string
	Status  string // "P" for passed, "F" for failed, "U" for untested
	Message string // error message if failed or skipped
}

type Result struct {
	Path                string
	Error               error
	ResourceCount       uint
	ResourceCountExists uint
	ResourceCountDiff   uint
	Coverage            float64
	Duration            time.Duration
	RawPlan             tfjson.Plan
	ActionNoopCount     uint
	ActionCreateCount   uint
	ActionReadCount     uint
	ActionUpdateCount   uint
	ActionDeleteCount   uint
}

type TerraformModuleStatus struct {
	Timestamp string
	Path      string
	Results   []Result
	Coverage  float64
}
