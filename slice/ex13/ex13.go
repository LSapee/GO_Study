package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name string
	Age  int
}

type Students []Student

func (s Students) Len() int           { return len(s) }
func (s Students) Less(i, j int) bool { return s[i].Age < s[j].Age }

func (s Students) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func main() {
	s := []Student{
		{"화랑", 31}, {"백두산", 52}, {"류", 42}, {"켄", 38}, {"송하나", 18},
	}

	sort.Sort(Students(s))
	fmt.Println(s)
	//함수 선언할때 주의 첫글자  ※ 대문자 = public
	//                   		※ 소문자 = praivate  이라 오류 날 수 있으니 오타 주의
}
