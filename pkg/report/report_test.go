package report

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/testgrid/metadata/junit"
	"github.com/elementtech/terracove/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestCreateCoverageXML(t *testing.T) {
	// Define test cases
	testCases := []struct {
		suitesRoot junit.Suites
		path       string
		expectErr  bool
	}{
		{
			suitesRoot: junit.Suites{},
			path:       "../../tests/terracove.xml",
			expectErr:  false,
		},
	}

	// Run test cases
	for _, tc := range testCases {
		err := CreateCoverageXML(tc.suitesRoot, tc.path)

		// Verify result
		if tc.expectErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.FileExists(t, tc.path)

			// Verify file content is valid XML
			file, err := os.Open(tc.path)
			assert.NoError(t, err)
			defer file.Close()

			dec := xml.NewDecoder(file)
			_, err = dec.Token()
			assert.NoError(t, err)
		}
	}
}

func TestCreateJson(t *testing.T) {
	// Define test cases
	testCases := []struct {
		suitesRoot []types.TerraformModuleStatus
		path       string
		expectErr  bool
	}{
		{
			suitesRoot: []types.TerraformModuleStatus{},
			path:       "../../tests/terracove.json",
			expectErr:  false,
		},
	}

	// Run test cases
	for _, tc := range testCases {
		err := CreateJson(tc.suitesRoot, tc.path)

		// Verify result
		if tc.expectErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.FileExists(t, tc.path)

			// Verify file content is valid JSON
			file, err := os.Open(tc.path)
			assert.NoError(t, err)
			defer file.Close()

			bytes, err := ioutil.ReadAll(file)
			assert.NoError(t, err)

			var data interface{}
			err = json.Unmarshal(bytes, &data)
			assert.NoError(t, err)
		}
	}
}

func TestCreateJunitStruct(t *testing.T) {
	// Create some sample data for testing
	terraformStatuses := []types.TerraformModuleStatus{
		{
			Path: "module1",
			Results: []types.Result{
				{
					Path:              "resource1",
					Duration:          time.Duration(1) * time.Second,
					ResourceCount:     2,
					ResourceCountDiff: 1,
					ActionDeleteCount: 0,
					ActionUpdateCount: 0,
					ActionReadCount:   0,
					ActionCreateCount: 1,
					ActionNoopCount:   0,
					Coverage:          50,
				},
			},
		},
	}

	// Call the function being tested
	junitSuites, err := CreateJunitStruct(terraformStatuses)
	if err != nil {
		t.Errorf("CreateJunitStruct returned an error: %v", err)
	}

	// Verify the result
	if len(junitSuites.Suites) != 1 {
		t.Errorf("CreateJunitStruct returned %d suites, expected 1", len(junitSuites.Suites))
	}
	suite := junitSuites.Suites[0]
	if suite.Name != "module1" {
		t.Errorf("CreateJunitStruct returned a suite with name %s, expected 'module1'", suite.Name)
	}
	if suite.Tests != 1 {
		t.Errorf("CreateJunitStruct returned a suite with %d tests, expected 1", suite.Tests)
	}
	if suite.Failures != 1 {
		t.Errorf("CreateJunitStruct returned a suite with %d failures, expected 1", suite.Failures)
	}
	result := suite.Results[0]
	if result.Name != "resource1" {
		t.Errorf("CreateJunitStruct returned a result with name %s, expected 'resource1'", result.Name)
	}
	if result.Time != 1.0 {
		t.Errorf("CreateJunitStruct returned a result with time %f, expected 1.0", result.Time)
	}
}

func TestPrettyPrinter(t *testing.T) {
	ts := junit.Suites{
		Suites: []junit.Suite{
			{
				Results: []junit.Result{
					{
						Name: "Test case 1",
						Failure: &junit.Failure{
							Message: "Assertion failed",
						},
					},
					{
						Name:    "Test case 2",
						Skipped: &junit.Skipped{},
					},
					{
						Name: "Test case 3",
						Errored: &junit.Errored{
							Message: "Unexpected error",
						},
					},
					{
						Name: "Test case 4",
					},
				},
			},
		},
	}

	PrettyPrinter(ts)
}
