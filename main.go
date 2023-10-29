package main

import (
    "github.com/go-rod/rod"

	"fmt"
)

func main() {
    u := "ws://127.0.0.1:9222/devtools/browser/a8c769c7-9af2-4643-8ea4-83f9553d99fa"
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