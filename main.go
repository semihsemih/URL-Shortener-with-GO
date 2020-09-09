package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := defaultMux()
	pathToUrls := make(map[string]string)
	mapHandler := mapController(pathToUrls, mux)

	yaml := `
- path: /go-terminal-ui2
  url: https://github.com/semihsemih/go-terminal-ui
- path: /blackjack
  url: https://github.com/semihsemih/Blackjack-Vue
`
	yamlHandler, err := yamlController([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Server start on :9000")
	http.ListenAndServe(":9000", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", mainPageHandler)
	return mux
}

func mainPageHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "You can use the URL shortener as '...... / path'.")
}
