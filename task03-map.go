package homework

var mapy = map[int]string{4: "Kim", 1: "Jimmy", 3: "Sara", 0: "Emily", 2: "Steven"}
var _ = sortMapValues(mapy)

func sortMapValues(input map[int]string) (result []string) {
	slc := make([]string, len(input)) //empty slice with input lenght

	for i, elmnt := range input {
		slc[i] = elmnt
	}

	return slc
}
