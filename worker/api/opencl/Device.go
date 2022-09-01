package opencl

import (
	"github.com/jgillich/go-opencl/cl"
)

type Device struct {
	Available              bool
	CompilerAvailable      bool
	LinkerAvailable        bool
	ImageSupport           bool
	ErrorCorrectionSupport bool
	EndianLittle           bool
	HostUnifiedMemory      bool

	Image2DMaxWidth    int
	Image2DMaxHeight   int
	Image3DMaxWidth    int
	Image3DMaxHeight   int
	Image3DMaxDepth    int
	ImageMaxArraySize  int
	ImageMaxBufferSize int

	NativeVectorWidthChar      int
	NativeVectorWidthShort     int
	NativeVectorWidthInt       int
	NativeVectorWidthLong      int
	NativeVectorWidthHalf      int
	NativeVectorWidthFloat     int
	NativeVectorWidthDouble    int
	PreferredVectorWidthChar   int
	PreferredVectorWidthShort  int
	PreferredVectorWidthInt    int
	PreferredVectorWidthLong   int
	PreferredVectorWidthHalf   int
	PreferredVectorWidthFloat  int
	PreferredVectorWidthDouble int

	Type                     uint
	HalfFPConfig             int
	DoubleFPConfig           int
	AddressBits              int
	MemBaseAddrAlign         int
	ExecutionCapabilities    int
	ProfilingTimerResolution int

	LocalMemType           int
	GlobalMemCacheType     int
	GlobalMemCacheSize     int
	GlobalMemCachelineSize int

	LocalMemSize          int64
	GlobalMemSize         int64
	MaxConstantBufferSize int64
	MaxMemAllocSize       int64

	MaxClockFrequency     int
	MaxComputeUnits       int
	MaxConstantArgs       int
	MaxParameterSize      int
	MaxReadImageArgs      int
	MaxWriteImageArgs     int
	MaxSamplers           int
	MaxWorkGroupSize      int
	MaxWorkItemDimensions int
	MaxWorkItemSizes      []int

	Name           string
	Vendor         string
	Profile        string
	Version        string
	DriverVersion  string
	OpenCLCVersion string
	Extensions     string
	BuiltInKernels string
}

func newDevice(d *cl.Device) (*Device, error) {
	device := &Device{
		Available:              d.Available(),
		CompilerAvailable:      d.CompilerAvailable(),
		LinkerAvailable:        d.LinkerAvailable(),
		ImageSupport:           d.ImageSupport(),
		ErrorCorrectionSupport: d.ErrorCorrectionSupport(),
		EndianLittle:           d.EndianLittle(),
		HostUnifiedMemory:      d.HostUnifiedMemory(),

		Image2DMaxWidth:    d.Image2DMaxWidth(),
		Image2DMaxHeight:   d.Image2DMaxHeight(),
		Image3DMaxWidth:    d.Image3DMaxWidth(),
		Image3DMaxHeight:   d.Image3DMaxHeight(),
		Image3DMaxDepth:    d.Image3DMaxDepth(),
		ImageMaxArraySize:  d.ImageMaxArraySize(),
		ImageMaxBufferSize: d.ImageMaxBufferSize(),

		NativeVectorWidthChar:      d.NativeVectorWidthChar(),
		NativeVectorWidthShort:     d.NativeVectorWidthShort(),
		NativeVectorWidthInt:       d.NativeVectorWidthInt(),
		NativeVectorWidthLong:      d.NativeVectorWidthLong(),
		NativeVectorWidthHalf:      d.NativeVectorWidthHalf(),
		NativeVectorWidthFloat:     d.NativeVectorWidthFloat(),
		NativeVectorWidthDouble:    d.NativeVectorWidthDouble(),
		PreferredVectorWidthChar:   d.PreferredVectorWidthChar(),
		PreferredVectorWidthShort:  d.PreferredVectorWidthShort(),
		PreferredVectorWidthInt:    d.PreferredVectorWidthInt(),
		PreferredVectorWidthLong:   d.PreferredVectorWidthLong(),
		PreferredVectorWidthHalf:   d.PreferredVectorWidthHalf(),
		PreferredVectorWidthFloat:  d.PreferredVectorWidthFloat(),
		PreferredVectorWidthDouble: d.PreferredVectorWidthDouble(),

		Type:                     uint(d.Type()),
		HalfFPConfig:             int(d.HalfFPConfig()),
		DoubleFPConfig:           int(d.DoubleFPConfig()),
		AddressBits:              d.AddressBits(),
		MemBaseAddrAlign:         d.MemBaseAddrAlign(),
		ExecutionCapabilities:    int(d.ExecutionCapabilities()),
		ProfilingTimerResolution: d.ProfilingTimerResolution(),

		LocalMemType:           int(d.LocalMemType()),
		GlobalMemCacheType:     int(d.GlobalMemCacheType()),
		GlobalMemCacheSize:     d.GlobalMemCacheSize(),
		GlobalMemCachelineSize: d.GlobalMemCachelineSize(),

		LocalMemSize:          d.LocalMemSize(),
		GlobalMemSize:         d.GlobalMemSize(),
		MaxConstantBufferSize: d.MaxConstantBufferSize(),
		MaxMemAllocSize:       d.MaxMemAllocSize(),

		MaxClockFrequency:     d.MaxClockFrequency(),
		MaxComputeUnits:       d.MaxComputeUnits(),
		MaxConstantArgs:       d.MaxConstantArgs(),
		MaxParameterSize:      d.MaxParameterSize(),
		MaxReadImageArgs:      d.MaxReadImageArgs(),
		MaxWriteImageArgs:     d.MaxWriteImageArgs(),
		MaxSamplers:           d.MaxSamplers(),
		MaxWorkGroupSize:      d.MaxWorkGroupSize(),
		MaxWorkItemDimensions: d.MaxWorkItemDimensions(),
		MaxWorkItemSizes:      d.MaxWorkItemSizes(),

		Name:           d.Name(),
		Vendor:         d.Vendor(),
		Profile:        d.Profile(),
		Version:        d.Version(),
		DriverVersion:  d.DriverVersion(),
		OpenCLCVersion: d.OpenCLCVersion(),
		//Extensions:     d.Extensions(), // FixMe There is an unknown error
		BuiltInKernels: d.BuiltInKernels(),
	}
	return device, nil
}
