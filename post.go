package ghost

import (
	"time"
)

type (
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

	PostParams struct {
		Title     *string `json:"title"`
		HTML      *string `json:"html,omitempty"`
		Markdown  *string `json:"markdown,omitempty"`
		Mobiledoc *string `json:"mobiledoc,omitempty"`
		Status    *string `json:"status,omitempty"`
	}

	Mobiledoc struct {
		Version  string          `json:"version"`
		Atoms    []interface{}   `json:"atoms"`
		Cards    [][]interface{} `json:"cards"`
		Markups  []interface{}   `json:"markups"`
		Sections [][]interface{} `json:"sections"`
	}

	Card struct {
		CardName *string `json:"cardName,omitempty"`
		HTML     *string `json:"html,omitempty"`
		Markdown *string `json:"markdown,omitempty"`
	}
)

// NewMarkdownMobiledoc wraps the given Markdown in a Mobiledoc struct
func NewMarkdownMobiledoc(md string) *Mobiledoc {
	return &Mobiledoc{
		Version: "0.3.1",
		Atoms:   []interface{}{},
		Cards: [][]interface{}{
			[]interface{}{
				"markdown",
				Card{
					CardName: str("markdown"),
					Markdown: str(md),
				},
			},
		},
		Markups: []interface{}{},
		Sections: [][]interface{}{
			[]interface{}{10, 0},
		},
	}
}

func str(s string) *string {
	return &s
}
