package models

import (
	"encoding/json"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type SpecialtyScore struct {
	Id             bson.ObjectId `bson:"_id,omitempty"`
	SchoolId       json.Number   `bson:"schoolid" json:"schoolid"`
	SchooleName    string        `bson:"schoolname" json:"schoolname"`
	SpecialtyName  string        `bson:"specialtyname" json:"specialtyname"`
	LocalProvince  string        `bson:"localprovince" json:"localprovince"`
	Province       string        `bson:"province" json:"province"`
	StudentType    string        `bson:"studenttype" json:"studenttype"`
	Year           json.Number   `bson:"year" json:"year"`
	Batch          string        `bson:"batch" json:"batch"`
	Var            json.Number   `bson:"var" json:"var"`
	VarScore       json.Number   `bson:"var_score" json:"var_score"`
	Max            interface{}   `bson:"max" json:"max"`
	Min            interface{}   `bson:"min" json:"min"`
	Zyid           string        `bson:"zyid" json:"zyid"`
	Seesign        json.Number   `bson:"seesign" json:"seesign"`
	FirstRate      interface{}   `bson:"firstrate" json:"firstrate"`
	FirstRateClass interface{}   `bson:"firstrateclass" json:"firstrateclass"`
	URL            string        `bson:"url" json:"url"`
	Zytype         string        `bson:"zytype"`
	SchoolType     string        `bson:"schooltype"`
}

func (bean SpecialtyScore) getTableName() string { return "specialty_score" }

func (bean SpecialtyScore) Save() string {
	if isNeedReset(bean.Max) {
		bean.Max = "--"
	}
	if isNeedReset(bean.Min) {
		bean.Min = "--"
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
