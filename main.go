package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {
	fmt.Println("starting up rproxy, the super-simple reverse proxy...")

	toHost := os.Getenv("RPROXY_TO_HOST")
	host, err := url.Parse(toHost)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to parse RPROXY_TO_HOST env var: %v", err))
		os.Exit(1)
	}

	proxy := httputil.NewSingleHostReverseProxy(host)

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		req.Host = host.Host
		req.URL.Host = host.Host
		proxy.ServeHTTP(res, req)
	})

	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println(fmt.Errorf("unable to start server: %v", err))
		os.Exit(1)
	}
}
