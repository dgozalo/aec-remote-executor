package main

import (
	"github.com/dgozalo/aec-remote-executor/pkg/worker"
	"log"
)

// main is the entrypoint of the worker
func main() {
	w := worker.Worker{}
	// Run the worker
	err := w.InitWorker()
	if err != nil {
		log.Fatal(err)
	}
}
