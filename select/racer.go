package racer

import (
	"fmt"
	"net/http"
	"time"
)

// 选择响应比较快的url
func Racer(a, b string) (winner string) {
	startA := time.Now()
	http.Get(a)
	aDuration := time.Since(startA)

	startB := time.Now()
	http.Get(b)
	bDuration := time.Since(startB)
	fmt.Println(a, aDuration)
	fmt.Println(b, bDuration)
	if aDuration < bDuration {
		return a
	}
	return b
}
