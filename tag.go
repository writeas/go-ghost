package ghost

import (
	"time"
)

// Tag represents a tag on Ghost.
type Tag struct {
	ID              *string   `json:"id"`
	Name            *string   `json:"name"`
	Slug            *string   `json:"slug"`
	Description     *string   `json:"description"`
	FeaturedImage   *string   `json:"featured_image"`
	Visibility      *string   `json:"visibility"`
	MetaTitle       *string   `json:"meta_title"`
	MetaDescription *string   `json:"meta_description"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	Parent          *Tag      `json:"parent"`
}
