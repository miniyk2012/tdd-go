package racer

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRace(t *testing.T) {
	fastURL := "http://www.baidu.com"
	slowURL := "http://www.sina.cn"

	want := fastURL
	got := Racer(slowURL, fastURL)

	if want != got {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}

func TestRacer(t *testing.T) {
	slowServer := makeTestServer(20 * time.Millisecond)
	fastServer := makeTestServer(0 * time.Millisecond)

	slowURL := slowServer.URL // http://127.0.0.1:52446
	fastURL := fastServer.URL // http://127.0.0.1:52447
	fmt.Printf("slowURL=%s\n", slowURL)
	fmt.Printf("fastURL=%s\n", fastURL)

	want := fastURL
	got := Racer(slowURL, fastURL)

	if want != got {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}

func TestSelectRead(t *testing.T) {
	start := time.Now()
	c := make(chan interface{})
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {

		time.Sleep(4 * time.Second)
		close(c)
	}()

	go func() {

		time.Sleep(3 * time.Second)
		ch1 <- 3
	}()

	go func() {

		time.Sleep(3 * time.Second)
		ch2 <- 5
	}()

	fmt.Println("Blocking on read...")
	select {
	case <-c:

		fmt.Printf("Unblocked %v later.\n", time.Since(start))

	case <-ch1:

		fmt.Printf("ch1 case...")
	case <-ch2:

		fmt.Printf("ch1 case...")
	default:

		fmt.Printf("default go...")
	}
}

func TestSelectWrite(t *testing.T) {
	// make(chan int) 是 unbuffered channel, send 之后 send 语句会阻塞执行，直到有人 receive 之后 send 解除阻塞，后面的语句接着执行。
	var ch1 chan int = make(chan int, 1)
	var ch2 chan int = make(chan int, 1)
	var chs = []chan int{ch1, ch2}
	var numbers = []int{1, 2, 3, 4, 5}

	getNumber := func(i int) int {
		fmt.Printf("numbers[%d]=%v\n", i, numbers[i])

		return numbers[i]
	}
	getChan := func(i int) chan int {
		fmt.Printf("chs[%d]=%v\n", i, chs[i])
		return chs[i]
	}

	select {
	case getChan(0) <- getNumber(2):

		fmt.Println("1th case is selected.")
	case getChan(1) <- getNumber(3):

		fmt.Println("2th case is selected.")
	default:

		fmt.Println("default!.")
	}

}

// make(chan int) 和 make(chan int, 1): https://www.jianshu.com/p/f12e1766c19f
func TestMakeChan(t *testing.T) {
	// make(chan int) 是 unbuffered channel, send 之后 send 语句会阻塞执行，
	// 直到有人 receive 之后 send 解除阻塞，后面的语句接着执行
	// make(chan int, 1) 是 buffered channel, 容量为 1。
	// 在 buffer 未满时往里面 send 值并不会阻塞， 只有 buffer 满时再 send 才会阻塞，
	// 所以执行到 c <- 0 时并不会阻塞 fmt.Println(a) 的执行，这时 a 可能是 "hello world" 也可能是空，
	// 看两个 goroutine 谁执行的更快

	var c = make(chan int, 1)
	var a string

	go func() {
		a = "hello world"
		<-c
	}()
	// 模拟耗时操作
	// sum := 0
	// for i := 0; i < 100000; i++ {
	// 	sum += i*23 + 45 - 7*4/2 - 99
	// }
	c <- 0
	fmt.Println(a)
}

func makeTestServer(delay time.Duration) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
	return server
}
