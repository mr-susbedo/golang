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

// Deprecated: Use the strings.Repeat() instead, my dumbass was just playing with it (bad approach below)
func PrintChar(c rune, len uint8) string {
	var str string
	for i := 0; i < int(len); i++ {
		str += string(c)
	}
	return str
}
