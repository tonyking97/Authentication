package models

/*
* @Author : Arputha Tony King P @Created at : Jul 2'19
* Definition : Schema for the ff_id to get user details
 */

type Ff_id_user struct {
	Ff_id string `json:"ff_id" bson:"_id" valid:"required~Empty id"`
}
