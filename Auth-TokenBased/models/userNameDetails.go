package models

/*
* @Author : Arputha Tony King P @Created at : Jul 2
* Definition : Schema for the get user name details
 */

type UserNameDetails struct {
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}
