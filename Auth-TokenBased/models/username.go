package models

/*
* @Author : Arputha Tony King P @Created at : Apr 19
* Definition : Schema for the username, which is used in checkusername to validate.
 */

type Username struct {
	Username string `json:"username" valid:"length(4|15)~Username should contain at least 4 characters and maximum of 15,alphanum~Username should contain only Alphabets and Numbers,required~Username is required"`
}
