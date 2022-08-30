package opencl

import (
	"github.com/jgillich/go-opencl/cl"
)

var ocl *OpenCL

type OpenCL struct {
	platforms []*cl.Platform
}

func NewOCL() (*OpenCL, error) {
	var err error
	if ocl == nil {
		ocl = &OpenCL{}
		ocl.platforms, err = cl.GetPlatforms()
		if err != nil {
			return nil, err
		}
	}
	return ocl, nil
}
