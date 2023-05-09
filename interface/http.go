package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"strconv"
)

type dollars float32

func (d dollars)String()string{return fmt.Sprintf("$%.2f", d)}

type database struct{
	data map[string]dollars
	mutex sync.Mutex
}

func(db database)list(w http.ResponseWriter, req *http.Request){
	for item, price := range db.data{
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func(db database)price(w http.ResponseWriter, req *http.Request){
	item := req.URL.Query().Get("item")
	price, ok := db.data[item]
	if !ok{
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func(db database)update(w http.ResponseWriter, req *http.Request){
	db.mutex.Lock()
	defer db.mutex.Unlock()

	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	p, err := strconv.ParseFloat(price, 64)
	if err != nil {
		fmt.Fprintf(w, "invalid price")
	} else {
		db.data[item] = dollars(p)
		fmt.Fprintf(w, "update price success")
	}
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	item := req.URL.Query().Get("item")
	delete(db.data, item)
	fmt.Fprintf(w, "delete success")
}

func main() {
	db := database{map[string]dollars{"shoes": 50, "socks": 5}, sync.Mutex{}}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}