package report

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"text/tabwriter"

	"github.com/GoogleCloudPlatform/testgrid/metadata/junit"
	"github.com/ahmetb/go-linq/v3"
	"github.com/jatalocks/terracove/internal/types"
)

func CreateCoverageXML(suitesRoot junit.Suites, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	enc := xml.NewEncoder(file)
	enc.Indent("", "\t")
	if err := enc.Encode(suitesRoot); err != nil {
		return err
	}
	return nil
}
func CreateJson(suitesRoot []types.TerraformModuleStatus, path string) error {

	file, err := json.MarshalIndent(suitesRoot, "", " ")
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile(path, file, 0644)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

// func CreateYaml(suitesRoot []types.TerraformModuleStatus, path string) error {

// 	file, err := yaml.Marshal(suitesRoot)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	err = ioutil.WriteFile(path, file, 0644)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	return nil
// }

func CreateJunitStruct(terraformStatuses []types.TerraformModuleStatus) (junit.Suites, error) {
	suitesRoot := junit.Suites{}
	for _, ts := range terraformStatuses {
		suites := junit.Suite{
			Name:     ts.Path,
			Tests:    len(ts.Results),
			Failures: linq.From(ts.Results).WhereT(func(r types.Result) bool { return r.ResourceCountDiff > 0 || r.Error != nil }).Count(),
			Time:     linq.From(ts.Results).SelectT(func(r types.Result) float64 { return r.Duration.Seconds() }).Max().(float64),
		}

		for _, r := range ts.Results {
			// rawPlanBytes, _ := json.Marshal(r.RawPlan)

			testCase := junit.Result{
				Name: r.Path,
				Time: r.Duration.Seconds(),
			}
			testCase.SetProperty("total", fmt.Sprint(r.ResourceCount))
			testCase.SetProperty("diff", fmt.Sprint(r.ResourceCountDiff))
			testCase.SetProperty("delete", fmt.Sprint(r.ActionDeleteCount))
			testCase.SetProperty("update", fmt.Sprint(r.ActionUpdateCount))
			testCase.SetProperty("read", fmt.Sprint(r.ActionReadCount))
			testCase.SetProperty("create", fmt.Sprint(r.ActionCreateCount))
			testCase.SetProperty("noop", fmt.Sprint(r.ActionNoopCount))
			testCase.SetProperty("coverage", fmt.Sprint(r.Coverage))

			switch {
			case r.Error != nil:
				testCase.Errored = &junit.Errored{Message: "module has planning error"}
				// *testCase.Error = fmt.Sprint(rawPlanBytes)
			case r.ResourceCount == 0:
				testCase.Skipped = &junit.Skipped{Message: "module does not contain any resources"}
			case r.Coverage != 100:
				testCase.Failure = &junit.Failure{Message: fmt.Sprintf("module has %v resources with diff", r.ResourceCountDiff)}
			}
			suites.Results = append(suites.Results, testCase)
		}
		suitesRoot.Suites = append(suitesRoot.Suites, suites)
	}

	return suitesRoot, nil
}

func PrettyPrinter(testsuites junit.Suites) {
	// Build the test report
	var report types.TestReport
	report.Name = "\nTerraform Diff Report"
	report.Tests = 0
	report.Failures = 0
	report.Errors = 0
	report.Skipped = 0

	for _, suite := range testsuites.Suites {
		for _, tcase := range suite.Results {
			var tc types.TestCase
			tc.Name = tcase.Name
			if tcase.Failure != nil {
				tc.Status = "⚠️"
				tc.Message = tcase.Failure.Message
				report.Failures++
			} else if tcase.Skipped != nil {
				tc.Status = "?"
				tc.Message = "Skipped"
				report.Skipped++
			} else if tcase.Errored != nil {
				tc.Status = "X"
				tc.Message = tcase.Errored.Message
				report.Errors++
			} else {
				tc.Status = "√"
				tc.Message = "Success"
			}
			report.TestCases = append(report.TestCases, tc)
			report.Tests++
		}
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)

	// Print the test report
	fmt.Fprintln(w, report.Name)
	fmt.Fprintln(w, "============================================================================")

	for _, tc := range report.TestCases {
		fmt.Fprintf(w, "%s\t\t%s\t\t\t%s\n", tc.Name, tc.Status, tc.Message)
	}
	w.Flush()

	fmt.Println("----------------------------------------------------------------------------")
	fmt.Printf("Test Results:\nFailed       ⚠️%8d\nPassed       √%8d\nSkipped      ?%8d\nErrors       X%8d\n",
		report.Failures, report.Tests-report.Failures-report.Skipped-report.Errors, report.Skipped, report.Errors)
}
