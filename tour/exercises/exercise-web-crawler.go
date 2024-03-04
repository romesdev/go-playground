package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type CrawlerState struct {
	mu sync.Mutex
	v  map[string]bool
}

var cs CrawlerState = CrawlerState{v: make(map[string]bool)}

func main() {
	exit := make(chan bool)
	go Crawl("https://golang.org/", 4, fetcher, exit)
	<-exit

}

func (cs *CrawlerState) Visit(key string) {
	cs.mu.Lock()
	// Lock so only one goroutine at a time can access the map sc.v.
	cs.v[key] = true
	cs.mu.Unlock()
}

// Value returns the current value of the crawler url for the given key.
func (cs *CrawlerState) Value(key string) bool {
	cs.mu.Lock()
	// Lock so only one goroutine at a time can access the map sc.v.
	defer cs.mu.Unlock()
	return cs.v[key]
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, exit chan bool) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:

	if depth <= 0 {
		exit <- true
		return
	}
	visited := cs.Value(url)

	if visited == false {
		cs.Visit(url)
	} else if visited == true {
		exit <- true
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		exit <- true
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	e := make(chan bool)
	for i := 0; i < len(urls); i++ {
		go Crawl(urls[i], depth-1, fetcher, e)
	}

	// wait for all child gorountines to exit
	for i := 0; i < len(urls); i++ {
		<-e
	}
	exit <- true
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
