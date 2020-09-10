package main

import (
	"flag"
	"fmt"
	"net/http"

	controller "short-links/controller"
)

func main() {
	yamlFileName := flag.String("yaml", "urls.yaml", "a yaml file in the format of 'path,url'")
	flag.Parse()

	mux := defaultMux()
	pathToUrls := make(map[string]string)
	mapController := controller.MapController(pathToUrls, mux)

	yamlFileContent := controller.GetYAMLFileContent(*yamlFileName)

	yamlController, err := controller.YAMLController([]byte(yamlFileContent), mapController)
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
