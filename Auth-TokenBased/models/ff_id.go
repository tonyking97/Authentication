package models

/*
* @Author : Arputha Tony King P @Created at : May 6'19
* Definition : Schema for the ff_id to create filter for DB query
 */

type Ff_id struct {
	Ff_id string `json:"ff_id" valid:"required~Empty id"`
}
