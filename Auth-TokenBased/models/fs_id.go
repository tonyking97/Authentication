package models

type Fs_id struct {
	Fs_id string `json:"fs_id" valid:"required~Empty id"`
}
