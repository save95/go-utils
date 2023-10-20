package locker

import "strings"

func Key(k string, ks ...string) string {
	if len(k) == 0 {
		k = "empty"
	}

	keys := []string{"lock", k}

	for _, s := range ks {
		keys = append(keys, s)
	}

	return strings.Join(keys, ":")
}
