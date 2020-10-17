package fibonacci

func fibonacci() func() int {
	before, after := 0, 1
	return func() int {
		result := before
		before, after = after, before+after
		return result
	}
}

// Number Fibonacci retuns
func Number(num int) int {
	fibonacciFunction := fibonacci()

	for ; num != 1; num-- {
		fibonacciFunction()
	}

	return fibonacciFunction()
}
