package main

type Dictionary map[string]string

func (dic Dictionary) Search(word string) string {
	// 通过 map[key] 的方式从 map 中获取值
	return dic[word]
}
