package maps

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")
func (dic Dictionary) Search(word string) (string, error) {
	// 通过 map[key] 的方式从 map 中获取值
	value, ok := dic[word]
	if ok {
		return value, nil
	}
	return "", ErrNotFound
}

func (dic Dictionary) Add(key string, value string) {
	// dic是个引用, 因此不需要传递指针. Map 作为引用类型是非常好的，因为无论 map 有多大，都只会有一个副本。
	dic[key] = value
}
