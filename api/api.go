package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	network "github.com/niharika199/GCP_go/compute/network"
	server "github.com/niharika199/GCP_go/compute/server"
	"log"
	"net/http"
	"os"
)

func main() {
	r := mux.NewRouter()
	//	r.HandleFunc("/", welcome)
	file, err := os.OpenFile("../logs/gcp.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	r.HandleFunc("/servercreate", create)
	r.HandleFunc("/serverdelete", Delete)
	r.HandleFunc("/networkcreate", netcreate)
	r.HandleFunc("/networkdelete", netdel)
	//      r.HandleFunc("/listservers",listservers)
	if err := http.ListenAndServe(":80", r); err != nil {
		log.Fatal(err)
	}
}

/*func welcome(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("layout.html")
	tmpl.Execute(w, nil)
}*/

func create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	input := server.Serverinput{}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		fmt.Println(w, "Invalid request payload")
		return
	}
	serverout := input.Createserver()

	log.Println("Created the instance", serverout.Name, serverout.PublicIP)
	fmt.Fprintln(w, "Created the server")
	fmt.Fprintln(w, serverout.Name)
	fmt.Fprintln(w, serverout.PublicIP)

}

func Delete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	input := server.Serverdelinput{}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		fmt.Println(w, "Invalid request payload")
		return
	}
	input.Deleteserver()
	// add the code to handle the error if not found
	log.Println("deleted the server", input.Name)
	fmt.Fprintln(w, "deleted the instance")
}

func netcreate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	input := network.Networkinput{}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		fmt.Println(w, "Invalid request payload")
		return
	}
	input.Createnetwork()
	fmt.Fprintln(w, "Created the Network")
}

func netdel(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	input := network.Networkdelinput{}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		fmt.Println(w, "Invalid request payload")
		return
	}
	input.Deletenetwork()
	fmt.Fprintln(w, "Deleted the Network")
}

/*nc listservers(w http.ResponseWriter, r *http.Request) {
        defer r.Body.Close()
        input := server.listserver{}
        if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
                fmt.Println(w,"Invalid request payload")
                return
        }
}*/
