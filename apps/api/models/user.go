package models

import "time"

type User struct {
    ID           uint      `gorm:"primaryKey" json:"id"`
    Email        string    `gorm:"unique;not null" json:"email"`
    PasswordHash string    `gorm:"not null" json:"-"`
    Role         string    `gorm:"not null;default:student" json:"role"`
	StudioID     uint      `json:"studio_id"`
    Studio       Studio    `gorm:"foreignKey:StudioID"`
    CreatedAt    time.Time `json:"created_at"`
    UpdatedAt    time.Time `json:"updated_at"`
}