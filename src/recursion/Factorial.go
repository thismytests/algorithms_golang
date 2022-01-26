package recursion

func Factorial(numb int) int {
	if numb == 1 {
		return 1
	}

	return numb * Factorial(numb-1)
}
