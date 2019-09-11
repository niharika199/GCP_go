package snapshot
import (
        "fmt"
        "log"

        "golang.org/x/net/context"
        "golang.org/x/oauth2/google"
        "google.golang.org/api/compute/v1"
)

func CreateSnapshot(project string,zone string,disk string,name string) {
        ctx := context.Background()

        c, err := google.DefaultClient(ctx, compute.CloudPlatformScope)
        if err != nil {
                log.Fatal(err)
        }

        computeService, err := compute.New(c)
        if err != nil {
                log.Fatal(err)
        }


        rb := &compute.Snapshot{
           Name : name,
                // TODO: Add desired fields of the request body.
        }

        resp, err := computeService.Disks.CreateSnapshot(project, zone, disk, rb).Context(ctx).Do()
        if err != nil {
                log.Fatal(err)
        }
        fmt.Printf("created the snapshot %v\n",name)
        fmt.Printf("%#v\n", resp)
}

func main(){
 CreateSnapshot("proj-211305","us-east1-b","go","go")
}
