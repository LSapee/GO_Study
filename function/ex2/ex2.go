package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to create a file")
		return
	}

	defer fmt.Println("반드시 호출됩니다.")
	defer f.Close()
	defer fmt.Println("파일이 닫혔습니다.")

	fmt.Println("파일에 Hello world를 씁니다.")
	fmt.Fprintln(f, "Hello World")
}
