package  region

import (
//	"fmt"
	"log"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
)

var Machinetype []string

func ListMachinetypes(project string, zone string) []string {
	ctx := context.Background()

	c, err := google.DefaultClient(ctx, compute.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	computeService, err := compute.New(c)
	if err != nil {
		log.Fatal(err)
	}

	req := computeService.MachineTypes.List(project, zone)
	if err := req.Pages(ctx, func(page *compute.MachineTypeList) error {
		for _, machineType := range page.Items {
			// TODO: Change code below to process each `machineType` resource:
			fmt.Printf("%#v : %#v\n\n ", machineType.Name, machineType.Description)
                        Machinetype =append(MachineType,machineType.Name)
}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
      return Machinetype
}

/*func main() {
	listMachinetypes("proj", "us-east1-b")
}*/
