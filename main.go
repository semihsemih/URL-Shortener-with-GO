package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := defaultMux()

	http.ListenAndServe(":9000", mux)
}

func defaultMux() *http.ServeMux  {
	mux := http.NewServeMux()
	mux.HandleFunc("/", sayHello)
	return mux
}

func sayHello(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello, World")
}