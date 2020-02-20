package server

import (
	"fmt"
	"golang.org/x/net/context"
	compute "google.golang.org/api/compute/v1"
	//	htransport "google.golang.org/api/transport/http"
	"golang.org/x/oauth2/google"
	"log"
	"time"
)

func (c *Serverinput) Createserver() Serveroutput {
	ctx := context.Background()
	//	client, _, err := htransport.NewClient(ctx)
	//	computeService, err := compute.New(client)

	cl, err := google.DefaultClient(ctx, compute.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	computeService, err := compute.New(cl)
	if err != nil {
		log.Fatal(err)
	}
	// Getting the image from their respective project
	resp, err := computeService.Images.GetFromFamily(c.ImageProj, c.Imagefamily).Context(ctx).Do()
	if err != nil {
		log.Fatal(err)
	}

	var Subnet string
	var s *string
	value := c.User + ":" + c.Pem
	s = &value
	if c.Network != "default" {
		Subnet = "regions/" + c.Region + "/subnetworks/" + c.Subnet
	} else {
		Subnet = ""
	}
	instance := &compute.Instance{
		Name:        c.Name,
		Description: "compute sample instance",
		MachineType: "/zones/" + c.Zone + "/machineTypes/" + c.MachineType,
		Disks: []*compute.AttachedDisk{
			{
				AutoDelete: true,
				Boot:       true,
				Type:       "PERSISTENT",
				InitializeParams: &compute.AttachedDiskInitializeParams{
					DiskName:    c.Name,
					SourceImage: resp.SelfLink,
				},
			},
		},
		NetworkInterfaces: []*compute.NetworkInterface{
			&compute.NetworkInterface{
				AccessConfigs: []*compute.AccessConfig{
					&compute.AccessConfig{
						Type: "ONE_TO_ONE_NAT",
						Name: "External NAT",
					},
				},
				Network:    "/global/networks/" + c.Network,
				Subnetwork: Subnet,
			},
		},
		Metadata: &compute.Metadata{
			Items: []*compute.MetadataItems{
				{
					Key:   "ssh-keys",
					Value: s,
				},
			},
		},
		ServiceAccounts: []*compute.ServiceAccount{
			{
				Email: "default",
				Scopes: []string{
					compute.DevstorageFullControlScope,
					compute.ComputeScope,
				},
			},
		},
	}
	ins, err := computeService.Instances.Insert(c.Project, c.Zone, instance).Do()
	//waiting for the instance operation to complete
	c.waitforOP(ins.Name)
	fmt.Println("CREATED INSTANCE SUCCESSFULLY")
	// fetching the instance response through get
	res, err := computeService.Instances.Get(c.Project, c.Zone, c.Name).Do()
	//	val, _ := json.Marshal(res)
	//	result :=serveroutput{Name = string(val.name)}
	//	fmt.Println(string(val))
	var j Serveroutput
	j = Serveroutput{Name: res.Name,
		PublicIP:  res.NetworkInterfaces[0].AccessConfigs[0].NatIP,
		PrivateIP: res.NetworkInterfaces[0].NetworkIP}
	return j

}

func (c *Serverinput) waitforOP(Name string) (progress bool, err error) {
	ctx := context.Background()
	//	client, _, err := htransport.NewClient(ctx)
	cl, err := google.DefaultClient(ctx, compute.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}
	computeService, err := compute.New(cl)

	//looping until the instance operatin progress is 100
	for {
		resp1, err := computeService.ZoneOperations.Get(c.Project, c.Zone, Name).Context(ctx).Do()
		if err != nil {
			fmt.Println("error in op:")
			return false, err
		}
		//	fmt.Println(resp1.Progress)
		if resp1.Progress == 100 {
			return true, nil
		}
		time.Sleep(5 * time.Second)
	}
}

/*func main() {
	c := Serverinput{
		Project:     "proj",
		Name:        "name",
		MachineType: "n1-standard-1",
		ImageProj:   "rhel-cloud",
		Imagefamily: "rhel-7",
		Zone:        "us-east1-b",
		Network:     "default",
		Subnet:      "private",
		Region:      "us-east1",
		User:        "username",
		Pem:         "pemkey",
	}
	resu := c.createserver()
	fmt.Printf("%+v\n", resu)

}*/
