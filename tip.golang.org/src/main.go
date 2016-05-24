package main

import (
	"log"
	"net/http"
)

func main() {
	p := new(Proxy)
	go p.run()
	http.Handle("/", p)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

type Proxy struct {
	mu    sync.Mutex
	proxy *httputil.ReverseProxy
}

func NewProxy() {

}

func (p *Proxy) run() {
	for {

	}
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/_tipstatus" {
		p.serveStatus(w, r)
		return
	}

	p.mu.Lock()

	proxy := p.proxy
	if proxy == nil {
		http.Error(w, "not ready", http.StatusInternalServerError)
		return
	}

	p.mu.Unlock()

	proxy.ServeHTTP(w, r)
}

func (p *Proxy) serveStatus(w http.ResponseWriter, r *http.Request) {

}
