package scores

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"volunteer/models"
)

type ProvinceResponse struct {
	Schools []models.ProvinceScore `json:"school"`
	Total   total                  `json:"totalRecord"`
}

func Province() {
	fPage := 1
	pResp := ProvinceResponse{}
	count, err := pResp.scores(fPage)
	if err != nil {
		log.Println("ERROR:", err)
	}

	nums := (count / 50) + 1
	log.Println(nums, count)

	for index := 1; index <= nums; index++ {
		pResp = ProvinceResponse{}
		pResp.scores(index)
	}
}

func (p ProvinceResponse) getURI() string {
	// http://data-gkcx.eol.cn/soudaxue/queryProvinceScore.html?messtype=json&page=30
	return "http://data.api.gkcx.eol.cn/soudaxue/queryProvinceScore.html?messtype=json&size=50&page="
}

func (data *ProvinceResponse) parse(resp *http.Response) (int, error) {
	defer resp.Body.Close()
	err := json.NewDecoder(resp.Body).Decode(data)
	if err != nil {
		log.Println("ERROR DECODE:", err)
		return 0, err
	}
	log.Println("Province", pRecord, data.Total)
	for _, v := range data.Schools {
		v.Save()
	}
	return strconv.Atoi(data.Total.Count)
}

func (p *ProvinceResponse) sleep() {
	sleep()
}

var pRecord = eRecord{0, 0}

func (p *ProvinceResponse) scores(page int) (int, error) {
	p.sleep()

	var uriBuf bytes.Buffer
	uriBuf.WriteString(p.getURI())
	uriBuf.WriteString(strconv.Itoa(page))
	resp, err := http.Get(uriBuf.String())
	if err != nil {
		log.Println("ERROR GET: ", uriBuf.String(), err)
		if pRecord.isExit(page) {
			os.Exit(-1)
		}
		return p.scores(page)
	}

	return p.parse(resp)
}
