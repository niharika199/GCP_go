package main

import (
	"encoding/json"
	"fmt"
	server "gcp/server"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", welcome)
	r.HandleFunc("/servercreate", create)
	if err := http.ListenAndServe(":80", r); err != nil {
		log.Fatal(err)
	}
}

func welcome(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("layout.html")
	tmpl.Execute(w, nil)
}

func create(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	input := server.Serverinput{
		Project:     r.FormValue("Project"),
		Name:        r.FormValue("Name"),
		MachineType: r.FormValue("MachineType"),
		ImageProj:   r.FormValue("ImageProj"),
		Imagefamily: r.FormValue("Imagefamily"),
		Zone:        r.FormValue("Zone"),
		Network:     r.FormValue("Network"),
		Subnet:      r.FormValue("Subnet"),
		Region:      r.FormValue("Region"),
		User:        r.FormValue("User"),
		Pem:         r.FormValue("Pem"),
	}
        fmt.Fprintln(w, "Creating the instance", r.FormValue("Name"))
	serverout := input.Createserver()
	fmt.Fprintln(w, "Created the server")
	fmt.Fprintln(w, serverout.Name)
	fmt.Fprintln(w, serverout.PublicIP)

}

