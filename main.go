package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	for i := 100; i < math.MaxInt64; i++ {
		str := strconv.Itoa(i)
		go checkAndPrint(str, i)
	}
}
func checkAndPrint(a string, b int) {
	x := len(a)
	frist := strings.Index(a, "5")
	last := strings.LastIndex(a, "5")
	i := 1
	j := x - 2
	if x == last-frist+1 && strings.Count(a, "0") == j-i+1 {
		fmt.Println(b)
	}
}
