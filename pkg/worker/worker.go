package worker

import (
	"github.com/dgozalo/aec-remote-executor/pkg/database"
	"github.com/dgozalo/aec-remote-executor/pkg/management"
	"github.com/pkg/errors"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	workflow2 "go.temporal.io/sdk/workflow"
	"log"
	"os"
)

const ExecutionWorkflowName = "execution-workflow"

type Worker struct {
}

func (r Worker) InitWorker() error {
	temporalURL := os.Getenv("TEMPORALITE_HOST_PORT")
	if temporalURL == "" {
		temporalURL = "localhost:7233"
	}
	c, err := client.Dial(client.Options{
		HostPort: temporalURL,
	})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	w := worker.New(c, ExecutionTaskQueue, worker.Options{})
	pg, err := database.NewPostgresDBAccess()
	if err != nil {
		log.Fatalln("unable to obtain database client")
	}
	management := management.NewManagementService(pg)
	activity := NewActivity(management)
	workflow := NewWorkflow(activity)

	w.RegisterWorkflowWithOptions(workflow.ExecutionWorkflow, workflow2.RegisterOptions{Name: ExecutionWorkflowName})
	w.RegisterActivity(activity.ExecutionActivity)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		return errors.Wrap(err, "impossible to start the Executions Worker")
	}
	return nil
}
