package main

import (
	"fmt"
)


func ExploreArray() {
	// var a [3]int
	// a := [3]int{1, 2, 3} // Array of len 3
	a := [...]int{1, 2, 3} // Array of len 3
	fmt.Printf("Index 1: %v\n", a[1])
	fmt.Printf("Index 2: %v\n", a[2])
	fmt.Printf("Len: %v\n", len(a))

	b := a // Copies array.
	b[1] = 20
	fmt.Printf("Array a: %v\n", a)
	fmt.Printf("Array b: %v\n", b)

	c := &a // Pointer to array a.
	c[1] = 20
	fmt.Printf("Array a: %v\n", a)
	fmt.Printf("c -> a: %v\n", c)

	var matrix [3][3]int
	matrix[0] = [3]int{1, 0, 0}
	matrix[1] = [3]int{0, 1, 0}
	matrix[2] = [3]int{0, 0, 1}
	fmt.Printf("Matrix: %v", matrix)
}


func ExploreSlices() {
	a := []int{1, 2, 3} // Slice
	fmt.Printf("Len: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))

	b := a // Same slice.
	b[1] = 20
	fmt.Printf("Slice a: %v\n", a)
	fmt.Printf("b -> a: %v\n", b)

	c := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // Slice
	d := c[:]
	e := c[3:]
	f := c[:6]
	g := c[3:6]  // Inclusive 3, exclusive 6.

	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	h := make([]int, 3, 100)
	fmt.Println(h)
	fmt.Printf("Len: %v\n", len(h))
	fmt.Printf("Capacity: %v\n", cap(h))

	k := []int{}
	fmt.Println(k)
	fmt.Printf("Len: %v\n", len(k))
	fmt.Printf("Capacity: %v\n", cap(k))
	// k = append(k, 1, 2, 3, 4, 5) 
	k = append(k, []int{1, 2, 3, 4, 5}...)
	fmt.Println(k)
	fmt.Printf("Len: %v\n", len(k))
	fmt.Printf("Capacity: %v\n", cap(k))

	l := []int{1, 2, 3, 4, 5}
	m := l[3:]  // Pop an element.
	fmt.Println(m)

	n := []int{1, 2, 3, 4, 5}
	fmt.Println(n)
	o := append(n[:2], n[3:]...) // Remove an element at 4.
	fmt.Println(o)
	fmt.Println(n)
}


func main() {
	// ExploreArray()
	ExploreSlices()
}
