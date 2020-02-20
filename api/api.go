package main

import (
	logger "github.com/niharika199/GCP_go/logs"
	"encoding/json"
	"fmt"

	network "github.com/niharika199/GCP_go/compute/network"
	server "github.com/niharika199/GCP_go/compute/server"

	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/servercreate", create)
	http.HandleFunc("/serverdelete", Delete)
	http.HandleFunc("/networkcreate", netcreate)
	http.HandleFunc("/networkdelete", netdel)
	//      r.HandleFunc("/listservers",listservers)
	if err := http.ListenAndServe(":80", logRequest(http.DefaultServeMux)); err != nil {
		log.Fatal(err)
	}
}

/*func welcome(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("layout.html")
	tmpl.Execute(w, nil)
}*/

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.RequestLogger.Println(r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	input := server.Serverinput{}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		fmt.Println(w, "Invalid request payload")
		return
	}
	serverout := input.Createserver()
	logger.GeneralLogger.Println("Created the instance", serverout.Name, serverout.PublicIP)
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
	logger.GeneralLogger.Println("deleted the server", input.Name)
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
