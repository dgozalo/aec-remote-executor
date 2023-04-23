package worker

import (
	"context"
	"fmt"
	"github.com/dgozalo/aec-remote-executor/pkg/graph/model"
	"github.com/dgozalo/aec-remote-executor/pkg/worker/compilers"
	"github.com/pkg/errors"
	"go.temporal.io/sdk/temporal"
)

func ExecutionActivity(ctx context.Context, execution model.NewExecution) (string, error) {
	exec := fmt.Sprintf("Compile code %s for language %s!", execution.Code, execution.Language)
	if execution.Language == "java" {
		result, err := compilers.RunCompile(execution.Code)
		if err != nil {
			return "", temporal.NewNonRetryableApplicationError(errors.Wrap(err, "there was a problem running the Java compilation").Error(), "", nil)
		}
		return result, nil
	}
	return exec, nil
}
