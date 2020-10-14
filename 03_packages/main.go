package main

import (
	"fmt"
	"math"

	"github.com/AntonyLe23/go_crash_course/03_packages/strutil"
)

func main() {
	fmt.Println("Hello world")
	fmt.Println(math.Floor(2.3))
	fmt.Println(math.Ceil(2.3))
	fmt.Println(math.Sqrt(16))
	fmt.Println(strutil.Reverse("12345"))
}
