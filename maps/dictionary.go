package maps

type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("word has existed")
	ErrWordDoesNotExist = DictionaryErr("word does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (dic Dictionary) Search(word string) (string, error) {
	// 通过 map[key] 的方式从 map 中获取值
	value, ok := dic[word]
	if ok {
		return value, nil
	}
	return "", ErrNotFound
}

func (dic Dictionary) Add(key string, value string) error {
	// dic是个引用, 因此不需要传递指针. Map 作为引用类型是非常好的，因为无论 map 有多大，都只会有一个副本。
	_, err := dic.Search(key)
	switch err {
	case ErrNotFound:
		dic[key] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (dic Dictionary) Update(word string, definition string) error {
	_, err := dic.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		dic[word] = definition
	default:
		return err
	}
	return nil
}

func (dic Dictionary) Delete(word string) {
	delete(dic, word)
}
