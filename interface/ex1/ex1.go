package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
	Name string
	Age  int
}

func (s Student) String() string {
	return fmt.Sprintf("안녕! 나는%d 살 %s라고 해", s.Age, s.Name)
}

func main() {
	studnet := Student{"철수", 12}
	var stringer Stringer

	stringer = studnet

	fmt.Printf("%s\n", stringer.String())
}
