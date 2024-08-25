package gcp

import (
	compute "cloud.google.com/go/compute/apiv1"
	"cloud.google.com/go/compute/apiv1/computepb"
	"context"
	"fmt"
	"github.com/iodasolutions/xbee-common/cmd"
	"github.com/iodasolutions/xbee-common/provider"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"log"
	"sync"
)

var (
	cfg  option.ClientOption
	once sync.Once
)

func loadConfig() option.ClientOption {
	once.Do(func() {
		cfg = option.WithCredentialsFile("")
	})
	return cfg
}

type Zone struct {
	Id string
}

func (z *Zone) instanceInfos() map[string]*provider.InstanceInfo {
	result := map[string]*provider.InstanceInfo{}
	return result
}

func (z *Zone) fillVolumes(ctx context.Context) *cmd.XbeeError {
	c, err := compute.NewDisksRESTClient(ctx, cfg)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer c.Close()

	projectID := "your-project-id"
	zone := "us-central1-a"

	req := &computepb.ListDisksRequest{
		Project: projectID,
		Zone:    zone,
	}

	it := c.List(ctx, req)
	fmt.Printf("Disks in zone %s:\n", zone)
	for {
		disk, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to list disks: %v", err)
		}
		fmt.Printf("- %s (Size: %d GB, Type: %s)\n", disk.GetName(), disk.GetSizeGb(), disk.GetType())
	}
}
