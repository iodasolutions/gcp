package gcp

import (
	"context"
	"github.com/iodasolutions/xbee-common/cmd"
	"github.com/iodasolutions/xbee-common/log2"
	"github.com/iodasolutions/xbee-common/provider"
)

type Provider struct {
}

func (pv Provider) Up() ([]*provider.InstanceInfo, *cmd.XbeeError) {
	log2.Infof("starting up gcp provider")
	return nil, nil
}

func (pv Provider) Delete() *cmd.XbeeError {
	log2.Infof("starting delete gcp provider")
	return nil
}
func (pv Provider) InstanceInfos() ([]*provider.InstanceInfo, *cmd.XbeeError) {
	ctx := context.Background()
	var result []*provider.InstanceInfo
	if zones, err := zonesForHosts(ctx); err != nil {
		return nil, err
	} else {
		for _, z := range zones {
			for _, info := range z.instanceInfos() {
				result = append(result, info)
			}
		}
		return result, nil
	}
}
func (pv Provider) Image() *cmd.XbeeError {
	return nil
}
