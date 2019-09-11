package network

import (
	"fmt"
	"log"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
)

func Listnws(project string) {
	ctx := context.Background()

	c, err := google.DefaultClient(ctx, compute.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}
	k := []string{}
	computeService, err := compute.New(c)
	if err != nil {
		log.Fatal(err)
	}

	req := computeService.Networks.List(project)
	if err := req.Pages(ctx, func(page *compute.NetworkList) error {
		for _, network := range page.Items {
			// TODO: Change code below to process each `network` resource:
			fmt.Printf("%#v\n", network)
			k = append(k, network.Name)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", k)
}

/*func main() {
	Listnws("proj")
}*/
