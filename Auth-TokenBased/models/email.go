package models

/*
* @Author : Arputha Tony King P @Created at : May 3'19
* Definition : Schema for the email, which is used in checkemailavailability to validate.
 */

type Email struct {
	Email string `json:"email" valid:"required~EmailID is required,email~Valid Email ID required"`
}
