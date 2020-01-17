package racer

import (
	"fmt"
	"net/http"
	"time"
)

// Racer 返回选择响应比较快的url
func Racer(a, b string) (winner string, err error) {
	// select就是用来监听和channel有关的IO操作，当 IO 操作发生时，触发相应的动作。
	// ping(a), ping(b)依次执行, 由于ping中有协程不阻塞, 因此ping(a)很快就返回, 然后运行ping(b),
	// 也很快返回. 此时2个chan都没法读, 因此select阻塞, 直到某个协程跑完, close(c)后, c变得可读, 优先返回跑的快的url.
	// close(c)的文档: After the last value has been received
	// from a closed channel c, any receive from c will succeed without
	// blocking, returning the zero value for the channel element.
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(10 * time.Second):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}

	//atime := measureResponseTime(a)
	//fmt.Printf("%s=%v\n", a, atime)
	//measureResponseTime(b)
	//btime := measureResponseTime(b)
	//fmt.Printf("%s=%v\n", b, btime)
	//if atime < btime {
	//	return a
	//} else {
	//	return b
	//}
}

func ping(a string) chan struct{} {
	fmt.Println(a)
	c := make(chan struct{})
	go func() {
		http.Get(a)
		close(c)
	}()
	return c
}

func measureResponseTime(a string) time.Duration {
	startA := time.Now()
	http.Get(a)
	aDuration := time.Since(startA)
	return aDuration
}
