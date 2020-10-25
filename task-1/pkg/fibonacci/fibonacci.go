package fibonacci

func fibonacci(n int, cache map[int]int) int {
	if val, ok := cache[n]; ok {
		return val
	}

	if n < 2 {
		return 1
	}

	num := fibonacci(n-1, cache) + fibonacci(n-2, cache)
	cache[n] = num
	return num
}

// Number Fibonacci retuns
func Number(num int) int {
	cache := make(map[int]int)
	result := fibonacci(num, cache)
	return result
}
