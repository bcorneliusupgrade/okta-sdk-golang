package okta

import (
	"context"
	"fmt"
	"time"
)

type RequestTypeResource resource

type RequestType struct {
	ID               string            `json:"id,omitempty"`
	Name             string            `json:"name,omitempty"`
	Description      string            `json:"description,omitempty"`
	Created          *time.Time        `json:"created,omitempty"`
	CreatedBy        string            `json:"createdBy,omitempty"`
	LastUpdated      *time.Time        `json:"lastUpdated,omitempty"`
	LastUpdatedBy    string            `json:"lastUpdatedBy,omitempty"`
	Status           string            `json:"status,omitempty"`
	OwnerID          string            `json:"ownerId,omitempty"`
	ResourceSettings *ResourceSettings `json:"resourceSettings,omitempty"`
	RequestSettings  *RequestSettings  `json:"requestSettings,omitempty"`
	ApprovalSettings *ApprovalSettings `json:"approvalSettings,omitempty"`
	LastUpdateSource string            `json:"lastUpdateSource,omitempty"`
	AccessDuration   interface{}       `json:"accessDuration,omitempty"`
	Links            interface{}       `json:"_links,omitempty"`
}

func NewRequestType() *RequestType {
	return &RequestType{}
}

func (a *RequestType) RequestType() bool {
	return true
}

// Creates a new Request Type in your Okta organization.
func (m *RequestTypeResource) CreateRequestType(ctx context.Context, body RequestType) (*RequestType, *Response, error) {
	url := "/governance/api/v1/request-types"

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var requestType *RequestType

	resp, err := rq.Do(ctx, req, &requestType)
	if err != nil {
		return nil, resp, err
	}

	return requestType, resp, nil
}

// fetches single Request Type from your Okta organization.
func (m *RequestTypeResource) GetRequestType(ctx context.Context, requestTypeId string) (*RequestType, *Response, error) {
	url := fmt.Sprintf("/governance/api/v1/request-types/%v", requestTypeId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var requestType *RequestType

	resp, err := rq.Do(ctx, req, &requestType)
	if err != nil {
		return nil, resp, err
	}

	return requestType, resp, nil
}

// fetches all Request Types from your Okta organization.
func (m *RequestTypeResource) GetRequestTypes(ctx context.Context) ([]*RequestType, *Response, error) {
	url := "/governance/api/v1/request-types"

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var requestType []*RequestType

	resp, err := rq.Do(ctx, req, &requestType)
	if err != nil {
		return nil, resp, err
	}

	return requestType, resp, nil
}

// Deletes a Request Type permanently.
func (m *RequestTypeResource) DeleteRequestType(ctx context.Context, requestTypeId string) (*Response, error) {
	url := fmt.Sprintf("/governance/api/v1/request-types/%v", requestTypeId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Publishes a Request Type.
func (m *RequestTypeResource) PublishRequestType(ctx context.Context, requestTypeId string) (*Response, error) {
	url := fmt.Sprintf("/governance/api/v1/request-types/%v", requestTypeId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Unpublishes a Request Type.
func (m *RequestTypeResource) UnpublishRequestType(ctx context.Context, requestTypeId string) (*Response, error) {
	url := fmt.Sprintf("/governance/api/v1/request-types/%v", requestTypeId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
