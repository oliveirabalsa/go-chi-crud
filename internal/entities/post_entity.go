package entities

import (
	"encoding/json"

	"github.com/google/uuid"
)

type Post struct {
	Id          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}

func (p *Post) ToJson() []byte {
	data, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return data
}
