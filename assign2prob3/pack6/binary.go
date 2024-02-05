package pack6

func BinarySearch(sl []int, key int) int {
	low := 0
	high := len(sl) - 1
	for low <= high {
		mid := (low + high) / 2
		if key == sl[mid] {
			return mid
		} else if key > sl[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}

	}
	return -1
}
