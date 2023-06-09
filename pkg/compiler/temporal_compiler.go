package compiler

import (
	"context"
	dbmodels "github.com/dgozalo/aec-remote-executor/pkg/database/model"
	"github.com/dgozalo/aec-remote-executor/pkg/graph/model"
	"github.com/dgozalo/aec-remote-executor/pkg/worker"
	"github.com/pkg/errors"
	v1 "go.temporal.io/api/enums/v1"
	"go.temporal.io/sdk/client"
	"log"
)

// TemporalCompiler is the service that handles the temporal compiler
type TemporalCompiler struct {
	TemporalClient client.Client
}

// TemporalExecution is the struct that contains the temporal execution information
type TemporalExecution struct {
	WorkflowID string
	RunID      string
}

// NewTemporalCompiler creates a new temporal compiler
func NewTemporalCompiler(c client.Client) *TemporalCompiler {
	return &TemporalCompiler{
		TemporalClient: c,
	}
}

// RunCompileWorker runs starts a mew Temporal Workflow that will compile and execute the code
func (c TemporalCompiler) RunCompileWorker(execution model.NewExecution) TemporalExecution {
	options := client.StartWorkflowOptions{
		ID:        "executions-workflow",
		TaskQueue: worker.ExecutionTaskQueue,
	}

	// Start the Workflow
	we, err := c.TemporalClient.ExecuteWorkflow(context.Background(), options, worker.ExecutionWorkflowName, execution)
	if err != nil {
		log.Fatalln("unable to complete Workflow", err)
	}
	return TemporalExecution{
		WorkflowID: we.GetID(),
		RunID:      we.GetRunID(),
	}
}

// GetCompilationStatus retrieves the status of the compilation from the Temporal workflow
func (c TemporalCompiler) GetCompilationStatus(execution *dbmodels.Execution) (*model.ExecutionResult, error) {
	ctx := context.Background()
	dwf, err := c.TemporalClient.DescribeWorkflowExecution(ctx, execution.WorkflowID, execution.RunID)
	if err != nil {
		return nil, errors.Wrap(err, "there was a problem describing the workflow")
	}
	finished := dwf.WorkflowExecutionInfo.Status != v1.WORKFLOW_EXECUTION_STATUS_RUNNING
	// If the workflow is not finished, we return the ExecutionStatusRunning status
	if !finished {
		return &model.ExecutionResult{
			Status: model.ExecutionStatusRunning,
		}, err
	}

	var result worker.ExecutionResult
	wr := c.TemporalClient.GetWorkflow(ctx, execution.WorkflowID, execution.RunID)
	err = wr.Get(ctx, &result)
	if err != nil {
		return nil, errors.Wrap(err, "there was a problem obtaining the status of the current temporal execution")
	}
	var testResults []*model.TestResult
	for _, testResult := range result.TestsResults {
		testResults = append(testResults, &model.TestResult{
			TestName: testResult.TestName,
			Expected: testResult.Expected,
			Actual:   testResult.Actual,
			Passed:   testResult.Passed,
		})
	}

	if result.Stderr != "" {
		return &model.ExecutionResult{
			Stdout: result.Stdout,
			Stderr: result.Stderr,
			Status: model.ExecutionStatusError,
		}, nil
	}

	return &model.ExecutionResult{
		Stdout:      result.Stdout,
		Stderr:      result.Stderr,
		TestResults: testResults,
		Status:      model.ExecutionStatusCompleted,
	}, nil
}
