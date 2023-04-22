package compiler

import (
	"context"
	"fmt"
	"github.com/dgozalo/aec-remote-executor/pkg/graph/model"
	models "github.com/dgozalo/aec-remote-executor/pkg/model"
	"github.com/dgozalo/aec-remote-executor/pkg/worker"
	"github.com/pkg/errors"
	v1 "go.temporal.io/api/enums/v1"
	"go.temporal.io/sdk/client"
	"log"
)

type TemporalCompiler struct {
	TemporalClient client.Client
}

type TemporalExecution struct {
	WorkflowID string
	RunID      string
}

func NewTemporalCompiler(c client.Client) *TemporalCompiler {
	return &TemporalCompiler{
		TemporalClient: c,
	}
}

func (c TemporalCompiler) RunCompileWorker(execution model.NewExecution) TemporalExecution {
	options := client.StartWorkflowOptions{
		ID:        "executions-workflow",
		TaskQueue: worker.ExecutionTaskQueue,
	}

	// Start the Workflow
	we, err := c.TemporalClient.ExecuteWorkflow(context.Background(), options, worker.ExecutionWorkflow, execution)
	if err != nil {
		log.Fatalln("unable to complete Workflow", err)
	}
	return TemporalExecution{
		WorkflowID: we.GetID(),
		RunID:      we.GetRunID(),
	}
}

func (c TemporalCompiler) GetCompilationStatus(execution *models.InternalExecution) (*model.ExecutionResult, error) {
	ctx := context.Background()
	dwf, err := c.TemporalClient.DescribeWorkflowExecution(ctx, execution.WorkflowID, execution.RunID)
	if err != nil {
		return nil, errors.Wrap(err, "there was a problem describing the workflow")
	}
	finished := dwf.WorkflowExecutionInfo.Status != v1.WORKFLOW_EXECUTION_STATUS_RUNNING
	if !finished {
		return &model.ExecutionResult{
			CompilationResult: "",
			Finished:          false,
			Error:             false,
		}, err
	}

	var result string
	wr := c.TemporalClient.GetWorkflow(ctx, execution.WorkflowID, execution.RunID)
	err = wr.Get(ctx, &result)
	if err != nil {
		return nil, errors.Wrap(err, "there was a problem obtaining the status of the current temporal execution")
	}
	return &model.ExecutionResult{
		CompilationResult: result,
		Finished:          true,
		Error:             false,
	}, nil
}

func printResults(greeting string, workflowID, runID string) {
	fmt.Printf("\nWorkflowID: %s RunID: %s\n", workflowID, runID)
	fmt.Printf("\n%s\n\n", greeting)
}
