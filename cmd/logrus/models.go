package main

import "github.com/google/uuid"

type User struct {
	ID   uuid.UUID `json:"id,omitempty"`
	Name string    `json:"name,omitempty"`
}

type Article struct {
	ID    uuid.UUID `json:"id,omitempty"`
	Title string    `json:"title,omitempty"`
	Text  string    `json:"text,omitempty"`

	UserID uuid.UUID `json:"user_id,omitempty"`
}
