package main

import (
	"goproject/ch20/fedex"
	// "goproject/ch20/koreaPost"
)

func SendBook(name string, sender *fedex.FedexSender) {
	sender.Send(name)
}

func main() {
	// sender := &koreaPost.PostSender{}
	// SendBook("어린왕자", sender)
	// SendBook("그리스인 조르바", sender)

	//오류남
}
