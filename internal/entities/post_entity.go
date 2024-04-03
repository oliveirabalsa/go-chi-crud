package entities

import (
	"encoding/json"

	"github.com/google/uuid"
)

type Post struct {
	Id          uuid.UUID `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"title"`
	Description string    `json:"description" gorm:"description"`
}

func (p *Post) ToJson() []byte {
	data, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return data
}
