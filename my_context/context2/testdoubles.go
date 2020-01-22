package context2

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	t        *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			fmt.Println(c)
			select {
			/*
			在子Context被传递到的goroutine中，应该对该子Context的Done信道（channel）进行监控，
			一旦该信道被关闭（即上层运行环境撤销了本goroutine的执行），应主动终止对当前请求信息的处理，释放资源并返回。
			 */
			case <-ctx.Done():
				// 进入这里的概率不大, 因为最后的Done()会让所有go routine退出
				fmt.Println("spy store got cancelled xx")
				s.t.Log("spy store got cancelled xx")
			default:
				time.Sleep(9 * time.Millisecond)
				result += string(c)
			}
		}
		time.Sleep(100 * time.Millisecond)
		data <- result
	}()

	select {
	case <-ctx.Done():
		fmt.Println("spy store got cancelled too")
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}
