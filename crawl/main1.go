package main

import (
	"fmt"
	"log"
	"sync"
	"flag"
	"links"
)

//!+sema
// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

var maxDepth int
var seen = make(map[string]bool)
// visit seen in goroutine, need a lock
var seenLock = sync.Mutex{}

func crawl(url string, depth int, wg *sync.WaitGroup){
	defer wg.Done()
	if depth >= maxDepth{
		return
	}
	fmt.Println(depth, url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token
	if err != nil {
		log.Print(err)
	}
	for _, link := range list{
		seenLock.Lock()
		if seen[link]{
			seenLock.Unlock()
			continue
		}
		seen[link] = true
		seenLock.Unlock()
		wg.Add(1)
		go crawl(link, depth+1, wg)
	}
}

func main() {
	flag.IntVar(&maxDepth, "depth", 3, "max crawl depth")
	flag.Parse()
	wg := &sync.WaitGroup{}
	for _, link := range flag.Args(){
		wg.Add(1)
		go crawl(link, 0, wg)
	}
	wg.Wait()
}