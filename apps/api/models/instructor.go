package models

import "time"

type Instructor struct {
    ID        int       `json:"id"`
    StudioID  uint      `json:"studio_id"`
    FirstName string    `json:"first_name"`
    LastName  string    `json:"last_name"`
    Email     string    `json:"email"`
    Phone     string    `json:"phone,omitempty"`
    Bio       string    `json:"bio,omitempty"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
