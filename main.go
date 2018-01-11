package main

import (
	// "fmt"
	// "io"
	// "io/ioutil"
	// "log"
	// "net/http"
	// "net/http/httptest"
	// "net/http/httputil"
	// "net/url"

	// "volunteer/spider"
	"volunteer/spider/scores"
)

func main() {
	// proxyServers := spider.GetProxyServer()
	// for k, v := range proxyServers {
	// 	fmt.Println(k, v)
	// 	spider.RmUnuse(k)
	// }
	// var uri = "http://data.api.gkcx.eol.cn/soudaxue/queryProvinceScore.html?messtype=json&size=50&page="
	// spider.ProxySpider(uri)
	// fmt.Println(spider.ProxyServers);
	// go scores.Province()
	// go scores.Specialty()
	// scores.Province()
	scores.Specialty()
	// proxy()
	// getIp("http://60.160.186.86:61234")
	// start()
}

type tgo struct {
	username string
	age      int
	address  string
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
