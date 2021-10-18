package koreaPost

import "fmt"

type PostSender struct {
}

func (k *PostSender) Send(percel string) {
	fmt.Printf("우체국에서 택배 %v를 보냅니다.\n", percel)
}
