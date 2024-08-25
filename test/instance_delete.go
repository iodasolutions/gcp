package main

import (
	compute "cloud.google.com/go/compute/apiv1"
	"cloud.google.com/go/compute/apiv1/computepb"
	"context"
	"github.com/iodasolutions/gcp/test/config"
	"github.com/iodasolutions/xbee-common/log2"
	"log"
)

func main() {
	name := "moninstance"
	ctx := context.Background()
	instancesClient, err := compute.NewInstancesRESTClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create compute service: %v", err)
	}
	defer instancesClient.Close()
	req := &computepb.DeleteInstanceRequest{
		Instance: name,
		Project:  config.ProjectId,
		Zone:     config.Zone,
	}
	op, err := instancesClient.Delete(ctx, req)
	if err != nil {
		log.Fatalf("unable to create instance: %w", err)
	}

	if err = op.Wait(ctx); err != nil {
		log.Fatalf("unable to wait for the operation: %w", err)
	}

	log2.Infof("Instance deleted")
}
