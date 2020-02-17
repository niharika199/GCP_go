package snapshot

import (
	"fmt"
	"log"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
)

func Snapshotdel(project string, snapshot string) {
	ctx := context.Background()

	c, err := google.DefaultClient(ctx, compute.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	computeService, err := compute.New(c)
	if err != nil {
		log.Fatal(err)
	}

	_, er := computeService.Snapshots.Delete(project, snapshot).Context(ctx).Do()
	if er != nil {
		log.Fatal(er)
	}

	// TODO: Change code below to process the `resp` object:
	fmt.Printf("DELETED the snapshot %v\n", snapshot)
}

func main() {
	Snapshotdel("proj-211305", "go")
}
