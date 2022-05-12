package utils

import (
	"regexp"
)

func ExtractParametricEndpoints(url string, prefix string) (string, bool) {
	pattern := prefix + "/{1}([a-zA-z0-9]+)/?"
	r, err := regexp.Compile(pattern)

	if err == nil {
		matches := r.FindStringSubmatch(url)

		if len(matches) == 1 || matches == nil {
			return "no key", false
		}

		p := matches[1]
		return p, true
	} else {
		return "error", false
	}
}

// func BasicAuth(username string) string {
// 	return base64.StdEncoding.EncodeToString([]byte(username))
// }
