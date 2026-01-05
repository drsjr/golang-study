package main

import(
	"strings"
	"fmt"
	"os"
)

func main() {
	// main_1()
	// main_2()
	main_3()
}

func main_1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func main_2() {
	s, sep := "", ""

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func main_3() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}