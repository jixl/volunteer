package scores

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/jixl/volunteer/models"
)

type SpecialtyResponse struct {
	page     int
	category string
	Schools  []models.SpecialtyScore `json:"school"`
	Total    total                   `json:"totalRecord"`
}

func Specialty() {
	category := [...]string{"文学类", "理学类", "哲学类", "教育学类", "管理学类",
		"经济学类", "农学", "工学类", "医学类", "历史学类", "艺术学类", "交通运输类",
		"生化与药品类", "资源开发与测绘类", "材料与能源类", "土建类", "水利类", "制造类",
		"电子信息类", "环保、气象与安全类", "财经类", "医药卫生类", "旅游类", "公共事业类",
		"文化教育类", "艺术设计传媒类", "公安类", "轻纺食品类", "法律类",
	}

	for i := 0; i < len(category); i++ {
		oneSpecialty(category[i])
	}
}
func oneSpecialty(cate string) {
	fPage := 1
	sResp := SpecialtyResponse{page: fPage, category: cate}
	count, err := sResp.scores()
	if err != nil {
		log.Println("ERROR:", err)
	}

	nums := (count / 50) + 1
	log.Println(cate, nums, count)
	for index := fPage + 1; index <= nums; index++ {
		sResp = SpecialtyResponse{page: index, category: cate}
		sResp.scores()
	}
}

func (s SpecialtyResponse) getURI() string {
	// http://data-gkcx.eol.cn/soudaxue/querySpecialtyScore.html?messtype=json&page=30
	uri := "http://data.api.gkcx.eol.cn/soudaxue/querySpecialtyScore.html?messtype=json&size=50&page="
	var uriBuf bytes.Buffer
	uriBuf.WriteString(uri)
	uriBuf.WriteString(strconv.Itoa(s.page))
	uriBuf.WriteString("&zytype=")
	uriBuf.WriteString(url.QueryEscape(s.category))

	return uriBuf.String()
}

func (data *SpecialtyResponse) parse(resp *http.Response) (int, error) {
	defer resp.Body.Close()
	err := json.NewDecoder(resp.Body).Decode(data)
	if err != nil {
		log.Println("ERROR DECODE:", err)
		return 0, err
	}
	log.Println("Specialty:", sRecord, data.Total)
	for _, v := range data.Schools {
		v.Zytype = data.category
		v.Save()
	}
	return strconv.Atoi(data.Total.Count)
}

func (p *SpecialtyResponse) sleep() {
	sleep()
}

var sRecord = eRecord{0, ""}

func (p *SpecialtyResponse) scores() (int, error) {
	p.sleep()
	uri := p.getURI()
	resp, err := http.Get(uri)
	log.Println(uri)
	if err != nil {
		log.Println("ERROR GET: ", uri, err)
		if sRecord.isExit(uri) {
			os.Exit(-1)
		}
		return p.scores()
	}

	return p.parse(resp)
}
