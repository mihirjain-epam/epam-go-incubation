package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
)

const nonTlsHost = "http://godoc.org"
const tlsHost = "https://godoc.org"

var responseCache *cache.Cache

func init() {
	responseCache = cache.New(5*time.Minute, 10*time.Minute)
}

func isRequestCached(w http.ResponseWriter, r *http.Request) bool {
	cachedResponse, found := responseCache.Get(r.URL.Path)
	if found {
		fmt.Println("Loading from cache")
		fmt.Fprintf(w, cachedResponse.(string))
		return true
	}
	return false
}

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
	fmt.Println("url : ", url)
	fmt.Println("HOST :: ", r.Host)
	if isRequestCached(w, r) {
		return
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

func startServer(wg *sync.WaitGroup) (*http.Server, *http.Server) {
	nonTlsSrv := &http.Server{Addr: ":8080"}
	http.HandleFunc("/github.com/stretchr/testify/assert", proxyOrCacheRequest)
	go func() {
		defer wg.Done() // let main know we are done cleaning up

		// always returns error. ErrServerClosed on graceful close
		if err := nonTlsSrv.ListenAndServe(); err != http.ErrServerClosed {
			// unexpected error. port in use?
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	tlsSrv := &http.Server{
		Addr: ":9443",
		TLSConfig: &tls.Config{
			MinVersion:               tls.VersionTLS13,
			PreferServerCipherSuites: true,
		},
	}
	go func() {
		defer wg.Done() // let main know we are done cleaning up

		// always returns error. ErrServerClosed on graceful close
		if err := tlsSrv.ListenAndServeTLS("server.crt", "server.key"); err != http.ErrServerClosed {
			// unexpected error. port in use?
			log.Fatalf("ListenAndServeTLS(): %v", err)
		}
	}()

	return nonTlsSrv, tlsSrv
}
func main() {
	httpServerExitDone := &sync.WaitGroup{}
	startServer(httpServerExitDone)
}
