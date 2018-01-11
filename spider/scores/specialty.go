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

type SpecialtyResponse struct {
	Schools []models.SpecialtyScore `json:"school"`
	Total   total                   `json:"totalRecord"`
}

func Specialty() {
	fPage := 1
	sResp := SpecialtyResponse{}
	count, err := sResp.scores(fPage)
	if err != nil {
		log.Println("ERROR:", err)
	}

	nums := (count / 50) + 1
	log.Println(nums, count)
	for index := fPage + 1; index <= nums; index++ {
		sResp = SpecialtyResponse{}
		sResp.scores(index)
	}
}

func (s SpecialtyResponse) getURI() string {
	return "http://data.api.gkcx.eol.cn/soudaxue/querySpecialtyScore.html?messtype=json&size=50&page="
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
		v.Save()
	}
	return strconv.Atoi(data.Total.Count)
}

func (p *SpecialtyResponse) sleep() {
	sleep()
}

var sRecord = eRecord{0, 0}
func (p *SpecialtyResponse) scores(page int) (int, error) {
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
