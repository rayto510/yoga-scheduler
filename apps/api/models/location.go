package models

import "time"

type Location struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    StudioID  uint      `json:"studio_id"`
    Name      string    `json:"name" binding:"required"`
    Address   string    `json:"address"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
