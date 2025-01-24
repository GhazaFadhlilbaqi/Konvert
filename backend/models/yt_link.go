package models

import (
	"github.com/google/uuid"
)

type YtVid struct {
	Id       uuid.UUID `json:"id"`
	Url      string    `jso:"url"`
	Title    string    `json:"title"`
	Duration string    `json:"duration"`
	Channel  string    `json:"Channel"`
}
