package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/labstack/echo"
	"go/build"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

var x = build.NoGoError{"wawa"}
var y = filepath.SkipDir

type eror interface {
	Eror() string
}

func New(text string) eror {
	return &errorString{text}
}

// errorString is a trivial implementation of error.
type errorString struct {
	s string
}

func (e *errorString) Eror() string {
	return e.s
}

func main() {
	helper()

	x := errorString{"text"}
	a := New("牛逼23")
	fmt.Println(a)
	fmt.Println(x)
	fmt.Println(a.Eror())
	s := []string{"tuid"}
	fmt.Print(s[0] == "tuid")
	line := "tuid,age,sex"
	ss := strings.Split(line, ",")
	fmt.Print(ss[0] == "tuid")
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))

}

func helper() {
	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// navigate to a page, wait for an element, click
	var example string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://golang.org/pkg/time/`),
		// wait for footer element is visible (ie, page is loaded)
		chromedp.WaitVisible(`body > footer`),
		// find and click "Expand All" link
		chromedp.Click(`#pkg-examples > div`, chromedp.NodeVisible),
		// retrieve the value of the textarea
		chromedp.Value(`#example_After .play .input textarea`, &example),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Go's time.After example:\n%s", example)
}
