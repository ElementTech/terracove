package report

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"os"
	"testing"

	"github.com/GoogleCloudPlatform/testgrid/metadata/junit"
	"github.com/jatalocks/terracove/internal/types"
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
