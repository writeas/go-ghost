package ghost

import (
	"time"
)

type (
	// Post represents a post on Ghost.
	Post struct {
		ID                 *string   `json:"id"`
		UUID               *string   `json:"uuid"`
		Title              *string   `json:"title"`
		Slug               *string   `json:"slug"`
		HTML               *string   `json:"html"`
		CommentID          *string   `json:"comment_id"`
		FeatureImage       *string   `json:"feature_image"`
		Featured           *bool     `json:"featured"`
		Page               *bool     `json:"page"`
		MetaTitle          *string   `json:"meta_title"`
		MetaDescription    *string   `json:"meta_description"`
		CreatedAt          time.Time `json:"created_at"`
		UpdatedAt          time.Time `json:"updated_at"`
		PublishedAt        time.Time `json:"published_at"`
		CustomExcerpt      *string   `json:"custom_excerpt"`
		OGImage            *string   `json:"og_image"`
		OGTitle            *string   `json:"og_title"`
		OGDescription      *string   `json:"og_description"`
		TwitterImage       *string   `json:"twitter_image"`
		TwitterTitle       *string   `json:"twitter_title"`
		TwitterDescription *string   `json:"twitter_description"`
		CustomTemplate     *string   `json:"custom_template"`
		PrimaryAuthor      *Author   `json:"primary_author"`
		PrimaryTag         *Tag      `json:"primary_tag"`
		URL                *string   `json:"url"`
		Excerpt            *string   `json:"excerpt"`
	}

	// PostParams are used for content publishing requests made to the Admin API.
	PostParams struct {
		Title     *string `json:"title"`
		HTML      *string `json:"html,omitempty"`
		Markdown  *string `json:"markdown,omitempty"`
		Mobiledoc *string `json:"mobiledoc,omitempty"`
		Status    *string `json:"status,omitempty"`
	}
)
