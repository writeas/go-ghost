package ghost

type Error struct {
	ID      *string `json:"id"`
	Message *string `json:"message"`
	Context *string `json:"context"`
	Type    *string `json:"type"`
	/*
		TODO: determine types of these properties
		Details  *string `json:"details"`
		Property *string `json:"property"`
		Help     *string `json:"help"`
		Code     *string `json:"code"`
	*/
}

func (e Error) Error() string {
	return *e.Message + ": " + *e.Context
}
