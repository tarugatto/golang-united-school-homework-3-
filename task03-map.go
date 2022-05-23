package homework

import "sort"

var mapy = map[int]string{5: "Kim", 2: "Jimmy", 4: "Sara", 1: "Emily", 3: "Steven"}
var _ = sortMapValues(mapy)

func sortMapValues(input map[int]string) (result []string) {
	names := make([]int, 0, len(mapy))
	for n := range mapy {
		names = append(names, n)
	}

	sort.Ints(names)

	var arr [6]string

	for i, names := range names {
		arr[i+1] = mapy[names]
	}

	return
}
