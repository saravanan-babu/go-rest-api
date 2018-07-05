package api

import (
	"encoding/json"
	"net/http"

	"../apimodel"
	"../common"
)

type ApiHandler func(w http.ResponseWriter, r *http.Request)

//Welcome
func Welcome(w http.ResponseWriter, r *http.Request) {

	data := "Welcome to new Golang API!"
	common.ToJsonHttpResponse(w, http.StatusOK, data)

}

//GetUsers
func GetUsers(w http.ResponseWriter, r *http.Request) {

	users := [2]apimodel.User{apimodel.User{Id: 10, Name: "sara"}, apimodel.User{Id: 20, Name: "flintoff"}}

	common.ToJsonHttpResponse(w, http.StatusOK, users)
}

//UpdateUsers
func UpdateUsers(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	var t apimodel.User
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	common.ToJsonHttpResponse(w, http.StatusOK, "Updated Successfully!")

}
