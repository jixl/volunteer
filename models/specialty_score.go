package models

import (
	"encoding/json"
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

func (bean *SpecialtyScore) beforeInsert() {
	if isNeedReset(bean.Max) {
		bean.Max = "--"
	}
	if isNeedReset(bean.Min) {
		bean.Min = "--"
	}
}

func FindSpecialty(opts *SearchOption) []SpecialtyScore {
	ds := NewSessionStore()
	defer ds.Close()
	table := SpecialtyScore{}.getTableName()
	coll := ds.C(table)
	query := coll.Find(opts.Choice)

	skip := getSkip(opts.Page, opts.PageSize)
	if skip > 0 {
		query.Skip(skip)
	}
	query.Limit(getPageSize(opts.PageSize))

	var list = []SpecialtyScore{}
	err := query.All(&list)
	if err != nil {
		log.Println("Find Error ", query, table)
	}
	return list
}
