package retry

// Fibonacci returns successive Fibonacci numbers starting from 1
func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

var MAXIMUM_INTERVAL_TIME = 300

// FibonacciNext returns next number in Fibonacci sequence greater than start
func FibonacciNext(start int) int {
	fib := Fibonacci()
	num := fib()
	for num <= start {
		num = fib()
	}

	if num > MAXIMUM_INTERVAL_TIME {
		return MAXIMUM_INTERVAL_TIME
	}

	return num
}
