package domain

import "time"

type Project struct {
	Name        string
	Description string
	Tasks       []*Task
	Tags        []*Tag
	CreatedAt   time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
}
