package models

/*
* @Author : Arputha Tony King P @Created at : Apr 29
* Definition : Schema for the checkusername, which is used in checkusername to return response.
 */

type CheckUsername struct {
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}
