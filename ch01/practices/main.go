package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func _() {
	var s, sep string

	for _, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func _() {

	for i, arg := range os.Args[0:] {
		fmt.Printf("index=%d, arg=%s\n", i, arg)
	}
}

const repetitions = 10_000

func _() {
	// using for loop the echo running time
	start := time.Now()
	var s, sep string
	for i := 0; i < repetitions; i++ {

		for _, arg := range os.Args[0:] {
			s += sep + arg
			sep = " "
		}
	}

	fmt.Printf("%.3f s elapsed\n", time.Since(start).Seconds())

}

func p4() {
	// using string.Join running time
	start := time.Now()

	for i := 0; i < repetitions; i++ {
		strings.Join(os.Args[0:], " ")
	}
	fmt.Printf("%.3f s elapsed\n", time.Since(start).Seconds())
}

func main() {
	//p1()
	//p2()
	//p3()  // 3.720 s elapsed
	p4() // 0.001 s elapsed
}
