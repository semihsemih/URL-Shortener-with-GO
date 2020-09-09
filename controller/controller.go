package controller

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

type pathURL struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func MapController(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		path := request.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(writer, request, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(writer, request)
	}
}

func YAMLController(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var pathUrls []pathURL
	err := yaml.Unmarshal(yamlBytes, &pathUrls)
	if err != nil {
		return nil, err
	}

	pathsToUrls := make(map[string]string)
	for _, pu := range pathUrls {
		pathsToUrls[pu.Path] = pu.URL
	}

	return MapController(pathsToUrls, fallback), nil
}
