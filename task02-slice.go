package homework

var avrg = []int64{23, 45, 74, 92, 34, 66}
var slc1 = avrg[2:5]
var _ = reverse(slc1)

func reverse(input []int64) (result []int64) {
	var slc2 []int64

	for i := len(input) - 1; i >= 0; i-- {
		slc2 = append(slc2, input[i])
	}

	return slc2
}
