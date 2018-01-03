package models

import (
	"encoding/json"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

const table = "province_score"

type ProvinceScore struct {
	Id            bson.ObjectId `bson:"_id,omitempty"`
	SchoolId      json.Number   `bson:"schoolid" json:"schoolid"`
	SchooleName   string        `bson:"schoolname" json:"schoolname"`
	LocalProvince string        `bson:"localprovince" json:"localprovince"`
	Province      string        `bson:"province" json:"province"`
	StudentType   string        `bson:"studenttype" json:"studenttype"`
	Year          json.Number   `bson:"year" json:"year"`
	Batch         string        `bson:"batch" json:"batch"`
	Var           json.Number   `bson:"var" json:"var"`
	VarScore      json.Number   `bson:"var_score" json:"var_score"`
	Max           interface{}   `bson:"max,omitempty" json:"max"`
	Min           interface{}   `bson:"min" json:"min"`
	Num           json.Number   `bson:"num" json:"num"`
	Fencha        json.Number   `bson:"fencha" json:"fencha"`
	ProvinceScore interface{}   `bson:"provincescore" json:"provincescore"`
	URL           string        `bson:"url" json:"url"`
}

/**
 * 添加ProvinceScore对象
 */
func AddProvinceScore(ps ProvinceScore) string {
	query := func(c *mgo.Collection) error {
		if isNeedReset(ps.Max) {
			ps.Max = "--"
		}
		if isNeedReset(ps.Min) {
			ps.Min = "--"
		}
		if isNeedReset(ps.ProvinceScore) {
			ps.ProvinceScore = nil
		}
		return c.Insert(ps)
	}

	err := witchCollection(table, query)
	if err != nil {
		log.Println("ERROR INSERT:", ps, err)
		return "false"
	}
	return "true"
}
