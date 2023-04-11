package types

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTypes(t *testing.T) {
	t.Run("TestReport", func(t *testing.T) {
		tr := TestReport{
			Name:     "test-report",
			Tests:    1,
			Failures: 1,
			Skipped:  0,
			Errors:   0,
			TestCases: []TestCase{
				{
					Name:    "test-case",
					Status:  "F",
					Message: "expected failure",
				},
			},
		}

		assert.Equal(t, "test-report", tr.Name)
		assert.Equal(t, 1, tr.Tests)
		assert.Equal(t, 1, tr.Failures)
		assert.Equal(t, 0, tr.Skipped)
		assert.Equal(t, 0, tr.Errors)
		assert.Len(t, tr.TestCases, 1)
		assert.Equal(t, "test-case", tr.TestCases[0].Name)
		assert.Equal(t, "F", tr.TestCases[0].Status)
		assert.Equal(t, "expected failure", tr.TestCases[0].Message)
	})

	t.Run("Result", func(t *testing.T) {
		r := Result{
			Path:                "/path/to/module",
			Error:               "",
			ResourceCount:       10,
			ResourceCountExists: 5,
			ResourceCountDiff:   3,
			Coverage:            60.0,
			Duration:            time.Second,
			// RawPlan:             tfjson.Plan{},
			ActionNoopCount:   1,
			ActionCreateCount: 2,
			ActionReadCount:   3,
			ActionUpdateCount: 4,
			ActionDeleteCount: 5,
		}

		assert.Equal(t, "/path/to/module", r.Path)
		assert.Nil(t, r.Error)
		assert.Equal(t, uint(10), r.ResourceCount)
		assert.Equal(t, uint(5), r.ResourceCountExists)
		assert.Equal(t, uint(3), r.ResourceCountDiff)
		assert.Equal(t, 60.0, r.Coverage)
		assert.Equal(t, time.Second, r.Duration)
		// assert.Equal(t, tfjson.Plan{}, r.RawPlan)
		assert.Equal(t, uint(1), r.ActionNoopCount)
		assert.Equal(t, uint(2), r.ActionCreateCount)
		assert.Equal(t, uint(3), r.ActionReadCount)
		assert.Equal(t, uint(4), r.ActionUpdateCount)
		assert.Equal(t, uint(5), r.ActionDeleteCount)
	})
}
