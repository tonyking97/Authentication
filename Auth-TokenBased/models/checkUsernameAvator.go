package models

/*
* @Author : Arputha Tony King P @Created at : Apr 29
* Definition : Schema for the username, which is used in checkusername to get first and last name for avatar.
 */

type CheckUsernameAvatar struct {
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
