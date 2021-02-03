package helper

import (
	"fmt"
	"time"
)

var message = make(chan string)

func PictureProcess() {
	fmt.Println("in PictureProcess")
	fmt.Println("<-0")
	message <- "hello1"
	fmt.Println("<-1")
	time.Sleep(time.Second * 2)
	fmt.Println(<-message)
}
