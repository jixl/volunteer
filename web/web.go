package web

import (
	"net/http"
)

func Routes() {
	http.HandleFunc("/json", sendJSON)
	http.HandleFunc("/parse", parseParams)
	http.HandleFunc("/province", queryProvince)
	http.HandleFunc("/specialty", querySpecialty)
}
