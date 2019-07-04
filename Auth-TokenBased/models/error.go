package models

import "net/http"

/*
* @Author : Arputha Tony King P @Created at : Apr 19
* Definition : Schema to create the Error message
 */

type ErrorMessage struct {
	Message string `json:"err"`
}

func CheckError(err error, w http.ResponseWriter) {
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
}
