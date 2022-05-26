package homework

var mapy = map[int]string{5: "Kim", 2: "Jimmy", 4: "Sara", 1: "Emily", 3: "Steven"}
var _ = sortMapValues(mapy)

func sortMapValues(input map[int]string) (result []string) {
	arr1 := make([]string, len(input)) //empty slice with input lenght

	for i, elmnt := range input {
		arr1[i-1] = elmnt
	}

	return arr1
}
