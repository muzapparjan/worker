package worker

import (
	"log"
	"worker/worker/api/cuda"
	"worker/worker/api/opencl"
)

var api map[string]IAPI

type IAPI interface {
}

func InitAPI() map[string]IAPI {
	api = make(map[string]IAPI)

	cl, err := opencl.NewOCL()
	if err != nil {
		log.Printf("Failed to initialize OpenCL API:%v", err)
	} else {
		api["OpenCL"] = cl
		log.Printf("Successfully initialized OpenCL API.")
	}

	cuda, err := cuda.NewCUDA()
	if err != nil {
		log.Printf("Failed to initialize CUDA API:%v", err)
	} else {
		api["CUDA"] = cuda
		log.Printf("Successfully initialized CUDA API.")
	}

	if len(api) <= 0 {
		log.Fatalf("Failed to initialize API!")
	}

	return api
}
