package urlshort

import (
	"net/http"
	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...
	return func(w http.ResponseWriter, r *http.Request) {
		if val, ok := pathsToUrls[r.URL.Path]; ok {
			http.Redirect(w, r, val, http.StatusTemporaryRedirect)
		}else{
			fallback.ServeHTTP(w, r)
		}
		
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...
	parsedYAML, err := ParseYAML(yml)
	if err != nil {
		return nil, err
	}
	pathMap := BuildMap(parsedYAML)
	return MapHandler(pathMap, fallback), err
}

// ParseYAML will parse the string to a slice of map, 
// each map has two entries -- path:val, url:val
func ParseYAML(yml []byte) (result []map[string]string, err error) {
	err = yaml.Unmarshal(yml, &result)
    return result,err
}

// BuildMap builds a map with key=path and val=url from the result of ParseYAML
func BuildMap(raw []map[string]string) (map[string]string) {
	result := make(map[string]string)
	for _, item := range(raw) { 
		result[item["path"]] = item["url"]
	}
	return result
}