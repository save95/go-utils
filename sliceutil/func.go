package sliceutil

func UniqueUint(vals ...uint) []uint {
	set := make(map[uint]struct{}, len(vals))
	for _, val := range vals {
		set[val] = struct{}{}
	}

	res := make([]uint, 0, len(set))
	for val := range set {
		res = append(res, val)
	}

	return res
}

func UniqueString(vals ...string) []string {
	set := make(map[string]struct{}, len(vals))
	for _, val := range vals {
		set[val] = struct{}{}
	}

	res := make([]string, 0, len(set))
	for val := range set {
		res = append(res, val)
	}

	return res
}
