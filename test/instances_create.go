package main

import (
	compute "cloud.google.com/go/compute/apiv1"
	"cloud.google.com/go/compute/apiv1/computepb"
	"context"
	"fmt"
	"github.com/iodasolutions/gcp/test/config"
	"github.com/iodasolutions/xbee-common/log2"
	"github.com/iodasolutions/xbee-common/provider"
	"google.golang.org/protobuf/proto"
	"log"
)

func main() {
	name := "moninstance"
	diskImage := fmt.Sprintf("projects/%s/global/images/family/%s", config.ImageProject, config.ImageFamily)
	machineType := fmt.Sprintf("zones/%s/machineTypes/%s", config.Zone, config.InstanceType)
	network := fmt.Sprintf("global/networks/%s", config.Network)
	ctx := context.Background()
	instancesClient, err := compute.NewInstancesRESTClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create compute service: %v", err)
	}
	defer instancesClient.Close()
	startupScript := provider.AuthorizedKeyScript(config.User)

	req := &computepb.InsertInstanceRequest{
		Project: config.ProjectId,
		Zone:    config.Zone,
		InstanceResource: &computepb.Instance{
			Name: proto.String(name),
			Disks: []*computepb.AttachedDisk{
				{
					InitializeParams: &computepb.AttachedDiskInitializeParams{
						DiskSizeGb:  proto.Int64(config.InstanceDiskSize),
						SourceImage: proto.String(diskImage),
					},
					AutoDelete: proto.Bool(true),
					Boot:       proto.Bool(true),
					Type:       proto.String(computepb.AttachedDisk_PERSISTENT.String()),
				},
			},
			MachineType: proto.String(machineType),
			NetworkInterfaces: []*computepb.NetworkInterface{
				{
					AccessConfigs: []*computepb.AccessConfig{
						{
							// Une configuration AccessConfig avec "type" défini sur "ONE_TO_ONE_NAT"
							// spécifie que cette interface réseau est configurée avec une IP externe.
							Name: config.String("External NAT"),
							Type: config.String("ONE_TO_ONE_NAT"),
						},
					},
					Name: proto.String(network),
				},
			},
			Metadata: &computepb.Metadata{
				Items: []*computepb.Items{
					{
						Key:   config.String("startup-script"),
						Value: &startupScript,
					},
				},
			},
		},
	}

	op, err := instancesClient.Insert(ctx, req)
	if err != nil {
		log.Fatalf("unable to create instance: %w", err)
	}

	if err = op.Wait(ctx); err != nil {
		log.Fatalf("unable to wait for the operation: %w", err)
	}

	log2.Infof("Instance created")

}
