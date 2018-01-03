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

type (
	gResponse struct {
		// Schools []interface{} `json:"school"`
		Schools []models.ProvinceScore `json:"school"`
		Total   total                 `json:"totalRecord"`
	}
)

func Province() {
	fPage := 1
	// getScores(fPage)
	count, err := getScores(fPage)
	if err != nil {
		log.Println("ERROR:", err)
	}

	nums := (count / 50) + 1
	log.Println(nums, count)
	for index := 1; index <= nums; index++ {
		// fmt.Println(index, nums, count)
		getScores(index)
	}
}

func ParseProvinceScores(resp *http.Response) (int, error) {
	defer resp.Body.Close()
	var gr gResponse
	err := json.NewDecoder(resp.Body).Decode(&gr)
	if err != nil {
		log.Println("ERROR DECODE:", err)
		return 0, err
	}
	log.Println(gr.Total)
	// fmt.Println(gr.Schools)
	for _, v := range gr.Schools {
		models.AddProvinceScore(v)
	}
	return strconv.Atoi(gr.Total.Count)
}

var p_URI = "http://data.api.gkcx.eol.cn/soudaxue/queryProvinceScore.html?messtype=json&size=50&page="

var pRecord = eRecord{0, 0}
func getScores(page int) (int, error) {
	sleep()
	var uriBuf bytes.Buffer
	uriBuf.WriteString(p_URI)
	uriBuf.WriteString(strconv.Itoa(page))
	resp, err := http.Get(uriBuf.String())
	if err != nil {
		log.Println("ERROR GET:", uriBuf.String(), err)
		if isExit(page) {
			os.Exit(-1)
		}
		return getScores(page)
	}
	return ParseProvinceScores(resp)
}
