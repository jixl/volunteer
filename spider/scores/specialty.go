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

type SpecialtyResponse struct {
	category
	Total  total                   `json:"totalRecord"`
	Scores []models.SpecialtyScore `json:"school"`
}

func Specialty(cate string, year string, page int) {
	if page <= 0 {
		page = firstPage
	}
	if cate == "all" {
		oneSpecialty(category{year, page, "", 0})
		return
	}

	kinds := [...]string{"文学类", "理学类", "哲学类", "教育学类", "管理学类",
		"经济学类", "农学类", "工学类", "医学类", "历史学类", "艺术学类", "交通运输类",
		"生化与药品类", "资源开发与测绘类", "材料与能源类", "土建类", "水利类", "制造类",
		"电子信息类", "环保、气象与安全类", "财经类", "医药卫生类", "旅游类", "公共事业类",
		"文化教育类", "艺术设计传媒类", "公安类", "轻纺食品类", "法律类",
	}

	for i := 0; i < len(kinds); i++ {
		oneSpecialty(category{year, page, kinds[i], 0})
	}
}
func oneSpecialty(cate category) {
	objResp := SpecialtyResponse{category: cate}
	count, err := spiderScores(objResp)
	if err != nil {
		log.Println("ERROR:", err)
	}

	nums := (count / 50) + 1
	log.Println("Specialty", cate, nums, count)

	for index := cate.page + 1; index <= nums; index++ {
		cate.page = index
		cate.execs = 0
		objResp = SpecialtyResponse{category: cate}
		count, _ = spiderScores(objResp)
		log.Println("Specialty", cate, count)
	}
}

func (s SpecialtyResponse) getURI() string {
	uri := "http://data.api.gkcx.eol.cn/soudaxue/querySpecialtyScore.html?messtype=json&size=50&page="
	var uriBuf bytes.Buffer
	uriBuf.WriteString(uri)
	uriBuf.WriteString(strconv.Itoa(s.page))
	uriBuf.WriteString("&fsyear=")
	uriBuf.WriteString(s.year)
	uriBuf.WriteString("&zytype=")
	uriBuf.WriteString(url.QueryEscape(s.category.name))

	return uriBuf.String()
}

func (data SpecialtyResponse) parse(resp *http.Response) (int, error) {
	defer resp.Body.Close()
	err := json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Println("ERROR DECODE:", err)
		return 0, err
	}
	for _, v := range data.Scores {
		v.Zytype = data.category.name
		models.Save(&v)
	}

	return strconv.Atoi(data.Total.Count)
}
