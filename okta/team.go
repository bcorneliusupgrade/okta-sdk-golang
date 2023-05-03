package okta

import (
	"context"
	"time"
)

type TeamResource resource

type Team struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Created     time.Time `json:"created"`
	LastUpdated time.Time `json:"lastUpdated"`
}

type Teams struct {
	Data []Team `json:"data"`
}

func NewTeam() *Team {
	return &Team{}
}

func (a *Team) Team() bool {
	return true
}

// fetches RequestType from your Okta organization.
func (m *TeamResource) GetTeams(ctx context.Context) (*Teams, *Response, error) {
	url := "/governance/api/v1/teams"

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var teams *Teams

	resp, err := rq.Do(ctx, req, &teams)
	if err != nil {
		return nil, resp, err
	}

	return teams, resp, nil
}
