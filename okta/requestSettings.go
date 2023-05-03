package okta

type RequestSettings struct {
	Type            string            `json:"type,omitempty"`
	RequesterFields []RequesterFields `json:"requesterFields,omitempty"`
}

func NewRequestSettings() *RequestSettings {
	return &RequestSettings{}
}

func (a *RequestSettings) RequestSettings() bool {
	return true
}
