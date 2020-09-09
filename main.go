package main

import (
	"fmt"
	"net/http"

	controller "short-links/controller"
)

func main() {
	mux := defaultMux()
	pathToUrls := make(map[string]string)
	mapController := controller.MapController(pathToUrls, mux)

	yaml := `
- path: /go-terminal-ui2
  url: https://github.com/semihsemih/go-terminal-ui
- path: /blackjack
  url: https://github.com/semihsemih/Blackjack-Vue
`
	yamlController, err := controller.YAMLController([]byte(yaml), mapController)
	if err != nil {
		panic(err)
	}

	fmt.Println("Server start on :9000")
	http.ListenAndServe(":9000", yamlController)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", mainPageHandler)
	return mux
}

func mainPageHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "You can use the URL shortener as '....../path'.")
}
