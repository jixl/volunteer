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

func ParseSpecialtyScores(resp *http.Response) (int, error) {
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

var s_URI = "http://data.api.gkcx.eol.cn/soudaxue/querySpecialtyScore.html?messtype=json&size=50&page="
var count = 0

func getSpecialtyScores(page int) (int, error) {
	sleep()
	var uriBuf bytes.Buffer
	uriBuf.WriteString(s_URI)
	uriBuf.WriteString(strconv.Itoa(page))
	resp, err := http.Get(uriBuf.String())
	if err != nil {
		count++
		log.Println("ERROR GET:", uriBuf.String(), err)
		if count >= 3 {
			os.Exit(-1)
		}
		return getSpecialtyScores(page)
	}
	return ParseSpecialtyScores(resp)
}
