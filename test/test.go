package main

import (
	compute "cloud.google.com/go/compute/apiv1"
	"cloud.google.com/go/compute/apiv1/computepb"
	"context"
	"fmt"
	"github.com/iodasolutions/xbee-common/provider"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/proto"
	"io"
	"log"
	"os"
)

func main() {
	//	ctx := context.Background()
	//	computeService, err := compute.NewService(ctx)
	//	startupScript := `#/bin/bash
	//echo Hello, Eric
	//`
	//	// Créer une requête de création de VM.
	//	instance := &compute.Instance{
	//		Name:        "xbee-test",
	//		Zone:        "europe-west9-a",
	//		MachineType: "f1-micro",
	//		Disks: []*compute.AttachedDisk{
	//			{
	//				AutoDelete: true,
	//				Boot:       true,
	//				Type:       "PERSISTENT",
	//				InitializeParams: &compute.AttachedDiskInitializeParams{
	//					DiskSizeGb: 10,
	//					// Utilisez l'image Ubuntu 20.04 LTS depuis le projet ubuntu-os-cloud
	//					SourceImage: "projects/ubuntu-os-cloud/global/images/family/ubuntu-2004-lts",
	//				},
	//			},
	//		},
	//		Metadata: &compute.Metadata{
	//			Items: []*compute.MetadataItems{
	//				{
	//					Key:   "startup-script",
	//					Value: &startupScript,
	//				},
	//			},
	//		},
	//		// Définir d'autres options de la VM.
	//	}
	//
	//	// Remplacez `votre-projet-id` par votre projet GCP et `votre-zone` par la zone de votre choix
	//	resp, err := computeService.Instances.Insert("votre-projet-id", "votre-zone", instance).Context(ctx).Do()
	//	if err != nil {
	//		fmt.Printf("Failed to create instance: %s\n", err)
	//		return
	//	}

	err := createInstance(os.Stdout, "ageless-domain-414706", "europe-west9-b", "moninstance", "e2-medium", "projects/ubuntu-os-cloud/global/images/family/ubuntu-2004-lts\n", "global/networks/default")
	if err != nil {
		log.Fatal(err)
	}
}

func String(v string) *string {
	return &v
}

// createInstance sends an instance creation request to the Compute Engine API and waits for it to complete.
func createInstance(w io.Writer, projectID, zone, instanceName, machineType, sourceImage, networkName string) error {
	// projectID := "ageless-domain-414706"
	// zone := "europe-west9-b"
	// instanceName := "mon_instance"
	// machineType := "n1-standard-1"
	// sourceImage := "ubuntu-os-cloud/ubuntu-focal"
	// networkName := "global/networks/default"

	ctx := context.Background()
	instancesClient, err := compute.NewInstancesRESTClient(ctx)
	if err != nil {
		return fmt.Errorf("NewInstancesRESTClient: %w", err)
	}
	defer instancesClient.Close()
	startupScript := provider.AuthorizedKeyScript("ubuntu")
	req := &computepb.InsertInstanceRequest{
		Project: projectID,
		Zone:    zone,
		InstanceResource: &computepb.Instance{
			Name: proto.String(instanceName),
			Disks: []*computepb.AttachedDisk{
				{
					InitializeParams: &computepb.AttachedDiskInitializeParams{
						DiskSizeGb:  proto.Int64(10),
						SourceImage: proto.String(sourceImage),
					},
					AutoDelete: proto.Bool(true),
					Boot:       proto.Bool(true),
					Type:       proto.String(computepb.AttachedDisk_PERSISTENT.String()),
				},
			},
			MachineType: proto.String(fmt.Sprintf("zones/%s/machineTypes/%s", zone, machineType)),
			NetworkInterfaces: []*computepb.NetworkInterface{
				{
					AccessConfigs: []*computepb.AccessConfig{
						{
							// Une configuration AccessConfig avec "type" défini sur "ONE_TO_ONE_NAT"
							// spécifie que cette interface réseau est configurée avec une IP externe.
							Name: String("External NAT"),
							Type: String("ONE_TO_ONE_NAT"),
						},
					},
					Name: proto.String(networkName),
				},
			},
			Metadata: &computepb.Metadata{
				Items: []*computepb.Items{
					{
						Key:   String("startup-script"),
						Value: &startupScript,
					},
				},
			},
		},
	}

	op, err := instancesClient.Insert(ctx, req)
	if err != nil {
		return fmt.Errorf("unable to create instance: %w", err)
	}

	if err = op.Wait(ctx); err != nil {
		return fmt.Errorf("unable to wait for the operation: %w", err)
	}

	fmt.Fprintf(w, "Instance created\n")

	return nil
}

// printImagesList prints a list of all non-deprecated image names available in given project.
func printImagesList(w io.Writer, projectID string) error {
	// projectID := "your_project_id"
	ctx := context.Background()
	imagesClient, err := compute.NewImagesRESTClient(ctx)
	if err != nil {
		return fmt.Errorf("NewImagesRESTClient: %w", err)
	}
	defer imagesClient.Close()

	// Listing only non-deprecated images to reduce the size of the reply.
	req := &computepb.ListImagesRequest{
		Project:    projectID,
		MaxResults: proto.Uint32(3),
		Filter:     proto.String("deprecated.state != DEPRECATED"),
	}

	// Although the `MaxResults` parameter is specified in the request, the iterator returned
	// by the `list()` method hides the pagination mechanic. The library makes multiple
	// requests to the API for you, so you can simply iterate over all the images.
	it := imagesClient.List(ctx, req)
	for {
		image, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		fmt.Fprintf(w, "- %s\n", image.GetName())
	}
	return nil
}

//func printImagesList(w io.Writer, projectID string) error {
//	// projectID := "your_project_id"
//	ctx := context.Background()
//	imagesClient, err := compute.NewImagesRESTClient(ctx)
//	if err != nil {
//		return fmt.Errorf("NewImagesRESTClient: %w", err)
//	}
//	defer imagesClient.Close()
//
//	// Listing only non-deprecated images to reduce the size of the reply.
//	req := &computepb.ListImagesRequest{
//		Project:    projectID,
//		MaxResults: proto.Uint32(3),
//		Filter:     proto.String("deprecated.state != DEPRECATED"),
//	}
//
//	// Although the `MaxResults` parameter is specified in the request, the iterator returned
//	// by the `list()` method hides the pagination mechanic. The library makes multiple
//	// requests to the API for you, so you can simply iterate over all the images.
//	it := imagesClient.List(ctx, req)
//	for {
//		image, err := it.Next()
//		if err == iterator.Done {
//			break
//		}
//		if err != nil {
//			return err
//		}
//		fmt.Fprintf(w, "- %s\n", image.GetName())
//	}
//	return nil
//}
