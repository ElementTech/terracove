package types

import (
	"time"

	tfjson "github.com/hashicorp/terraform-json"
	// tfjson "github.com/hashicorp/terraform-json"
)

type OutputOptions struct {
	Json    bool
	Minimal bool
	HTML    bool
	// Yaml         bool
	Junit       bool
	JsonOutPath string
	HTMLOutPath string
	// YamlOutPath  string
	JunitOutPath string
}

type RecursiveOptions struct {
	Exclude []string
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
	Error               string
	ResourceCount       uint
	ResourceCountExists uint
	ResourceCountDiff   uint
	Coverage            float64
	Duration            time.Duration
	PlanJSON            tfjson.Plan
	PlanRaw             string

	ActionNoopCount   uint
	ActionCreateCount uint
	ActionReadCount   uint
	ActionUpdateCount uint
	ActionDeleteCount uint
}

type TerraformModuleStatus struct {
	Timestamp string
	Path      string
	Results   []Result
	Coverage  float64
}
