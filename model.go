package gcp

import (
	"encoding/json"
	"fmt"
	"github.com/iodasolutions/xbee-common/cmd"
	"github.com/iodasolutions/xbee-common/provider"
	"github.com/iodasolutions/xbee-common/types"
	"github.com/iodasolutions/xbee-common/util"
)

type Model struct {
	ProjectId    string `json:"projectId"`
	Zone         string `json:"zone,omitempty"`
	InstanceType string `json:"instanceType,omitempty"`
	Image        struct {
		Family  string `json:"family,omitempty"`
		Project string `json:"project,omitempty"`
	} `json:"image,omitempty"`
	//
	//AvailabilityZone string `json:"availabilityZone,omitempty"`
	//OsArch           string `json:"osarch,omitempty"`
	//VolumeType       string `json:"volumeType,omitempty"`
	//Size             int    `json:"size,omitempty"`
}

func fromMap(aMap map[string]interface{}) (*Model, error) {
	var result Model
	data, err := util.NewJsonIO(aMap).SaveAsBytes()
	if err != nil {
		return nil, fmt.Errorf("unexpected when encoding to json : %v", err)
	}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("unexpected when decoding json to AWS model : %v", err)
	}
	return &result, nil
}

type ProviderHost struct {
	Specification *Model
	Name          string
	Ports         []string
	User          string
	Volumes       []string
	ExternalIp    string
	PackId        *types.IdJson
	PackHash      string
	SystemId      *types.IdJson
	SystemHash    string
}

func hostFrom(req *provider.Host) (*ProviderHost, *cmd.XbeeError) {
	m, err := fromMap(req.Provider)
	if err != nil {
		return nil, cmd.Error("cannot unmarshal json provider data for host %s : %v", req.Name, err)
	}
	return &ProviderHost{
		Specification: m,
		Name:          req.Name,
		Ports:         req.Ports,
		User:          req.User,
		Volumes:       req.Volumes,
		ExternalIp:    req.ExternalIp,
		PackId:        req.PackId,
		PackHash:      req.PackHash,
		SystemId:      req.SystemId,
		SystemHash:    req.SystemHash,
	}, nil
}

type Volume struct {
	provider.GenericVolume
	Specification *Model
}

func volumeFrom(req *provider.Volume) (*Volume, *cmd.XbeeError) {
	m, err := fromMap(req.Provider)
	if err != nil {
		return nil, cmd.Error("cannot unmarshal json provider data for volume %s : %v", req.Name, err)
	}
	return &Volume{
		GenericVolume: provider.FromVolume(req),
		Specification: m,
	}, nil
}

func VolumesFrom() (map[string]map[string]*Volume, *cmd.XbeeError) {
	volumes := provider.VolumesForEnv()
	result := make(map[string]map[string]*Volume)
	for _, vReq := range volumes {
		v, err := volumeFrom(vReq)
		if err != nil {
			return nil, err
		}
		if _, ok := result[v.Specification.Region]; !ok {
			result[v.Specification.Region] = map[string]*Volume{}
		}
		result[v.Specification.Region][v.Name] = v
	}
	return result, nil
}
