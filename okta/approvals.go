package okta

type Approvals struct {
	ApproverType     string            `json:"approverType,omitempty"`
	Description      string            `json:"description,omitempty"`
	ApproverMemberOf []string          `json:"approverMemberOf,omitempty"`
	ApproverFields   []RequesterFields `json:"approverFields,omitempty"`
}

func NewApprovals() *Approvals {
	return &Approvals{}
}

func (a *Approvals) IsApprover() bool {
	return true
}
