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
		// 把Fetch到的东西写入w中
		fmt.Fprint(w, store.Fetch())
	}
}
