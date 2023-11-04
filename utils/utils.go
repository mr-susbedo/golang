package utils

import "strings"

type Protocols string

const (
	HTTP  Protocols = "http://"
	HTTPS Protocols = "https://"
)

func EnsurePrefix(s string, prefix Protocols) string {

	if !strings.HasPrefix(s, string(prefix)) {
		return string(prefix) + s
	}
	return s
}
