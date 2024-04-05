package model

import (
	"fmt"
	"time"

	"github.com/takaodaze/beep/csv"
)

type Project string

const (
	gp  Project = "gp"
	trc Project = "trc"
)

type Data struct {
	Id      int
	Date    time.Time
	Project Project
}

func (d Data) ToRow() string {
	id := fmt.Sprintf("%d", d.Id)
	date := d.Date.Format("2006-01-02 15:04:05")
	return csv.ToRow([]string{id, date, string(d.Project)})
}

func NewGp() Data {
	d := Data{
		Id:      1,
		Date:    time.Now(),
		Project: gp,
	}
	return d
}
