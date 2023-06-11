package worker

import (
	"github.com/dgozalo/aec-remote-executor/pkg/graph/model"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
	"time"
)

// Workflow is the struct that contains the Temporal workflow and the activity
type Workflow struct {
	activity *Activity
}

// NewWorkflow creates a new Workflow struct
func NewWorkflow(activity *Activity) *Workflow {
	return &Workflow{activity: activity}
}

// ExecutionWorkflow is the Workflow function that will be executed by Temporal when a new execution is received
func (w Workflow) ExecutionWorkflow(ctx workflow.Context, execution model.NewExecution) (*ExecutionResult, error) {
	options := workflow.ActivityOptions{
		ScheduleToCloseTimeout: time.Minute * 3,
		StartToCloseTimeout:    time.Minute,
		RetryPolicy: &temporal.RetryPolicy{
			BackoffCoefficient:     1.0,
			MaximumAttempts:        3,
			NonRetryableErrorTypes: []string{"COMPILE_FAILURE", "ASSIGNMENT_NOT_FOUND"},
		},
	}

	wfLogger := workflow.GetLogger(ctx)
	wfLogger.Info("Starting workflow for Language %s and Code %s", execution.Language, execution.Code)

	ctx = workflow.WithActivityOptions(ctx, options)

	result := ExecutionResult{}
	err := workflow.ExecuteActivity(ctx, w.activity.ExecutionActivity, execution).Get(ctx, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
