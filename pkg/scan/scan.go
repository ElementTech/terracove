package scan

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/briandowns/spinner"
	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/terraform"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/jatalocks/terracove/internal/types"
	"github.com/jatalocks/terracove/pkg/report"
)

func getAllDirectories(dirs []string, ValidateOptions types.ValidateOptions) map[string][]string {
	subpaths := make(map[string][]string)
	for _, dir := range dirs {
		filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil || !info.IsDir() || (strings.HasPrefix(info.Name(), ".") && (info.Name() != ".")) {
				fmt.Println("Skipping ", info.Name())
				return filepath.SkipDir
			}
			if moduleType := checkModuleType(path, ValidateOptions); moduleType != "" {
				subpaths[dir] = append(subpaths[dir], filepath.ToSlash(path))
			}
			return nil
		})
	}
	return subpaths
}

func checkModuleType(path string, ValidateOptions types.ValidateOptions) string {
	// Check for Terragrunt module
	if _, err := os.Stat(filepath.Join(path, ValidateOptions.ValidateTerragruntBy)); err == nil {
		return "terragrunt"
	}

	// Check for Terraform module
	if _, err := os.Stat(filepath.Join(path, ValidateOptions.ValidateTerraformBy)); err == nil {
		return "terraform"
	}

	// Module is not a Terraform or Terragrunt module
	return ""
}

func Flatten[T any](lists [][]T) []T {
	var res []T
	for _, list := range lists {
		res = append(res, list...)
	}
	return res
}

func TestTerraformModules(paths []string, OutputOptions types.OutputOptions, ValidateOptions types.ValidateOptions) map[string][]string {
	dirsMap := getAllDirectories(paths, ValidateOptions)
	timestamp := time.Now().Format(time.RFC3339)
	var statuses []types.TerraformModuleStatus
	var wg sync.WaitGroup

	for root, v := range dirsMap {
		wg.Add(1)
		go func(root string, v []string) {
			defer wg.Done()

			var results []types.Result
			var mu sync.Mutex

			var wg2 sync.WaitGroup
			for _, dir := range v {
				wg2.Add(1)
				go func(dir string) {
					defer wg2.Done()
					moduleType := checkModuleType(dir, ValidateOptions)
					if moduleType == "" {
						return
					}
					tfOptions := &terraform.Options{
						TerraformDir:    dir,
						TerraformBinary: moduleType,
						Logger:          logger.Discard,
						PlanFilePath:    ".terracove.plan",
					}
					testingContext := testing.T{}
					now := time.Now()
					spinner := spinner.New(spinner.CharSets[33], 100*time.Millisecond)
					spinner.Suffix = fmt.Sprintf(" %v: ", dir)
					spinner.Start()
					_, err := terraform.InitAndPlanE(&testingContext, tfOptions)
					res := types.Result{
						Path:     dir,
						Error:    err,
						Duration: time.Since(now),
					}
					if res.Error == nil {
						plan, err := terraform.ShowWithStructE(&testingContext, tfOptions)
						if err == nil {
							res.RawPlan = plan.RawPlan
							resourceCount := len(plan.ResourceChangesMap)
							var resourceCountExists uint
							var resourceCountDiff uint
							var actionCount = map[tfjson.Action]int{}
							for _, change := range plan.ResourceChangesMap {
								action := change.Change.Actions[0]
								if action == tfjson.ActionCreate || action == tfjson.ActionUpdate || action == tfjson.ActionDelete {
									resourceCountDiff++
								} else {
									resourceCountExists++
								}
								actionCount[action]++
							}
							res.ResourceCount = uint(resourceCount)
							res.ResourceCountExists = resourceCountExists
							res.ResourceCountDiff = resourceCountDiff
							res.ActionNoopCount = uint(actionCount[tfjson.ActionNoop])
							res.ActionCreateCount = uint(actionCount[tfjson.ActionCreate])
							res.ActionReadCount = uint(actionCount[tfjson.ActionRead])
							res.ActionUpdateCount = uint(actionCount[tfjson.ActionUpdate])
							res.ActionDeleteCount = uint(actionCount[tfjson.ActionDelete])
							res.Coverage = float64(percentage(float64(resourceCountExists), float64(resourceCount)))
						} else {
							res.Error = err
						}
					}
					mu.Lock()
					results = append(results, res)
					mu.Unlock()
				}(dir)
			}
			wg2.Wait()
			mu.Lock()
			defer mu.Unlock()
			statuses = append(statuses, types.TerraformModuleStatus{
				Path:      root,
				Results:   results,
				Timestamp: timestamp,
				Coverage:  averagePercentage(results),
			})
		}(root, v)
	}
	wg.Wait()
	junitStruct, err := report.CreateJunitStruct(statuses)
	if err != nil {
		fmt.Println(err)
	}
	if OutputOptions.Junit {

		if err := report.CreateCoverageXML(junitStruct, OutputOptions.JunitOutPath); err != nil {
			fmt.Println("Error while creating junit XML: ", err)
		} else {
			fmt.Printf("%v created succesfully\n", OutputOptions.JunitOutPath)
		}
	}
	if OutputOptions.Json {
		if err := report.CreateJson(statuses, OutputOptions.JsonOutPath); err != nil {
			fmt.Println("Error while creating JSON: ", err)
		} else {
			fmt.Printf("%v created succesfully\n", OutputOptions.JsonOutPath)
		}
	}
	// if OutputOptions.Yaml {
	// 	if err := report.CreateYaml(statuses, OutputOptions.YamlOutPath); err != nil {
	// 		fmt.Println("Error while creating YAML: ", err)
	// 	} else {
	// 		fmt.Printf("%v created succesfully\n", OutputOptions.YamlOutPath)
	// 	}
	// }
	report.PrettyPrinter(junitStruct)
	return nil
}

func percentage(num float64, denom float64) float64 {
	if denom == 0 {
		return 100
	}
	return math.Round((num/denom)*10000) / 100
}

func averagePercentage(results []types.Result) float64 {
	var percentages []float64

	for _, p := range results {
		percentages = append(percentages, p.Coverage)
	}

	sum := 0.0

	for _, p := range percentages {
		sum += p
	}

	if len(percentages) > 0 {
		return math.Round((sum/float64(len(percentages)))*100) / 100
	}

	return 0
}
