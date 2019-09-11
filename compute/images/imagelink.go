package main

import (
	"log"
	//"fmt"
	//"encoding/json"
	//"time"
	"golang.org/x/net/context"
	compute "google.golang.org/api/compute/v1"
	htransport "google.golang.org/api/transport/http"
)

type Imageinput struct {
	ImageProj   string
	Imagefamily string
}

func (c *Imageinput) getimagelink() {
	ctx := context.Background()
	client, _, err := htransport.NewClient(ctx)
	computeService, err := compute.New(client)
	if err != nil {
		log.Printf("some error", err)
	}
	resp, err := computeService.Images.GetFromFamily(c.ImageProj, c.Imagefamily).Context(ctx).Do() //in order to get the public images we need to get from their projects hence here project is of rhel-cloud
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v\n", resp)
}
func main() {
	c := Imageinput{
		ImageProj:   "rhel-cloud",
		Imagefamily: "rhel-7",
	}
	c.getimagelink()
}
