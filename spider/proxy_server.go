package spider

import (
	"fmt"
	"github.com/jixl/volunteer/spider/scores"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

var (
	servers = []string{
		"119.29.12.129:8888",
		"27.184.125.53:8118",
		"121.196.226.246:84",
		"116.199.2.196:82",
		"110.73.3.123:8123",
	}

	serverMap = map[int]string{}
	rmTbl     = map[int]string{}
)

func GetProxyServer() map[int]string {
	for i, v := range servers {
		serverMap[i] = v
	}
	return serverMap
}

func RmUnuse(key int) {
	value, exists := serverMap[key]
	if exists {
		rmTbl[key] = value
		delete(serverMap, key)
	}
	fmt.Printf("REMOVE SERVER: %d::%s\n", key, value)
}

func ProxySpider(uri string) {
	for _, puri := range GetProxyServer() {
		crawl(uri, puri)
	}
	fmt.Println("OK!!!")
}

/**
* 返回response
 */
func crawl(uri string, puri string) *http.Response {
	request, _ := http.NewRequest("GET", uri, nil)
	//随机返回User-Agent 信息
	request.Header.Set("User-Agent", getAgent())
	request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	request.Header.Set("Connection", "keep-alive")
	proxy, err := url.Parse("http://" + puri)
	//设置超时时间
	timeout := time.Duration(20 * time.Second)
	fmt.Printf("使用代理:%s\t%s\n", uri, proxy)
	client := &http.Client{}
	if puri != "local" {
		client = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxy),
			},
			Timeout: timeout,
		}
	}

	response, err := client.Do(request)
	if err != nil || response.StatusCode != 200 {
		fmt.Printf("line-99:遇到了错误-并切换ip %s:%s:%s \n", uri, proxy, err)
	}
	// defer response.Body.Close()
	scores.ParseProvinceScores(response)
	return response
}

/**
* 随机返回一个User-Agent
 */
func getAgent() string {
	agent := []string{
		"Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:50.0) Gecko/20100101 Firefox/50.0",
		"Opera/9.80 (Macintosh; Intel Mac OS X 10.6.8; U; en) Presto/2.8.131 Version/11.11",
		"Opera/9.80 (Windows NT 6.1; U; en) Presto/2.8.131 Version/11.11",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; 360SE)",
		"Mozilla/5.0 (Windows NT 6.1; rv:2.0.1) Gecko/20100101 Firefox/4.0.1",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; The World)",
		"User-Agent,Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_8; en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
		"User-Agent, Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Maxthon 2.0)",
		"User-Agent,Mozilla/5.0 (Windows; U; Windows NT 6.1; en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	len := len(agent)
	return agent[r.Intn(len)]
}
