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
	category
	Total  total                  `json:"totalRecord"`
	Scores []models.ProvinceScore `json:"school"`
}

func Province() {
	kinds := [...]string{"综合类", "理工类", "农林类", "医药类", "语言类", "财经类",
		"医药类", "财经类", "政法类", "体育类", "艺术类", "民族类", "军事类", "其它",
	}

	for i := 0; i < len(kinds); i++ {
		oneProvince(category{firstPage, kinds[i], 0})
	}
}

func oneProvince(cate category) {
	objResp := ProvinceResponse{category: cate}
	count, err := spiderScores(objResp)
	if err != nil {
		log.Println("ERROR:", err)
	}

	nums := (count / 50) + 1
	log.Println(cate, nums, count)

	for index := cate.page + 1; index <= nums; index++ {
		cate.page = index
		cate.execs = 0
		objResp = ProvinceResponse{category: cate}
		count, _ = spiderScores(objResp)
		log.Println("Province", cate, count)
	}
}

func (obj ProvinceResponse) getURI() string {
	// http://data-gkcx.eol.cn/soudaxue/queryProvinceScore.html?messtype=json&page=30
	uri := "http://data.api.gkcx.eol.cn/soudaxue/queryProvinceScore.html?messtype=json&size=50&page="
	var uriBuf bytes.Buffer
	uriBuf.WriteString(uri)
	uriBuf.WriteString(strconv.Itoa(obj.category.page))
	uriBuf.WriteString("&schoolproperty=")
	uriBuf.WriteString(url.QueryEscape(obj.category.name))

	return uriBuf.String()
}

func (data ProvinceResponse) parse(resp *http.Response) (int, error) {
	defer resp.Body.Close()
	err := json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Println("ERROR DECODE:", err)
		return 0, err
	}
	for _, v := range data.Scores {
		v.SchoolType = data.category.name
		v.Save()
	}

	return strconv.Atoi(data.Total.Count)
}
