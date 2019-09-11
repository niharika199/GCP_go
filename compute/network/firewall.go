package network

import (
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	compute "google.golang.org/api/compute/v1"
	"log"
)

type Firewallinput struct {
	Project    string
	Name       string
	Network    string
	Priority   int64
	IPProtocol string
	Ports      []string
}

func (i *Firewallinput) Firewallcrt() {
	ctx := context.Background()

	c, err := google.DefaultClient(ctx, compute.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	computeService, err := compute.New(c)
	if err != nil {
		log.Fatal(err)
	}

	rb := &compute.Firewall{
		Name:     i.Name,
		Network:  "projects/" + i.Project + "/global/networks/" + i.Network,
		Priority: i.Priority,
		Allowed: []*compute.FirewallAllowed{
			&compute.FirewallAllowed{
				IPProtocol: i.IPProtocol,
				Ports:      i.Ports,
			},
		},
	}
	resp, err := computeService.Firewalls.Insert(i.Project, rb).Context(ctx).Do()
	if err != nil {
		log.Fatal(err)

	}
	fmt.Printf("%#v\n", resp)
}

/*func main() {
	input := Firewallinput{
		Project:    "proj-211305",
		Name:       "testt",
		Network:    "net1",
		Priority:   333,
		IPProtocol: "tcp",
		Ports:      []string{"22", "80", "8080"},
	}
	input.Firewallcrt()
}*/
