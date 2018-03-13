package main

import (
	"net/http"
	"fmt"
	"log"
	"sync"
)

var mu sync.Mutex

var count int

func main() {
	http.HandleFunc("/",handler)
	//http.HandleFunc("/count",counter)
	log.Fatal(http.ListenAndServe("localhost:8000",nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	//mu.Lock()
	//count ++
	//mu.Unlock()
	//fmt.Fprintf(w,"URL.PATH = %q \n",r.URL.Path)

	fmt.Fprintf(w,"method:%s url:%s proto:%s\n",r.Method,r.URL,r.Proto)

	for k,v := range r.Header {
		fmt.Fprintf(w,"header[%q] = %q \n",k,v)
	}

	fmt.Fprintf(w, "Host = %q\n", r.Host)

	fmt.Fprintf(w, "RemoteArr = %q\n", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k,v := range r.Form {
		fmt.Fprintf(w,"form[%q] = %q \n",k,v)
	}

}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w,"count : %d \n",count)
	mu.Unlock()
}
