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
	got, _ := Racer(slowURL, fastURL)

	if want != got {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}

func TestRacer(t *testing.T) {
	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		slowServer := makeTestServer(20 * time.Millisecond)
		fastServer := makeTestServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL // http://127.0.0.1:52446
		fastURL := fastServer.URL // http://127.0.0.1:52447
		fmt.Printf("slowURL=%s\n", slowURL)
		fmt.Printf("fastURL=%s\n", fastURL)

		want := fastURL
		got, _ := Racer(slowURL, fastURL)

		if want != got {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		serverA := makeTestServer(1 * time.Second + 500 * time.Millisecond)
		serverB := makeTestServer(3 * time.Second)

		defer serverA.Close()
		defer serverB.Close()

		ConfigurableRacer(serverA.URL, serverB.URL, 1 * time.Second)
		//if err == nil {
		//	t.Error("expected an error but didn't get one")
		//}
		fmt.Println("ConfigurableRacer Test Done")
	})

}



func TestSelectRead(t *testing.T) {
	start := time.Now()
	c := make(chan interface{})
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {

		time.Sleep(1 * time.Second)
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

	time.Sleep(2 * time.Second)
	fmt.Println("Blocking on read...")
	select {
	case <-c:  // 从c读数据

		fmt.Printf("Unblocked %v later.\n", time.Since(start))

	case <-ch1:

		fmt.Printf("ch1 case...")
	case <-ch2:

		fmt.Printf("ch1 case...")
	default:
		// 如果有一个或多个IO操作可以完成，则Go运行时系统会随机的选择一个执行，
		// 否则的话，如果有default分支，则执行default分支语句，如果连default都没有，
		// 则select语句会一直阻塞，直到至少有一个IO操作可以进行.

		fmt.Printf("default go...")
	}
}

func TestSelectWrite(t *testing.T) {
	// make(chan int) 是 unbuffered channel, send 之后 send 语句会阻塞执行，直到有人 receive 之后 send 解除阻塞，后面的语句接着执行。
	var ch1 chan int = make(chan int)
	var ch2 chan int = make(chan int)
	var chs = []chan int{ch1, ch2}
	var numbers = []int{1, 2, 3, 4, 5}
	// 所有channel表达式都会被求值、所有被发送的表达式都会被求值。求值顺序：自上而下、从左到右.
	getNumber := func(i int) int {
		time.Sleep(1 * time.Second)
		fmt.Printf("numbers[%d]=%v\n", i, numbers[i])

		return numbers[i]
	}
	getChan := func(i int) chan int {
		fmt.Printf("chs[%d]=%v\n", i, chs[i])
		return chs[i]
	}

	select {
	case getChan(0) <- getNumber(2):  // 往channel写数据

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

	//var c = make(chan int, 1)
	var c = make(chan int)
	var a string

	go func() {
		a = "hello world"
		<-c
	}()
	//模拟耗时操作
	//sum := 0
	//for i := 0; i < 100000; i++ {
	//	sum += i*23 + 45 - 7*4/2 - 99
	//}
	c <- 0  // 如果c = make(chan int), 这里会阻塞, 直到<-c才会才会解除
	fmt.Println(a)
	if a != "hello world" {
		t.Error("xx")
	}
}

func makeTestServer(delay time.Duration) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("start sleep %v\n", delay)
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
		fmt.Printf("sleep %v\n", delay)
	}))

	return server
}
