package concurrency

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	// fmt.Println(url)
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

// CheckWebsite returns true if the URL returns a 200 status code, false otherwise
func CheckWebsite(url string) bool {
	response, err := http.Head(url)
	if err != nil {
		return false
	}

	if response.StatusCode != http.StatusOK {
		return false
	}

	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	t.Run("mockChecker", func(t *testing.T) {
		actualResults := CheckWebsites(mockWebsiteChecker, websites)

		want := len(websites)
		got := len(actualResults)

		if want != got {
			t.Fatalf("Wanted %v, got %v", want, got)
		}

		expectedResults := map[string]bool{
			"http://google.com":          true,
			"http://blog.gypsydave5.com": true,
			"waat://furhurterwe.geds":    false,
		}

		if !reflect.DeepEqual(actualResults, expectedResults) {
			t.Fatalf("Wanted %v, got %v", expectedResults, actualResults)
		}
	})
	t.Run("true_Checker", func(t *testing.T) {
		websites := []string{
			"http://google.com",
			"http://baidu.com",
			"waat://furhurterwe.geds",
		}
		actualResults := CheckWebsites(CheckWebsite, websites)
		fmt.Print(actualResults)
	})
}
