package homework

var mapy = map[int]string{5: "Kim", 2: "Jimmy", 4: "Sara", 1: "Emily", 3: "Steven"}
var _ = sortMapValues(mapy)

func sortMapValues(input map[int]string) (result []string) {
	slc := make([]string, len(input)+1) //empty slice with input lenght

	for i, elmnt := range input {
		slc[i] = elmnt
	}

	slc1 := slc[1:6]

	return slc1
}
