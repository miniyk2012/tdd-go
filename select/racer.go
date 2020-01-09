package racer

import (
	"fmt"
	"net/http"
	"time"
)

// Racer 返回选择响应比较快的url
func Racer(a, b string) (winner string) {
	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)

	fmt.Println(a, aDuration)
	fmt.Println(b, bDuration)
	if aDuration < bDuration {
		return a
	}
	return b
}

func measureResponseTime(a string) time.Duration {
	startA := time.Now()
	http.Get(a)
	aDuration := time.Since(startA)
	return aDuration
}
