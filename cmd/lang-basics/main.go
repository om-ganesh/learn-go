package main

import (
	"fmt"

	"github.com/om-ganesh/goutils"
)

func main() {
	fmt.Println("Learn importing modules")
	fmt.Println(goutils.Hello())
	str := "Hello Kul Umesh !!!"
	fmt.Println("AlphaNumeric ", str, "? ", goutils.IsAlphanumeric(str))
}
