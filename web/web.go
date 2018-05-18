package web

import (
	"net/http"
)

func Routes() {
	http.HandleFunc("/json", sendJSON)
	http.HandleFunc("/parse", parseParams)
	http.HandleFunc("/jbody", parseJsonBody)
	http.HandleFunc("/province", queryProvince)
	http.HandleFunc("/specialty", querySpecialty)
	http.HandleFunc("/spider/start", spiderRun)
}
