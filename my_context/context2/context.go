package context2

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())  // 应该把context传入, 让Fetch负责上下文管理
		if err != nil {  // 如果取消, 会抛出错误
			return // todo: log error however you like
		}

		fmt.Fprint(w, data)
	}
}
