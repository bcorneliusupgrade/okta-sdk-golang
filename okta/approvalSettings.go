package okta

type ApprovalSettings struct {
	Type      string      `json:"type,omitempty"`
	Approvals []Approvals `json:"approvals,omitempty"`
}

func NewApprovalSettings() *ApprovalSettings {
	return &ApprovalSettings{}
}

func (a *ApprovalSettings) IsApprovalSettings() bool {
	return true
}
