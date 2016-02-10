package model

type Stadium struct {
	Sid    int `gorm:"primary_key"`
	Name   string
	Events []Event
}
