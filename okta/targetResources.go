package okta

type TargetResources struct {
	ResourceID string `json:"resourceId"`
}

func NewTargetResources() *TargetResources {
	return &TargetResources{}
}

func (a *TargetResources) TargetResources() bool {
	return true
}
