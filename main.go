package main

import (
	"fmt"
	"github.com/marzvix/stringmod/quotes"
	"github.com/marzvix/stringmod/strings"
)

func main() {
	o, e := strings.CountOddEven("12345")
	fmt.Println(o, e) // 3 2
	fmt.Println(quotes.GetEmoji("turtle"))
}
