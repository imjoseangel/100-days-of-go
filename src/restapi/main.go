///bin/true; exec /usr/bin/env go run "" "$@"
package main

import (
	"log"
	"net/http"
)

func main() {

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
