package context2

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestContext2(t *testing.T) {
	data := "Hello world"
	t.Run("returns data from store", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}
	})

	for _, data := range [][]string{
		{"", "request cancle when data is empty data"},
		{"hello world lalallalal", "request cancle when data is hello world lalallalal"},
	} {
		t.Run(data[1], func(t *testing.T) {
			store := &SpyStore{response: data[0], t: t}
			svr := Server(store)

			request := httptest.NewRequest(http.MethodGet, "/", nil)
			grandfatherCtx := request.Context()

			fatherCtx, fatherCancel := context.WithCancel(grandfatherCtx)
			childCtx, _ := context.WithCancel(fatherCtx)
			time.AfterFunc(50*time.Millisecond, fatherCancel) // 终止fatherCtx, 也能终止childCtx
			request = request.WithContext(childCtx)

			response := &SpyResponseWriter{}
			fmt.Println(response)

			svr.ServeHTTP(response, request)
			if response.written {
				t.Error("a response should not have been written")
			}
			<-fatherCtx.Done()
			<-childCtx.Done()
			fmt.Println("终止")
		})
	}

}
