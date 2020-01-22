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
	data := ""
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
		{"hello world", "request cancle when data is hello world"},
	} {
		t.Run(data[1], func(t *testing.T) {
			store := &SpyStore{response: data[0], t: t}
			svr := Server(store)

			request := httptest.NewRequest(http.MethodGet, "/", nil)

			cancellingCtx, cancel := context.WithCancel(request.Context())
			time.AfterFunc(5*time.Millisecond, cancel)
			request = request.WithContext(cancellingCtx)

			response := &SpyResponseWriter{}
			fmt.Print(response)

			svr.ServeHTTP(response, request)
			if response.written {
				t.Error("a response should not have been written")
			}
		})
	}

}
