package main

import (
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
	"log"
)

func Listimages(Imageproject string) {
	ctx := context.Background()

	c, err := google.DefaultClient(ctx, compute.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	computeService, err := compute.New(c)
	if err != nil {
		log.Fatal(err)
	}

	req := computeService.Images.List(Imageproject)
	if err := req.Pages(ctx, func(page *compute.ImageList) error {
		for _, image := range page.Items {
			// TODO: Change code below to process each `image` resource:
			fmt.Printf("%#v\n", image.Name)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
}

func main() {
	Listimages("rhel-cloud")
}
