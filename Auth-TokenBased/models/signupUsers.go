package models

/*
* @Author : Arputha Tony King P @Created at : Apr 19
* Definition : Schema to create the signUp user details and insert in the MongoDB collection
 */

type SignUpDetails struct {
	Id          string `json:"id" bson:"_id,omitempty"`
	Email       string `json:"email" valid:"required~EmailID is required,email~Valid Email ID required"`
	FirstName   string `json:"firstname" valid:"alpha~Firstname should contain only Alphabets,required~Firstname is required"`
	LastName    string `json:"lastname" valid:"alpha~Lastname should contain only Alphabets,required~Lastname is required"`
	Username    string `json:"username" valid:"length(4|15)~Username should contain at least 4 characters and maximum of 15,alphanum~Username should contain only Alphabets and Numbers,required~Username is required"`
	Password    string `json:"password" valid:"length(6|15)~Password should contain at least 6 characters and maximum of 15,required~Password is required,matches(^.*[a-z].*$)~Password should contain One Lowercase,matches(^.*[A-Z].*$)~Password should contain One Uppercase,matches(^.*[0-9].*$)~Password should contain One Number,matches(^.*[!@#$^&*_+=\\-<>?.].*$)~Password should contain One Symbol. Allowed special characters are ! @ # $ ^ & * _ - + = < > ? ."`
	CreatedTime int64  `json:"created_time"`
	UpdatedTime int64  `json:"updated_time"`
}
