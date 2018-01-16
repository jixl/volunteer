package scores

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/jixl/volunteer/models"
)

type ProvinceResponse struct {
	page      int
	category  string
	execCount int
	Schools   []models.ProvinceScore `json:"school"`
	Total     total                  `json:"totalRecord"`
}

func Province() {
	category := [...]string{"综合类", "理工类", "农林类", "医药类", "语言类", "财经类",
		"医药类", "财经类", "政法类", "体育类", "艺术类", "民族类", "军事类", "其它",
	}

	for i := 0; i < len(category); i++ {
		oneProvince(category[i])
	}
}

func oneProvince(cate string) {
	fPage := 1
	pResp := ProvinceResponse{page: fPage, category: cate, execCount: 0}
	count, err := spiderScores(pResp)
	if err != nil {
		log.Println("ERROR:", err)
	}

	nums := (count / 50) + 1
	log.Println(cate, nums, count)

	for index := fPage + 1; index <= nums; index++ {
		pResp = ProvinceResponse{page: index, category: cate, execCount: 0}
		count, _ = spiderScores(pResp)
		log.Println("Province", index, count)
	}
}

func (obj ProvinceResponse) getURI() string {
	// http://data-gkcx.eol.cn/soudaxue/queryProvinceScore.html?messtype=json&page=30
	uri := "http://data.api.gkcx.eol.cn/soudaxue/queryProvinceScore.html?messtype=json&size=50&page="
	var uriBuf bytes.Buffer
	uriBuf.WriteString(uri)
	uriBuf.WriteString(strconv.Itoa(obj.page))
	uriBuf.WriteString("&schoolproperty=")
	uriBuf.WriteString(url.QueryEscape(obj.category))

	return uriBuf.String()
}

func (data ProvinceResponse) parse(resp *http.Response) (int, error) {
	defer resp.Body.Close()
	err := json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Println("ERROR DECODE:", err)
		return 0, err
	}
	for _, v := range data.Schools {
		v.SchoolType = data.category
		v.Save()
	}
	return strconv.Atoi(data.Total.Count)
}

func (obj ProvinceResponse) isExit() bool {
	obj.execCount++
	return obj.execCount >= 3
}
