package utils_app

func FindIndex(slice []int, value int) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}
