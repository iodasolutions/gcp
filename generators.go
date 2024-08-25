package gcp

import (
	"context"
	"github.com/iodasolutions/xbee-common/cmd"
	"github.com/iodasolutions/xbee-common/log2"
	"github.com/iodasolutions/xbee-common/provider"
	"github.com/iodasolutions/xbee-common/util"
)

type response struct {
	z   *Zone
	err *cmd.XbeeError
}

func zonesForHosts(ctx context.Context) (map[string]*Zone, *cmd.XbeeError) {
	hosts, err := HostsByZone()
	if err != nil {
		return nil, err
	}
	//volumes, err := VolumesFrom()
	//if err != nil {
	//	return nil, err
	//}

	var channels []<-chan *response
	for regionName, hostsForRegion := range hosts {
		volumesForRegion := volumes[regionName]
		channels = append(channels, newRegion(ctx, regionName, hostsForRegion, volumesForRegion))
	}
	ch := util.Multiplex(ctx, channels...)
	result := map[string]*Region2{}
	for resp := range ch {
		if resp.err != nil {
			log2.Errorf("%v", resp.err)
		} else {
			result[resp.r.Name] = resp.r
		}
	}
	return result, nil
}

func HostsByZone() (map[string]map[string]*ProviderHost, *cmd.XbeeError) {
	hosts := provider.Hosts()
	result := map[string]map[string]*ProviderHost{}
	for _, hReq := range hosts {
		h, err := hostFrom(hReq)
		if err != nil {
			return nil, err
		}
		if _, ok := result[h.Specification.Zone]; !ok {
			result[h.Specification.Zone] = map[string]*ProviderHost{}
		}
		result[h.Specification.Zone][h.Name] = h
	}
	return result, nil
}
