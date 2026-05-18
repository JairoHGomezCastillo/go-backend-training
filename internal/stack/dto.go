package stack

import "time"

type Stack struct {
	Id              int
	Name            string
	Description     string
	Purpose         Purpose
	Segment         Segment
	ApplicationName string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	CreatedBy       string
	UpdatedBy       string
}
type Segment struct {
	Id           int
	Name         string
	IsProductive int
}
type Purpose struct {
	Id           int
	Name         string
	Description  string
	IsProductive int
}
