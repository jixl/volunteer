package models

import (
	"gopkg.in/mgo.v2"
)

const URL = "127.0.0.1:27017" //mongodb连接字符串
// const URL = "test:test@192.168.0.240:27017"

var (
	mgoSession *mgo.Session
	tables     []interface{}
	HasSession bool

	DbCfg struct {
		Type, Host, Name, User, Passwd, Path, SSLMode string
	}

	dataBase = "test"
)

/**
 * 公共方法，获取session，如果存在则拷贝一份
 */
func getSession() *mgo.Session {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial(URL)
		if err != nil {
			panic(err) //直接终止程序运行
		}
		mgoSession.SetMode(mgo.Monotonic, true)
	}
	//最大连接池默认为4096
	return mgoSession.Clone()
}

//公共方法，获取collection对象
func witchCollection(collection string, s func(*mgo.Collection) error) error {
	session := getSession()
	defer session.Close()
	c := session.DB(dataBase).C(collection)
	return s(c)
}

// /**
//  * 获取一条记录通过objectid
//  */
// func GetPersonById(id string) *Person {
// 	objid := bson.ObjectIdHex(id)
// 	person := new(Person)
// 	query := func(c *mgo.Collection) error {
// 		return c.FindId(objid).One(&person)
// 	}
// 	witchCollection("person", query)
// 	return person
// }

// //获取所有的person数据
// func PagePerson() []Person {
// 	var persons []Person
// 	query := func(c *mgo.Collection) error {
// 		return c.Find(nil).All(&persons)
// 	}
// 	err := witchCollection("person", query)
// 	if err != nil {
// 		return persons
// 	}
// 	return persons
// }

// //更新person数据
// func UpdatePerson(query bson.M, change bson.M) string {
// 	exop := func(c *mgo.Collection) error {
// 		return c.Update(query, change)
// 	}
// 	err := witchCollection("person", exop)
// 	if err != nil {
// 		return "true"
// 	}
// 	return "false"
// }

// /**
//  * 执行查询，此方法可拆分做为公共方法
//  * [SearchPerson description]
//  * @param {[type]} collectionName string [description]
//  * @param {[type]} query          bson.M [description]
//  * @param {[type]} sort           bson.M [description]
//  * @param {[type]} fields         bson.M [description]
//  * @param {[type]} skip           int    [description]
//  * @param {[type]} limit          int)   (results      []interface{}, err error [description]
//  */
// func SearchPerson(collectionName string, query bson.M, sort string, fields bson.M, skip int, limit int) (results []interface{}, err error) {
// 	exop := func(c *mgo.Collection) error {
// 		return c.Find(query).Sort(sort).Select(fields).Skip(skip).Limit(limit).All(&results)
// 	}
// 	err = witchCollection(collectionName, exop)
// 	return
// }
