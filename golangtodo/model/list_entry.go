package model

import (
	"fmt"

	"gorm.io/gorm"
)

type Status int

const (
	ACTIVE Status = iota
	COMPLETE
)

type ListEntry struct {
	gorm.Model
	ID         int
	Content    string
	ItemStatus Status
	ListID     int
}

func NewListEntry(content string) ListEntry {
	return ListEntry{
		Content:    content,
		ItemStatus: ACTIVE,
	}
}

func (l *ListEntry) String() string {
	var check string
	if l.ItemStatus == COMPLETE {
		check = "[x]"
	} else {
		check = "[ ]"
	}
	return fmt.Sprintf("%s: %s", check, l.Content)
}
