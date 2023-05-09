package main

import (
	"fmt"
	"log"
	"flag"
	"links"
)

type linklist struct{
	url string
	depth int
}
var maxDepth int

func crawl(link linklist)[]linklist {
	if link.depth >= maxDepth{
		return nil
	}
	fmt.Println(link.depth, link.url)
	urls, err := links.Extract(link.url)
	if err != nil {
		log.Print(err)
	}
	var list []linklist
	for _, url := range urls{
		list = append(list, linklist{url, link.depth + 1})
	}
	return list
}

//!+
func main() {
	flag.IntVar(&maxDepth, "depth", 3, "max crawl depth")
	flag.Parse()
	fmt.Println(maxDepth)
	worklist := make(chan []linklist)  // lists of URLs, may have duplicates
	unseenLinks := make(chan linklist) // de-duplicated URLs

	// Add command-line arguments to worklist.
	go func(){
		var list []linklist
		for _, url := range flag.Args(){
			list = append(list, linklist{url, 0})
		}
		worklist <- list
	}()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link.url] {
				seen[link.url] = true
				unseenLinks <- link
			}
		}
	}
}
