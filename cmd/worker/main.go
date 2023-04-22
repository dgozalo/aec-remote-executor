package main

import (
	"github.com/dgozalo/aec-remote-executor/pkg/worker"
	"log"
)

func main() {
	w := worker.Worker{}
	err := w.InitWorker()
	if err != nil {
		log.Fatal(err)
	}
}
