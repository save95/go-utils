package sliceutil

import "strings"

func Join2String(sep string, strs ...string) string {
	texts := make([]string, 0)
	for _, str := range strs {
		if len(str) > 0 {
			texts = append(texts, str)
		}
	}
	return strings.Join(texts, "-")
}
