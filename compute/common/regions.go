package region

import (
	//	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
	"log"
)

var Region []string

func ListRegions(project string) []string {
	ctx := context.Background()

	c, err := google.DefaultClient(ctx, compute.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	computeService, err := compute.New(c)
	if err != nil {
		log.Fatal(err)
	}

	req := computeService.Regions.List(project)
	if err := req.Pages(ctx, func(page *compute.RegionList) error {
		for _, region := range page.Items {
			// TODO: Change code below to process each `region` resource:
			//mt.Printf("%#v\n", region.Name)
			Region = append(Region, region.Name)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return Region
}

/*func main() {
	k :=ListRegions("proj")
        for _,y := range k{
              fmt.Println(y)
}
}*/
