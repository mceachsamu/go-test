package main

import "fmt"

const (
	numIters = 100
	fizz     = 3
	buzz     = 5
)

func main() {
	// hello, ok := os.LookupEnv("HELLO")
	// if !ok {
	// 	log.Fatal("could not find env variable \"HELLO\"")
	// }

	// fmt.Printf("%v %v", test.Hello, hello)

	for i := 1; i <= numIters; i++ {
		var p string
		if divides(i, fizz) {
			p = p + ("Fizz")
		}
		if divides(i, buzz) {
			p = p + ("Buzz")
		}
		if !divides(i, fizz) && !divides(i, buzz) {
			p = p + fmt.Sprintf("%v", i)
		}
		println(p)
	}

	m := make(map[string]int)
	m["Fizz"] = 3
	m["Buzz"] = 5

	fizzBuzz(numIters, m)
}

func fizzBuzz(iters int, m map[string]int) {
	for i := 1; i < (iters); i++ {
		//if divides atleast one, dont print number
		div := false
		for k, v := range m {
			if divides(i, v) {
				div = true
				print(k)
			}
		}
		if div {
			println("")
		} else {
			println(i)
		}
	}
}

func divides(i int, x int) bool {
	return i%x == 0
}

func hello() float64 {
	return 0.0
}
