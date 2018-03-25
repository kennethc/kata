package util

// IsFizz checks if a given integer is divisible by 3
func IsFizz(x int) bool {
	if x%3 == 0 {
		return true
	}
	return false
}

// IsBuzz checks if a given integer is divisible by 5
func IsBuzz(x int) bool {
	if x%5 == 0 {
		return true
	}
	return false
}
