package homework

var arr = [15]float32{4.5, 8, 3, 7, 1, 6.4, 8, 4, 1, 3, 11, 7.1, 5, 9.6, 1}
var _ = average(arr)

func average(input [15]float32) (result float32) {
	sum := 0
	for i := 0; i < len(input); i++ {
		sum += i
	}
	return float32(sum)
}
