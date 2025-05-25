package models

import "time"

type Class struct {
    ID           int        `json:"id" gorm:"primaryKey"`
    StudioID     uint       `json:"studio_id"`
    InstructorID int        `json:"instructor_id"`
    LocationID   *int       `json:"location_id,omitempty"` // pointer allows NULL
    Name         string     `json:"name"`
    Description  string     `json:"description,omitempty"`
    Capacity     int        `json:"capacity"`
    StartTime    time.Time  `json:"start_time"`
    EndTime      time.Time  `json:"end_time"`
    CreatedAt    time.Time  `json:"created_at"`
    UpdatedAt    time.Time  `json:"updated_at"`
}
