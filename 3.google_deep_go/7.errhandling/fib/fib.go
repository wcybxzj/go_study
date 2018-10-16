package fib

func Fibonnaci() func() int {
	a, b = 0, 1
	return func() int {
		a, b = b, a+b
	}
}
