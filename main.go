package main

import (
	// "fmt"
	// "io"
	// "io/ioutil"
	"log"
	"net/http"
	"os"
	// "net/http/httptest"
	// "net/http/httputil"
	// "net/url"
	// "volunteer/spider"
	"github.com/jixl/volunteer/web"
)

func main() {
	// proxyServers := spider.GetProxyServer()
	// for k, v := range proxyServers {
	// 	fmt.Println(k, v)
	// 	spider.RmUnuse(k)
	// }
	// getIp("http://60.160.186.86:61234")
	startWeb()
}

func startWeb() {
	web.Routes()
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
		os.Exit(-1)
	}
}

// func proxy() {
// 	backendServer := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintln(rw, "this call was relayed by the reverse proxy")
// 	}))
// 	defer backendServer.Close()

// 	rpURL, err := url.Parse(backendServer.URL)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	frontendProxy := httptest.NewServer(httputil.NewSingleHostReverseProxy(rpURL))
// 	defer frontendProxy.Close()

// 	resp, err := http.Get(frontendProxy.URL)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	b, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Printf("%s", b)
// }
