package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	compute "google.golang.org/api/compute/v1"
	htransport "google.golang.org/api/transport/http"
	"log"
	"time"
)

type serverinput struct {
	Project     string
	Name        string
	MachineType string
	Image       string
	Zone        string
	Network     string
}

var k string

// default images url mentioned in the function
func Imagelist(image string) string {
	switch image {
	case "centos":
		k = "https://www.googleapis.com/compute/v1/projects/centos-cloud/global/images/centos-6-v20150710"
	case "coreos":
		k = "https://www.googleapis.com/compute/v1/projects/coreos-cloud/global/images"
	case "debian":
		k = "https://www.googleapis.com/compute/v1/projects/debian-cloud/global/images/debian-9-stretch-v20181011"
	case "redhat":
		k = "https://www.googleapis.com/compute/v1/projects/rhel-cloud/global/images/rhel-7-v20181011"
		//   case "opensuse":
		//     k="https://www.googleapis.com/compute/v1/projects/opensuse-cloud/global/images"
	case "suse":
		k = "https://www.googleapis.com/compute/v1/projects/suse-cloud/global/images/sles-15-v20180816"
	case "ubuntu":
		k = "https://www.googleapis.com/compute/v1/projects/ubuntu-os-cloud/global/images/ubuntu-1504-vivid-v20150616a"
	case "windows":
		k = "https://www.googleapis.com/compute/v1/projects/windows-cloud/global/images/windows-server-2016-dc-v20181009"
	}
	return k

}

func (c *serverinput) createserver() {
	ctx := context.Background()
	client, _, err := htransport.NewClient(ctx)
	computeService, err := compute.New(client)
	if err != nil {
		log.Printf("some error", err)
	}

	/*resp, err := computeService.Images.GetFromFamily(c.ImageProj,c.Imagefamily).Context(ctx).Do()//in order to get the public images we need to get from their proj/?
	  ects hence here project is of rhel-cloud
	          if err != nil {
	                  log.Fatal(err)
	                }
	  log.Printf("%v",resp.SelfLink)*/
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
					SourceImage: Image(c.Image),
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
				Network: "/global/networks/default",
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
	op, err := computeService.Instances.Insert(c.Project, c.Zone, instance).Do()
	time.Sleep(1000)
	log.Printf("created instance successfully: %#v,err: %v", op, err)
	log.Println(op.Name)
	res, err := computeService.Instances.Get(c.Project, c.Zone, c.Name).Do()
	time.Sleep(10000)
	val, _ := json.Marshal(res)
	fmt.Println(string(val))
}
func main() {
	c := serverinput{
		Project:     "proj-211305",
		Name:        "testtt",
		MachineType: "n1-standard-1",
		Image:       "redhat",
		Zone:        "us-west1-b",
	}
	c.createserver()
}
