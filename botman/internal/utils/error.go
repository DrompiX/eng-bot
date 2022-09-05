package utils

import "strings"

func ErrorToCapital(e error) string {
	if e := e.Error(); len(e) > 0 {
		return strings.ToUpper(e[:1]) + e[1:]
	}
	return ""
}
