package worker

import (
	"context"
	"fmt"
	"github.com/dgozalo/aec-remote-executor/pkg/graph/model"
	"github.com/dgozalo/aec-remote-executor/pkg/management"
	"github.com/dgozalo/aec-remote-executor/pkg/worker/compilers/java"
	"github.com/pkg/errors"
	"go.temporal.io/sdk/temporal"
	"strconv"
	"strings"
)

// Activity is the struct that contains the Temporal activity information
type Activity struct {
	management *management.ManagementService
}

// NewActivity creates a new temporal activity
func NewActivity(management *management.ManagementService) *Activity {
	return &Activity{management: management}
}

// ExecutionActivity is the activity function that will compile and execute the code
func (a Activity) ExecutionActivity(ctx context.Context, execution model.NewExecution) (*ExecutionResult, error) {
	codeTemplate, err := a.management.GetAssignmentCodeTemplateForCode(execution.AssignmentID, execution.Language)
	if err != nil {
		return nil, temporal.NewNonRetryableApplicationError(errors.Wrap(err, "there was a problem getting the assignment").Error(), "ASSIGNMENT_NOT_FOUND", nil)
	}
	fmt.Printf("Compile code %s for language %s!", execution.Code, execution.Language)
	result := &ExecutionResult{}
	if execution.Language == "java" {
		stdout, stderr, results, err := java.RunCompile(execution.Code, codeTemplate.TestRunnerCode)
		if err != nil {
			return nil, temporal.NewNonRetryableApplicationError(errors.Wrap(err, "there was a problem running the Java compilation").Error(), "COMPILE_FAILURE", nil)
		}
		result.parseTestResults(results)
		result.Stdout = stdout
		result.Stderr = stderr
		return result, nil
	}
	return result, nil
}

func (t *ExecutionResult) parseTestResults(rawTestResults string) {
	testLines := strings.Split(rawTestResults, "\n")
	testResults := make([]TestResult, 0)
	for _, line := range testLines {
		if strings.HasPrefix(line, "TestCase#") {
			chunks := strings.Split(line, "::")
			passed, err := strconv.ParseBool(chunks[3])
			if err != nil {
				fmt.Errorf("error parsing test result %s", line)
			}
			testResult := TestResult{
				TestName: chunks[0],
				Actual:   chunks[1],
				Expected: chunks[2],
				Passed:   passed,
			}
			testResults = append(testResults, testResult)
		}
	}
	t.TestsResults = testResults
}
