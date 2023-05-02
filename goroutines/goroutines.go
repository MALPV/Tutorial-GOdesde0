package goroutines

import (
	"fmt"
	"strings"
	"time"
)

func ShowStringLower(s string, channel chan bool) {

	separator := strings.Split(s, "")

	for _, char := range separator {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(char)
	}

	channel <- true

}