package protocols

import (
	"fmt"
)

func MyRaspberryPiReader() string {
	return "42"
}

func protocols() {
	fmt.Println(MyRaspberryPiReader())
}
