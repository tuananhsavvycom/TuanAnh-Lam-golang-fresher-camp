package common

import "time"

type SQLModel struct {
	Status    int        `json:"status" gorm:"status"`
	CreatedAt *time.Time `json:"created_at" gorm:"created_at"`
	UpdateAt  *time.Time `json:"updated_at" gorm:"updated_at"`
}
