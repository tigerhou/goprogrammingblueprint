package main

import "fmt"

func main() {
	x := []int{2, 3, 4}
	y := make([]int, 4)
	z := copy(y, x)
	fmt.Print(x, y, z)

	maps := map[int]string{
		12: "test",
	}

	if key, ok := maps[12]; ok {
		fmt.Println(key)
	}
}
