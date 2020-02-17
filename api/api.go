package main

import (
	"encoding/json"
	"fmt"
	network "gcp/compute/network"
	server "gcp/compute/server"
	"github.com/gorilla/mux"
	//	"html/template"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	//	r.HandleFunc("/", welcome)
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
        output:=input.Createnetwork()
        fmt.Fprintln(w, "Created the Network",output.name)
}*/
