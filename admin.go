package gcp

import "github.com/iodasolutions/xbee-common/cmd"

type Admin struct {
}

func (pv Admin) DestroyVolumes(names []string) *cmd.XbeeError {
	return nil
}
