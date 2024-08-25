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

	op, err := computeService.Disks.Delete(config.ProjectId, config.Zone, "example-disk").Context(ctx).Do()
	if err != nil {
		log.Fatalf("Failed to create disk: %v", err)
	}

	fmt.Printf("Disk creation operation status: %v\n", op.Status)
}
