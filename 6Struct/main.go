package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	//Case 1
	//fmt.Println(Vertex{1, 2})

	//Case 2
	// v := Vertex{1, 2}
	// fmt.Println(v.X)
	// v.X = 4
	// fmt.Println(v.X)

	//Case 3
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
