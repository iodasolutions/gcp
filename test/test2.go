package main

import (
	"context"
	"fmt"
	"google.golang.org/api/compute/v1"
)

func main() {
	ctx := context.Background()
	computeService, err := compute.NewService(ctx)
	startupScript := `#/bin/bash
	echo Hello, Eric
	`
	// Créer une requête de création de VM.
	instance := &compute.Instance{
		Name:        "xbee-test",
		Zone:        "europe-west9-a",
		MachineType: "f1-micro",
		Disks: []*compute.AttachedDisk{
			{
				AutoDelete: true,
				Boot:       true,
				Type:       "PERSISTENT",
				InitializeParams: &compute.AttachedDiskInitializeParams{
					DiskSizeGb: 10,
					// Utilisez l'image Ubuntu 20.04 LTS depuis le projet ubuntu-os-cloud
					SourceImage: "projects/ubuntu-os-cloud/global/images/family/ubuntu-2004-lts",
				},
			},
		},
		Metadata: &compute.Metadata{
			Items: []*compute.MetadataItems{
				{
					Key:   "startup-script",
					Value: &startupScript,
				},
			},
		},
		// Définir d'autres options de la VM.
	}

	// Remplacez `votre-projet-id` par votre projet GCP et `votre-zone` par la zone de votre choix
	_, err = computeService.Instances.Insert("votre-projet-id", "votre-zone", instance).Context(ctx).Do()
	if err != nil {
		fmt.Printf("Failed to create instance: %s\n", err)
		return
	}

}
