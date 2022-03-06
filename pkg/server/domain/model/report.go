package model

import (
	"gorm.io/gorm"
)

type Report struct {
	gorm.Model
	Body    string `json:"body"`
	Weather string `json:"weather"`
}

type Reports struct {
	Reports []*Report `json:"reports"`
}

type CountReports struct {
	Count int `json:"count"`
}
