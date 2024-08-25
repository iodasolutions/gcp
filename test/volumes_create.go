package main

import (
	"context"
	"fmt"
	"github.com/iodasolutions/gcp/test/config"
	"log"

	"google.golang.org/api/compute/v1"
)

func main() {
	ctx := context.Background()

	computeService, err := compute.NewService(ctx)
	if err != nil {
		log.Fatalf("Failed to create compute service: %v", err)
	}

	disk := &compute.Disk{
		Name:   "example-disk",
		SizeGb: config.VolumeSize,
		Type:   fmt.Sprintf("zones/%s/diskTypes/%s", config.Zone, config.VolumeType),
	}

	op, err := computeService.Disks.Insert(config.ProjectId, config.Zone, disk).Context(ctx).Do()
	if err != nil {
		log.Fatalf("Failed to create disk: %v", err)
	}

	fmt.Printf("Disk creation operation status: %v\n", op.Status)
}
