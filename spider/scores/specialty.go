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
	gSpecialty struct {
		// Schools []interface{} `json:"school"`
		Schools []models.SpecialtyScore `json:"school"`
		Total   total                  `json:"totalRecord"`
	}
)

func Specialty() {
	fPage := 1
	// getScores(fPage)
	count, err := getSpecialtyScores(fPage)
	if err != nil {
		log.Println("ERROR:", err)
	}

	nums := (count / 50) + 1
	log.Println(nums, count)
	for index := fPage + 1; index <= nums; index++ {
		// fmt.Println(index, nums, count)
		getSpecialtyScores(index)
	}
}

func ParseSpecialtyScores(resp *http.Response) (int, error) {
	defer resp.Body.Close()
	var gr gSpecialty
	err := json.NewDecoder(resp.Body).Decode(&gr)
	if err != nil {
		log.Println("ERROR DECODE:", err)
		return 0, err
	}
	log.Println("Specialty:", sRecord, gr.Total)
	log.Println(gr.Schools)
	for _, v := range gr.Schools {
		models.AddSpecialtyScore(v)
	}
	return strconv.Atoi(gr.Total.Count)
}

var s_URI = "http://data.api.gkcx.eol.cn/soudaxue/querySpecialtyScore.html?messtype=json&size=50&page="
var sRecord = eRecord{0, 0}
func getSpecialtyScores(page int) (int, error) {
	sleep()
	var uriBuf bytes.Buffer
	uriBuf.WriteString(s_URI)
	uriBuf.WriteString(strconv.Itoa(page))
	resp, err := http.Get(uriBuf.String())

	if err != nil {
		log.Println("ERROR GET:", uriBuf.String(), err)
		if sRecord.isExit(page) {
			os.Exit(-1)
		}
		return getSpecialtyScores(page)
	}

	return ParseSpecialtyScores(resp)
}
