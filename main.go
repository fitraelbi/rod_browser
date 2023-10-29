package main

import (
    "github.com/go-rod/rod"

	"fmt"
)

func main() {
    u := "ws://127.0.0.1:9222/devtools/browser/f470edf7-0268-4f8f-8829-ccb85bf7b032"
    browser := rod.New().ControlURL(u).MustConnect()

	
	fmt.Println("Set Page")
	page := browser.MustPage("https://www.wikipedia.org/")
    fmt.Println("Navigating to page")
    page.MustWaitStable()
    fmt.Println(
        page.MustEval("() => document.title"),	
	)
	// page := browser.MustPage("http://api.ipify.org")
    // fmt.Println("Set title")
	// // IP address should be the same, since we are using local
	// // proxy, however the response signals that the proxy works
	// println(page.MustElement("html").MustText())
}