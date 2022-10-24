package model

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type List struct {
	gorm.Model
	ID       int
	ListName string `gorm:"uniqueIndex"`
	Items    []ListEntry
}

func NewList(title string) List {
	return List{
		ListName: title,
	}
}

func (l *List) String() string {
	to_join := make([]string, 0)
	to_join = append(to_join, l.ListName)
	for _, v := range l.Items {
		to_join = append(to_join, fmt.Sprint(&v))
	}
	return strings.Join(to_join, "\n")
}
