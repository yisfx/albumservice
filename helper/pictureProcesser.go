package helper

import (
	"fmt"
	"time"
)

func PictureProcess() {
	go Out()
}

func Out() {
	var message = make(chan int)

	go func() {
		for {
			select {
			case message <- 21:
				fmt.Println("write")
			case s := <-message:
				fmt.Println(s)

			default:
				time.Sleep(time.Second * 2)
				fmt.Println("default")
			}
		}
	}()
	message <- 21
}
