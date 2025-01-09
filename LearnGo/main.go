package main

import (
	"fmt"
	"time"
)

func main() {
	// arr := []int{12, 432, 22, 14, 86}
	// var day string
	// arr = basic.SelectionSort(arr)
	// fmt.Println(arr)
	// fmt.Println("Nhập vào một ngày trong tuần (ví dụ: Monday, Tuesday, ...):")
	// fmt.Scanln(&day)
	// if day == "" {
	// 	fmt.Println("Không được để trống")
	// } else {
	// 	day = strings.ToLower(day)
	// 	basic.CheckDay(day)
	// }

	//goroutine va channel
	messageChannel := make(chan string)
	go say("Hello world", messageChannel)
	for i := 0; i < 5; i++ {
		message := <-messageChannel // goroutine nếu không lấy ra thì sẽ bị chặn
		fmt.Println(message)
	}
	// chương trình vẫn tiếp tục chạy nếu goroutine bị chặn
	fmt.Println("All messages received!")
}

func say(s string, ch chan string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		// Gửi thông điệp qua channel
		ch <- fmt.Sprintf("%s %d", s, i+1)
	}
}
