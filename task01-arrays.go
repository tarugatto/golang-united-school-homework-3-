package homework

var arr = [15]float32{2.5, 6, 3, 2, 1, 1.4, 3, 4, 1, 3, 3.3, 4.1, 5, 5.2, 1}
var _ = average(arr)

func average(input [15]float32) (result float32) {
	sum := 0
	for i := 0; i < len(input); i++ {
		sum += i
	}
	return float32(sum / len(input))
}
