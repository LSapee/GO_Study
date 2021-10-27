package main

import (
	"bufio"
	"fmt"
	"os"
)

//파일 읽기
func ReadFile(filename string) (string, error) {
	//파일 열기 자세한건 >>  https://pkg.go.dev/os  참조 os.Open의 구조
	//func Open(name string)(*File, error)  반환하는 값이 2개 *File ,err 이다.
	file, err := os.Open(filename)
	//err 값 nil이 아닐경우 애러메시지 반환
	if err != nil {
		return "", err
	}
	// defer은 함수가 값을 반환하기 직전에 사용됨 마지막엔 파일을 닫아주기
	defer file.Close()
	// 파일 내용 읽기
	rd := bufio.NewReader(file)
	// ReadString은 값을 2개 반환하기 떄문에
	line, _ := rd.ReadString('\n')
	return line, nil
}

// 파일 생성
func WriteFile(filename string, line string) error {
	// os.Create의 반환값은 2개 *File  ,err 이다.
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = fmt.Fprint(file, line)
	return err
}

const filename string = "data.txt"

func main() {
	line, err := ReadFile(filename)
	if err != nil {
		err = WriteFile(filename, "This is WriteFile")
		if err != nil {
			fmt.Println("파일 생성에 실패했습니다.", err)
			return
		}
		line, err = ReadFile(filename)
		if err != nil {
			fmt.Println("파일 읽기에 실패했습니다.", err)
			return
		}
	}
	fmt.Println("파일 내용", line)
}
