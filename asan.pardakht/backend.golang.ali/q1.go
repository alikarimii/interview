package main

import (
	"fmt"
)

func main() {
	orginal := []int{1, 2, 3}
	list := orginal[:2]
	// orginal = append(orginal, 4) comment/uncomment and run
	list = append(list, 2)
	// what is a result and why
	fmt.Println(orginal, list)
}
