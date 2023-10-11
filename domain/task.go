package domain

import (
	"time"
)

type Task struct {
	ID          int
	Name        string
	Description string
	Due         time.Time
	Done        bool
	Tags        []*Tag
	CreatedAt   time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
}
