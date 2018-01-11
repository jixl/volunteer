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
	Max           interface{}   `bson:"max" json:"max"`
	Min           interface{}   `bson:"min" json:"min"`
	Num           json.Number   `bson:"num" json:"num"`
	Fencha        json.Number   `bson:"fencha" json:"fencha"`
	ProvinceScore interface{}   `bson:"provincescore" json:"provincescore"`
	URL           string        `bson:"url" json:"url"`
}

func (bean ProvinceScore) getTableName() string { return "province_score" }

func (bean ProvinceScore) Save() string {
	if isNeedReset(bean.Max) {
		bean.Max = "--"
	}
	if isNeedReset(bean.Min) {
		bean.Min = "--"
	}
	if isNeedReset(bean.ProvinceScore) {
		bean.ProvinceScore = nil
	}

	query := func(c *mgo.Collection) error {
		log.Println(bean)
		return c.Insert(bean)
	}

	err := witchCollection(bean.getTableName(), query)
	if err != nil {
		log.Println("ERROR INSERT:", bean.getTableName(), bean, err)
		return "false"
	}

	return "true"
}
