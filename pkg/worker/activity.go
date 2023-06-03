package worker

import (
	"context"
	"fmt"
	"github.com/dgozalo/aec-remote-executor/pkg/graph/model"
	"github.com/dgozalo/aec-remote-executor/pkg/worker/compilers/java"
	"github.com/pkg/errors"
	"go.temporal.io/sdk/temporal"
	"strconv"
	"strings"
)

func ExecutionActivity(ctx context.Context, execution model.NewExecution) (*ExecutionResult, error) {
	fmt.Printf("Compile code %s for language %s!", execution.Code, execution.Language)
	result := &ExecutionResult{}
	if execution.Language == "java" {
		stdout, stderr, results, err := java.RunCompile(execution.Code)
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
				Expected: chunks[1],
				Actual:   chunks[2],
				Passed:   passed,
			}
			testResults = append(testResults, testResult)
		}
	}
	t.TestsResults = testResults
}
