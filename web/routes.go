package web

import (
	"encoding/json"
	"fmt"
	"strconv"
	"github.com/jixl/volunteer/models"
	"github.com/jixl/volunteer/spider/scores"
	"net/http"
	"strings"
)

func sendJSON(rw http.ResponseWriter, r *http.Request) {
	u := struct {
		Name string
		Age  int
	}{"hah哈哈", 10}

	// rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	json.NewEncoder(rw).Encode(&u)
}

func parseParams(rw http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(rw, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

type test_struct struct {
	Test string
}
func parseJsonBody(rw http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var t test_struct
	err := decoder.Decode(&t)

	if err != nil {
		panic(err)
	}

	fmt.Println(t.Test)

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(t)
}

func queryProvince(rw http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	data := models.FindProvince(&models.SearchOption{Choice: r.Form})
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(data)
	// fmt.Println(data)
	// fmt.Fprintf(rw, "Hello astaxie!")
}

func querySpecialty(rw http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	data := models.FindSpecialty(&models.SearchOption{Choice: r.Form})
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(data)
}

func spiderRun(rw http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	spiderType := r.FormValue("spider")
	data := map[string]string{"spider": spiderType}
	page, err := strconv.Atoi(r.FormValue("page"))
	if err != nil {
		page = 1
	}

	year := r.FormValue("year")
	if spiderType == "province" {
		go scores.Province(year, page)
	} else if spiderType == "specialty" {
		cate := r.FormValue("category")
		go scores.Specialty(cate, year, page)
	} else {
		data["error"] = "not has " + spiderType + " spider"
	}
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(data)
}
