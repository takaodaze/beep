package model

import "time"

type Project int

const (
	gp Project = iota
	trc
)

type Data struct {
	Id      int
	Date    time.Time
	Project Project
}

func NewGp() Data {
	d := Data{
		Id:      1,
		Date:    time.Now(),
		Project: gp,
	}
	return d
}
