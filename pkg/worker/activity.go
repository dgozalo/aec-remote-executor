package worker

import (
	"context"
	"fmt"
	"github.com/dgozalo/aec-remote-executor/pkg/graph/model"
)

func ExecutionActivity(ctx context.Context, execution model.NewExecution) (string, error) {
	exec := fmt.Sprintf("Compile code %s for language %s!", execution.Code, execution.Language)
	return exec, nil
}
