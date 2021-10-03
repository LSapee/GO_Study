package main

// import "fmt"

// func main() {
// 	//배열 array 미리 크기를 정해야함
// 	names := [5]string{"jk", "ho", "lan"}
// 	//슬라이스  slice 미리 크기를 정하지 않아도 됨  append는 slice에 요소를 추가할때 사용
// 	names2 := []string{}
// 	//for문  돌아가면서 names에 요소를 하나씩 names2에 추가해주기
// 	for i := 0; i < len(names); i++ {
// 		fmt.Println(names[i])
// 		names2 = append(names2, names[i])
// 	}
// 	// for문 이후에 names의 lan을 aaa로 변경 후 4,5번 요소를 추가
// 	names[2] = "aaa"
// 	names[3] = "qqq"
// 	names[4] = "fuck"
// 	fmt.Println(names, names2)
// }
