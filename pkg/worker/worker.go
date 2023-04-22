package worker

import (
	"github.com/pkg/errors"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"log"
	"os"
)

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
	w.RegisterWorkflow(ExecutionWorkflow)
	w.RegisterActivity(ExecutionActivity)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		return errors.Wrap(err, "impossible to start the Executions Worker")
	}
	return nil
}
