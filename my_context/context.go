package my_context

import (
	"fmt"
	"net/http"
)

type Store interface {
	Fetch() string
	Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		data := make(chan string, 1)
		go func() {
			data <- store.Fetch()
			fmt.Println("如果提前终止, 这里就不会运行到")
		}()

		select {
		case d := <-data:
			//把Fetch到的东西写入w中
			fmt.Fprint(w, d)
		case <-ctx.Done():
			store.Cancel()
		}
	}
}
