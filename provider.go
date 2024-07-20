package gcp

import (
	"github.com/iodasolutions/xbee-common/cmd"
	"github.com/iodasolutions/xbee-common/provider"
)

type Provider struct {
}

func (pv Provider) Up() (*provider.InitialStatus, *cmd.XbeeError) {
	return nil, nil
}

func (pv Provider) Delete() *cmd.XbeeError {
	return nil
}
func (pv Provider) InstanceInfos() (map[string]*provider.InstanceInfo, *cmd.XbeeError) {
	return nil, nil
}
func (pv Provider) Image() *cmd.XbeeError {
	return nil
}
