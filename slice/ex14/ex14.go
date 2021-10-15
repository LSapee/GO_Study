package main

import (
	"fmt"
	"sort"
)

// 예제1
// func main() {
// 	array := [5]int{1, 2, 3, 4, 5}
// 	slice := array[1:3]

// 	slice = append(slice, 100)
// 	fmt.Println(array)

// 	//[1,2,3,100,5]
// }
// 예제 2
// func main() {
// 	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	slice = slice[:len(slice)-2]

// 	fmt.Println(slice)

// }
// 예제3
// func main() {
// 	slice := []int{1, 2, 3, 4, 5, 6}
// 	t, slice := slice[len(slice)-1], slice[:len(slice)-1]

// 	fmt.Println(t, slice)

// 	//6 1,2,3,4,5,
// }
//예제 4

type Profile struct {
	Name string
	Age  int
	Gool int
	Pass float64
}

type Profiles []Profile

func (p Profiles) Len() int           { return len(p) }
func (p Profiles) Less(i, j int) bool { return p[i].Gool < p[j].Gool }
func (p Profiles) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	p := []Profile{
		{"나통키", 13, 45, 78.4},
		{"오맹태", 16, 24, 67.4},
		{"오동도", 18, 54, 50.8},
		{"황금산", 16, 36, 89.7},
	}
	sort.Sort(sort.Reverse(Profiles(p)))
	fmt.Println(p)
}
