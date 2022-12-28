package main

import (
	"io"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

func main() {
	// Hello world, the web server
	var yamlObj struct{}

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")

		
		dec := yaml.NewDecoder(req.Body)
		dec.SetStrict(true)
		if err := dec.Decode(yamlObj); err != nil {
			io.WriteString(w, "error happened")
		}
	}

	http.HandleFunc("/hello", helloHandler)
    log.Println("Listing for requests at http://localhost:8000/hello")
	log.Fatal(http.ListenAndServe(":8000", nil))
}