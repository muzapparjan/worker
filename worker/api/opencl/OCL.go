package opencl

import (
	"github.com/jgillich/go-opencl/cl"
)

var ocl *OpenCL

type OpenCL struct {
	platforms []*Platform
}

func NewOCL() (*OpenCL, error) {
	var err error
	if ocl == nil {
		ocl, err = newOCL()
		if err != nil {
			return nil, err
		}
	}
	return ocl, nil
}

func newOCL() (*OpenCL, error) {
	ocl = &OpenCL{}
	clPlatforms, err := cl.GetPlatforms()
	if err != nil {
		return nil, err
	}
	ocl.platforms = make([]*Platform, 0, len(clPlatforms))
	for _, clPlatform := range clPlatforms {
		platform, err := newPlatform(clPlatform)
		if err != nil {
			return nil, err
		}
		ocl.platforms = append(ocl.platforms, platform)
	}
	return ocl, nil
}
