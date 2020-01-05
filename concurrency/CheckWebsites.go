package concurrency

// WebsiteChecker checks a url, returning a bool
type WebsiteChecker func(string) bool // 表明WebsiteChecker是一个函数, 它的签名是: func(string) bool

type result struct {
	name  string
	value bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultsChannel := make(chan result)

	for _, url := range urls {
		go func(u string) { // 匿名函数
			resultsChannel <- result{u, wc(u)} // 检查每个网址
		}(url)
	}
	for i := 0; i < len(urls); i++ {
		result := <-resultsChannel
		results[result.name] = result.value
	}
	return results
}
