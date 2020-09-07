package main

import (
	"fmt"
	"net/http"
)



func main() {
	mux := defaultMux()

	pathToUrls := map[string]string{
		"/godoc": "https://godoc.org",
		"/go-terminal-ui": "https://github.com/semihsemih/go-terminal-ui",
	}

	mapHandler := mapHandler(pathToUrls, mux)

	fmt.Println("Server start on :8080")
	http.ListenAndServe(":9000", mapHandler)
}

func defaultMux() *http.ServeMux  {
	mux := http.NewServeMux()
	mux.HandleFunc("/", sayHello)
	return mux
}

func sayHello(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello, World")
}

func mapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc  {
	return func(writer http.ResponseWriter, request *http.Request) {
		path := request.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(writer, request, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(writer, request)
	}
}