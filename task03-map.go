package homework

var mapy = map[int]string{5: "Kim", 2: "Jimmy", 4: "Sara", 1: "Emily", 3: "Steven"}
var _ = sortMapValues(mapy)

func sortMapValues(input map[int]string) (result []string) {
	arr1 := make([]string, len(input)) //empty slice with input lenght

	for i := 0; i < len(input); i++ {
		arr1[i] = input[i+1]
	}

	return arr1
}
