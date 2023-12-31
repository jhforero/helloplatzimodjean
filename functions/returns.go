package main

import "fmt"

func sum(values ...int) int {
	total := 0

	for _, num := range values {
		total += num
	}

	return total
}

func printNames(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

func getValues(x int) (double int, triple int, quad int) {
	// return 2 * x, 3 * x, 4 * x
	double = 2 * x
	triple = 3 * x
	quad = 4 * x

	return
}

func main() {
	fmt.Println(sum(1, 2, 3, 4))
	printNames("Jean", "Helberth", "Forero", "Martinez")
	fmt.Println(getValues(2))
}
