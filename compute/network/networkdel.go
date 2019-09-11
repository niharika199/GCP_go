package network

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	compute "google.golang.org/api/compute/v1"
	htransport "google.golang.org/api/transport/http"
	"log"
)

type Networkdelinput struct {
	Project string
	Name    string
}

func (c *Networkdelinput) Deletenetwork() {
	ctx := context.Background()
	client, _, err := htransport.NewClient(ctx)
	computeService, err := compute.New(client)
	if err != nil {
		log.Printf("some error", err)
	}

	resp, err := computeService.Networks.Delete(c.Project, c.Name).Context(ctx).Do()
	if err != nil {
		log.Fatal(err)
	}
	val, _ := json.Marshal(resp)
	fmt.Println(string(val))
}

/*func main() {
	c := Networkdelinput{
		Project: "proj-211305",
		Name:    "ne1",
	}
	c.Deletenetwork()
}*/
