package urlshort

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

type pathToURL struct {
	Path string
	URL  string
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		url := pathsToUrls[request.URL.Path]
		if url != "" {
			http.Redirect(response, request, url, http.StatusPermanentRedirect)
		} else {
			fallback.ServeHTTP(response, request)
		}
	})
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement parseYAML and buildMap
	parsedYaml, err := parseYAML(yml)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(parsedYaml)
	return MapHandler(pathMap, fallback), nil
}

func parseYAML(yml []byte) ([]pathToURL, error) {
	// TODO: Implement this...
	var pathsToUrls []pathToURL
	err := yaml.Unmarshal(yml, &pathsToUrls)
	return pathsToUrls, err
}

func buildMap(pathsToUrls []pathToURL) map[string]string {
	//TODO: Implement this...
	builtMap := make(map[string]string)
	for _, ptu := range pathsToUrls {
		builtMap[ptu.Path] = ptu.URL
	}
	return builtMap
}
