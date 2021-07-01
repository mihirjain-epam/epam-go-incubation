package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"
)

const nonTlsHost = "http://godoc.org"
const tlsHost = "https://godoc.org"

func proxyOrCacheRequest(w http.ResponseWriter, r *http.Request) {
	var url string
	if r.Host == "localhost:8080" {
		url = nonTlsHost + r.URL.Path
	} else if r.Host == "localhost:9443" {
		url = tlsHost + r.URL.Path
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// should only cache safe methods
	if r.Method != http.MethodGet && r.Method != http.MethodHead && r.Method != http.MethodOptions {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var responseCache = cache.New(5*time.Minute, 10*time.Minute)
	cachedResponse, found := responseCache.Get(r.URL.Path)
	fmt.Println("url : ", url)
	fmt.Println("HOST :: ", r.Host)
	if found {
		fmt.Println("Loading from cache")
		fmt.Fprintf(w, cachedResponse.(string))
	} else {
		response, err := http.Get(url)
		if err != nil || response.StatusCode != http.StatusOK {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		bodyStr := string(body)
		responseCache.Set(r.URL.Path, bodyStr, cache.DefaultExpiration)
		w.Header().Set("Cache-Control", "max-age=60")
		fmt.Println("Loading from proxy")
		fmt.Fprint(w, bodyStr)
	}
}

func main() {
	http.HandleFunc("/github.com/stretchr/testify/assert", proxyOrCacheRequest)
	go http.ListenAndServe(":8080", nil)
	err1 := http.ListenAndServeTLS(":9443", "RootCA.crt", "RootCA.key", nil)
	if err1 != nil {
		log.Fatal(err1)
	}

}
