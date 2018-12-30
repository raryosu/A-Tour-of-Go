package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type FetchedUrl struct {
	url map[string]bool
	mux sync.Mutex
}

func (f *FetchedUrl) Add(url string) {
	f.mux.Lock()
	f.url[url] = true
	f.mux.Unlock()
}

func (f *FetchedUrl) Exists(url string) bool {
	f.mux.Lock()
	defer f.mux.Unlock()
	return f.url[url]
}

func Crawl(url string, depth int, fetcher Fetcher, fetched FetchedUrl) {
	if depth <= 0 {
		return
	}
	if fetched.Exists(url) {
		return
	}

	ch := make(chan []string)

	var fetch func(url string, ch chan []string)
	fetch = func(url string, ch chan []string) {
		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			close(ch)
			return
		}
		fmt.Printf("found: %s %q\n", url, body)
		ch <- urls
		close(ch)
	}
	go fetch(url, ch)

	fetched.Add(url)
	urls := <-ch
	for _, u := range urls {
		Crawl(u, depth-1, fetcher, fetched)
	}
	return
}

func main() {
	fetched := FetchedUrl{url: make(map[string]bool)}
	Crawl("https://golang.org/", 4, fetcher, fetched)
}

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
