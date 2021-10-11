package main

import "fmt"

//구조체 strct
type name struct {
	name string
	age  int
	car  bool
}

func main() {
	//map =  mapname := map[keytype]valuetype{key:value,..}
	/*
			jk := map[string]string{"name": "jk", "age": "1"}
			for key, value := range j {
				fmt.Println(key, vale)
		}
			fmt.Println(k)
	*/

	jk := name{"jk", 28, false}
	fmt.Println(jk.age)

}
