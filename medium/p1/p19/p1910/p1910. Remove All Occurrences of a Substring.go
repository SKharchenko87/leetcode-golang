package p1910

import (
	"reflect"
	"slices"
	"strings"
)

func removeOccurrences(s string, part string) string {
	i := strings.Index(s, part)
	for i != -1 {
		s = s[:i] + s[i+len(part):]
		i = strings.Index(s, part)
	}
	return s
}

func removeOccurrences0(s string, part string) string {
	bytes := []byte(s)
	partBytes := []byte(part)
	for i := 0; i < len(bytes)-len(part)+1; i++ {
		if reflect.DeepEqual(bytes[i:i+len(part)], partBytes) {
			bytes = slices.Delete(bytes, i, i+len(part))
			i -= len(part)
			if i < 0 {
				i = -1
			}
		}
	}
	return string(bytes)
}
