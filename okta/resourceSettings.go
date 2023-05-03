package okta

type ResourceSettings struct {
	Type            string            `json:"type,omitempty"`
	TargetResources []TargetResources `json:"targetResources,omitempty"`
}

func NewResourceSettings() *ResourceSettings {
	return &ResourceSettings{}
}

func (a *ResourceSettings) ResourceSettings() bool {
	return true
}
