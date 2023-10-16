package models

import (
	"time"
)

type Product struct {
	Id         int
	Name       string
	ExpireDate time.Time
	UpdatedAt time.Time
	CreatedAt time.Time
	Price int
	Status bool
}
