package okta

type RequesterFields struct {
	ID       string `json:"id,omitempty"`
	Prompt   string `json:"prompt,omitempty"`
	Required bool   `json:"required,omitempty"`
	Type     string `json:"type,omitempty"`
}

func NewRequesterFields() *RequesterFields {
	return &RequesterFields{}
}

func (a *RequesterFields) IsRequesterFields() bool {
	return true
}
