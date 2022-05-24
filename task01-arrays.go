package homework

//array
var arr = [15]float32{2.5, 6, 3, 2, 1, 1.4, 3, 4, 1, 3, 3.3, 4.1, 5, 5.2, 1}
var _ = average(arr)

func average(input [15]float32) (result float32) {
	var sum float32 = 0
	for _, i := range input {
		sum += i
	}
	middle := sum / float32(len(input)) //find average
	return middle
}
