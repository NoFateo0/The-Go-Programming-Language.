package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	f, err := os.OpenFile("E:\\Sources\\The_Go_PL\\ozon.txt", os.O_APPEND|os.O_RDWR, 0777)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	for range os.Args[1:] {
		f.WriteString(<-ch)
	}
	fmt.Printf("%.2fs elapsed", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	fmt.Printf("%.2fs  %7d  %s\n", secs, nbytes, url)
	ch <- fmt.Sprintf("%.2fs  %7d  %s\n", secs, nbytes, url)
}
