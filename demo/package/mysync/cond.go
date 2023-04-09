package mysync

import (
	"fmt"
)

func Run() {

	// var mu sync.Mutex
	var ch = make(chan int, 3)

	for i := 0; i <= 8; i++ {
		go func(i int) {

		}(i)
	}

	fmt.Println(len(ch))
}
