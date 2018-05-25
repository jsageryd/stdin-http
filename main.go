package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	port := flag.String("p", "8080", "Port")
	flag.Parse()
	if *port == "" {
		*port = "8080"
	}

	log.Println("Accepting data from stdin...")

	addr := ":" + *port
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Serving %d bytes at %s...\n", len(b), addr)

	http.ListenAndServe(addr, http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL.Path)
			http.ServeContent(w, r, "data", time.Now(), bytes.NewReader(b))
		},
	))
}
