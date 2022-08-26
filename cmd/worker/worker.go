package main

import (
	"log"
	"worker/worker"
)

func main() {
	worker, err := worker.NewWorker()
	if err != nil {
		log.Fatalf("Failed to create worker:\n%v", err)
	}
	err = worker.Work()
	if err != nil {
		log.Fatalf("Error occured while working:\n%v", err)
	}
}
