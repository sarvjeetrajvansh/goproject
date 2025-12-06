package main

import "fmt"

func main() {

	a := []int{1, 2, 3}
	b := make([]int, len(a))
	copy(b, a)
	a[0] = 12
	fmt.Println(a, b)
	s := make([]int, 0, 3)
	fmt.Println(len(s), cap(s))

	s = append(s, 1)
	s = append(s, 2)
	s = append(s, 3)
	fmt.Println(len(s), cap(s))

	s = append(s, 4)
	fmt.Println(len(s), cap(s))

	c := []int{1, 2, 3, 4}
	d := c[:2]
	d[0] = 99
	fmt.Println(c)
	fmt.Println(d)

}
