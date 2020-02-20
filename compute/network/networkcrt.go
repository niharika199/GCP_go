package network

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	compute "google.golang.org/api/compute/v1"
	htransport "google.golang.org/api/transport/http"
	"log"
	"time"
)

type Networkinput struct {
	Project   string
	Name      string
	IPv4Range string
}

type SubnetInput struct {
	Project     string
	Name        string
	Region      string
	Network     string
	IpCidrRange string
}

func waitforNetwork(Name string, Project string) (progress bool, err error) {
	ctx := context.Background()
	client, _, err := htransport.NewClient(ctx)
	computeService, err := compute.New(client)
	for {
		resp, err := computeService.GlobalOperations.Get(Project, Name).Context(ctx).Do()
		if err != nil {
			return false, err
		}
		//		fmt.Println(resp.Progress)
		if resp.Progress == 100 {
			return true, nil
		}
		time.Sleep(2 * time.Second)
	}
}

func (c *Networkinput) Createnetwork() {
	ctx := context.Background()
	client, _, err := htransport.NewClient(ctx)
	computeService, err := compute.New(client)
	if err != nil {
		log.Printf("some error", err)
	}

	network := &compute.Network{
		Name:                  c.Name,
		AutoCreateSubnetworks: false,
		IPv4Range:             c.IPv4Range,
		RoutingConfig: &compute.NetworkRoutingConfig{
			RoutingMode: "Regional",
		},
		ForceSendFields: []string{"AutoCreateSubnetworks"}, //to avoid into legacy mode
	}
	ntwk, err := computeService.Networks.Insert(c.Project, network).Context(ctx).Do()
	if err != nil {
		log.Fatal(err)
	}
	waitforNetwork(ntwk.Name, c.Project)
	fmt.Printf("CREATED THE NETWORK SUCCESSFULLY")
	res, err := computeService.Networks.Get(c.Project, c.Name).Context(ctx).Do()
	if err != nil {
		log.Fatal(err)
	}
	val, _ := json.Marshal(res)
	fmt.Printf(string(val))
}

func (c *SubnetInput) Createsubnet() {

	ctx := context.Background()
	client, _, err := htransport.NewClient(ctx)
	computeService, err := compute.New(client)

	subnetIn := &compute.Subnetwork{
		Region:                c.Region,
		Name:                  c.Name,
		Network:               "projects/" + c.Project + "/global/networks/" + c.Network,
		IpCidrRange:           c.IpCidrRange,
		PrivateIpGoogleAccess: true,
	}
	sbnt, err := computeService.Subnetworks.Insert(c.Project, c.Region, subnetIn).Context(ctx).Do()
	if err != nil {
		log.Fatal(err)
	}
	waitforNetwork(sbnt.Name, c.Project)
	fmt.Printf("CREATED THE SUBNET SUCCESSFULLY")
	resp, err := computeService.Subnetworks.Get(c.Project, c.Region, c.Name).Context(ctx).Do()
	if err != nil {
		log.Fatal(err)
	}
	val, _ := json.Marshal(resp)
	fmt.Println(string(val))

	//      fmt.Printf("%#v\n", resp)
}

/*func main() {
	 c:= SubnetInput{
		Project:     "proj",
		Name:        "private",
		Region:      "us-east1",
		Network:     "net1",
		IpCidrRange: "10.112.0.0/16",
	}

	n := Networkinput{
		Project: "proj-211305",
		Name:    "ne1",
	}
	n.Createnetwork()
	c.Createsubnet()
}*/
