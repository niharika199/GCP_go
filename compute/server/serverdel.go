package server

import (
	"fmt"
	"golang.org/x/net/context"
	compute "google.golang.org/api/compute/v1"
	htransport "google.golang.org/api/transport/http"
	"log"
	"time"
)

func (c *Serverdelinput) waitforOP(Name string) (progress bool, err error) {
	ctx := context.Background()
	client, _, err := htransport.NewClient(ctx)
	computeService, err := compute.New(client)
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

func (c *Serverdelinput) Deleteserver() {
	ctx := context.Background()
	client, _, err := htransport.NewClient(ctx)
	computeService, err := compute.New(client)
	if err != nil {
		log.Printf("some error", err)
	}
	ins, err1 := computeService.Instances.Delete(c.Project, c.Zone, c.Name).Do()
	if err1 != nil {
		fmt.Println("ERROR:%v", err1)
	}
	c.waitforOP(ins.Name)
	fmt.Println("DELETED THE INSTANCE SUCCESSFULLY") 
}

/*func main() {
	c := serverinput{
		Project: "proj-211305",
		Name:    "priv",
		Zone:    "us-east1-b",
	}
	c.deleteserver()
}*/
