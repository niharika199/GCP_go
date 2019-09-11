package server

import (
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
	"log"
)

func Listservers(proj string, zone string) {
	ctx := context.Background()
	c, err := google.DefaultClient(ctx, compute.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	computeService, err := compute.New(c)
	if err != nil {
		log.Fatal(err)
	}
	k := []string{}
	var j Svrlistoutput

	req := computeService.Instances.List(proj, zone)

	if err := req.Pages(ctx, func(page *compute.InstanceList) error {
		for _, instance := range page.Items {
			// TODO: Change code below to process each `instance` resource:
			fmt.Printf("%+v\n", instance.Name)
			k = append(k, instance.Name)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(k); i++ {
		res, _ := computeService.Instances.Get(proj, zone, k[i]).Do()

		j = Svrlistoutput{Name: res.Name,
			PublicIP:  res.NetworkInterfaces[0].AccessConfigs[0].NatIP,
			PrivateIP: res.NetworkInterfaces[0].NetworkIP,
			Status:    res.Status}
		fmt.Printf("%+v\n", j)
	}
}

/*func main() {
	Listservers("proj-211305", "us-east1-b")
}*/
