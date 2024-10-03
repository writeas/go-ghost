package ghost

// Author represents an author on Ghost.
type Author struct {
	ID              *string `json:"id"`
	Name            *string `json:"name"`
	Slug            *string `json:"slug"`
	ProfileImage    *string `json:"profile_image"`
	CoverImage      *string `json:"cover_image"`
	Bio             *string `json:"bio"`
	Website         *string `json:"website"`
	Location        *string `json:"location"`
	Facebook        *string `json:"facebook"`
	Twitter         *string `json:"twitter"`
	MetaTitle       *string `json:"meta_title"`
	MetaDescription *string `json:"meta_description"`
	URL             *string `json:"url"`
}
