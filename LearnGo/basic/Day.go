// switch
package basic

import "fmt"

func CheckDay(i string) {
	switch i {
	case "saturday", "sunday":
		fmt.Println("Đây là ngày cuối tuần!")
	case "monday", "tuesday", "wednesday", "thursday", "friday":
		fmt.Println("Đây là ngày trong tuần.")
	default:
		fmt.Println("Ngày bạn nhập không hợp lệ.")
	}
}
