package model

import "time"

type Source struct {
	ID        int64 `gorm: "primary_key"`
	Name      string
	URL       string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
