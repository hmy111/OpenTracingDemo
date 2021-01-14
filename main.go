package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main()  {
	port:=8080
	addr:=fmt.Sprintf(":%d", port)
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/home", homeHandler)
	mux.HandleFunc("/async", serviceHandler)
	mux.HandleFunc("/service", serviceHandler)
	mux.HandleFunc("/db", dbHandler)
	fmt.Printf("Go to http://localhost:%d/home to start a request!\n", port)
	log.Fatal(http.ListenAndServe(addr, mux))
}
// Acts as our index page
func indexHandler(w http.ResponseWriter, r *http.Request)  {
	_, _ = w.Write([]byte(`<a href="/home"> 点击这里开始一个请求 </a>`))
}
func homeHandler(w http.ResponseWriter, r *http.Request){
	_,_ = w.Write([]byte("请求开始了!"))

	go func() {
		http.Get("http://10.1.140.48:8080/async")
	}()
	http.Get("http://10.1.140.48:8080/service")
	time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
	w.Write([]byte("请求结束"))
}
// Mocks a service endpoint that makes a DB call
func serviceHandler(w http.ResponseWriter, r *http.Request) {
	// ...
	http.Get("http://10.1.140.48:8080/async")
	time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
	// ...
}
// Mocks a DB call
func dbHandler(w http.ResponseWriter, r *http.Request) {
	http.Get("http://10.1.140.48:8080/db")
	time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
	// here would be the actual call to a DB.
}