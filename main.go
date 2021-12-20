package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var addr = flag.String("p", "8084", "http server port, default 8084")
var dir = flag.String("d", "./", "root dir in file system, default ./")

func Handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	if path == "favicon.ico" {
		http.NotFound(w, r)
		return
	}
	if path == "" {
		path = "index.html"
	}
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Fprintf(w, "404")
		return
	}
	fmt.Fprintf(w, "%s\n", contents)
}

func StaticServer(w http.ResponseWriter, req *http.Request) {
	staticHandler := http.FileServer(http.Dir(*dir))
	staticHandler.ServeHTTP(w, req)
	return
}

func main() {
	flag.Parse()
	addrstr := fmt.Sprintf(":%s", *addr)

	http.HandleFunc("/test12", Handler)
	http.HandleFunc("/", StaticServer)
	s := &http.Server{
		Addr: addrstr,
	}

	log.Printf("serve at htpp://%s for %s", addrstr, *dir)

	log.Fatal(s.ListenAndServe())
}
