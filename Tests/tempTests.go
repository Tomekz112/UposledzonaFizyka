package main

import "fmt"

func main() {
	c(b)
}

func c(test func(string, string)) {
	s := ""
	for {
		fmt.Scanln(&s)
		test(s, "comrade")
	}
}

func b(text, from string) {
	fmt.Println(text, "from", from, "!")
}
