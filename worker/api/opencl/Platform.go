package opencl

import "github.com/jgillich/go-opencl/cl"

type Platform struct {
	Name       string
	Vendor     string
	Version    string
	Profile    string
	Extensions string
	Devices    []*Device
}

func newPlatform(p *cl.Platform) (*Platform, error) {
	platform := &Platform{
		Name:       p.Name(),
		Vendor:     p.Vendor(),
		Version:    p.Version(),
		Profile:    p.Profile(),
		Extensions: p.Extensions(),
	}
	clDevices, err := p.GetDevices(cl.DeviceTypeAll)
	platform.Devices = make([]*Device, 0, len(clDevices))
	if err != nil {
		return nil, err
	}
	for _, clDevice := range clDevices {
		device, err := newDevice(clDevice)
		if err != nil {
			return nil, err
		}
		platform.Devices = append(platform.Devices, device)
	}
	return platform, nil
}
