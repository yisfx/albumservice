package helper

import (
	"fmt"
	"time"
)

var message = make(chan string)

func In(album string) {

	message <- album
	// time.Sleep(time.Second * 7)
	// for {
	// 	if time.Now().Second()%3 == 0 {
	// 		message <- "ininin"
	// 	}
	// 	time.Sleep(time.Second * 7)
	// }
}

func Out() {
	for {
		select {
		case s := <-message:
			go buildAlbum(s)
		default:
			time.Sleep(time.Second * 20)
			fmt.Println("default")
		}
	}
}

func buildAlbum(album string) {
	albumHelper = NewAlbumHelper()
	albumHelper.
}
