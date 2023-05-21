package worker

import (
	"github.com/dgozalo/aec-remote-executor/pkg/graph/model"
	"go.temporal.io/sdk/workflow"
	"time"
)

func ExecutionWorkflow(ctx workflow.Context, execution model.NewExecution) error {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}

	ctx = workflow.WithActivityOptions(ctx, options)
	err := workflow.ExecuteActivity(ctx, ExecutionActivity, execution).Get(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}
