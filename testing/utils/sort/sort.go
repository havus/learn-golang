package sort

func BubbleSortAsc(elements []int) {
	stillLooping := true

	for stillLooping {
		stillLooping = false

		for i := 0; i < (len(elements) - 1); i ++ {
			if elements[i] > elements[i+1] {
				stillLooping = true
				elements[i], elements[i+1] = elements[i+1], elements[i]
			}
		}
	}
}
