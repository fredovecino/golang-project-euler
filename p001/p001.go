package main

import (
	"fmt"
)

func main() {
	limit := 1000

	fmt.Println("Slow:", slow(limit))
	fmt.Println("Medium:", medium(limit))
	fmt.Println("Fast:", fast(limit))

}

func slow(limit int) int {
	sum := 0

	for i := 0; i < limit; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	return sum
}

func medium(limit int) int {
	sum := 0
	for i := 3; i < limit; i += 3 {
		sum += i
	}

	for i := 5; i < limit; i += 5 {
		sum += i
	}

	for i := 15; i < limit; i += 15 {
		sum -= i
	}

	return sum
}

func fast(limit int) float32 {
	top := float32(lastDiv(3, limit))
	sum := ((3 + top) / 2.0) * (top / 3)

	top = float32(lastDiv(5, limit))
	sum += ((5 + top) / 2.0) * (top / 5)

	top = float32(lastDiv(15, limit))
	sum -= ((15 + top) / 2.0) * (top / 15)

	return sum
}

func lastDiv(d, n int) int {
	for i := n - 1; i > 0; i-- {
		if i%d == 0 {
			return i
		}
	}
	return -1
}
