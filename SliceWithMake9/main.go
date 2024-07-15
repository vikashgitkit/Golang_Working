package main

import "fmt"

func main() {

	//Case1- Slice With make
	// a := make([]int, 5)
	// printSlice("a", a)

	// b := make([]int, 0, 5)
	// printSlice("b", b)

	// c := b[:2]
	// printSlice("c", c)

	// d := c[2:5]
	// printSlice("d", d)

	//Case2- Slice of Slice
	// Create a tic-tac-toe board.
	// board := [][]string{
	// 	[]string{"_", "_", "_"},
	// 	[]string{"_", "_", "_"},
	// 	[]string{"_", "_", "_"},
	// }

	// // The players take turns.
	// board[0][0] = "X"
	// board[2][2] = "O"
	// board[1][2] = "X"
	// board[1][0] = "O"
	// board[0][2] = "X"

	// for i := 0; i < len(board); i++ {
	// 	fmt.Printf("%s\n", strings.Join(board[i], " "))
	// }

	//Case3- Appending Slice
	// var s []int
	// printSlice(s)

	// //append works on nil slices
	// s = append(s, 0)
	// printSlice(s)

	// //The slice grows as needed
	// s = append(s, 1)
	// printSlice(s)

	// //We can add more than one element at a time
	// s = append(s, 2, 3, 4)
	// printSlice(s)

	//Case4- Range
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

}

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

// func printSlice(s string, x []int) {
// 	fmt.Printf("%s len=%d cap=%d %v\n",
// 		s, len(x), cap(x), x)
// }

// For Case3
// func printSlice(s []int) {
// 	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
// }
