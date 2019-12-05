package main

import (
	"log"
	"net/http"
)

func main() {
	r := Route()
	log.Fatal(http.ListenAndServe(":8080", r))
}
