package models

/*
* @Author : Arputha Tony King P @Created at : Apr 19
* Definition : Schema for the users, which is used in log in of the user.
 */

type User struct {
	Username     string `json:"username" valid:"length(4|15)~Username should contain at least 4 characters and maximum of 15,alphanum~Username should contain only Alphabets and Numbers,required~Username is required"`
	Password     string `json:"password" valid:"length(6|15)~Password should contain at least 6 characters and maximum of 15,required~Password is required,matches(^.*[a-z].*$)~Password should contain One Lowercase,matches(^.*[A-Z].*$)~Password should contain One Uppercase,matches(^.*[0-9].*$)~Password should contain One Number,matches(^.*[!@#$^&*_+=\\-<>?.].*$)~Password should contain One Symbol. Allowed special characters are ! @ # $ ^ & * _ - + = < > ? ."`
	Localization string
	Ip_details   IpDetails
}
