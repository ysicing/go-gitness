// Copyright (c) 2025-2025 All rights reserved.
//
// The original source code is licensed under the Apache License 2.0.
//
// You may review the terms of both licenses in the LICENSE file.

package gitness

import (
	"context"
	"fmt"
)

// ConnectorsService handles communication with connector related methods
type ConnectorsService struct {
	client *Client
}

// ConnectorType represents the type of connector
type ConnectorType string

// Connector types
const (
	ConnectorTypeGithub ConnectorType = "github"
)

// ConnectorStatus represents the status of connector
type ConnectorStatus string

// Connector statuses
const (
	ConnectorStatusOK      ConnectorStatus = "ok"
	ConnectorStatusError   ConnectorStatus = "error"
	ConnectorStatusPending ConnectorStatus = "pending"
)

// ConnectorAuthType represents the type of connector authentication
type ConnectorAuthType string

// Connector auth types
const (
	ConnectorAuthTypeBearer ConnectorAuthType = "bearer"
)

// ConnectorAuth represents connector authentication credentials
type ConnectorAuth struct {
	AuthType ConnectorAuthType `json:"auth_type,omitempty"`
	Token    *string           `json:"token,omitempty"`
}

// GithubConnectorData represents github connector specific data
type GithubConnectorData struct {
	APIURL   *string        `json:"api_url,omitempty"`
	Insecure *bool          `json:"insecure,omitempty"`
	Auth     *ConnectorAuth `json:"auth,omitempty"`
}

// Connector represents a Gitness connector based on TypesConnector schema
type Connector struct {
	Created          *int64               `json:"created,omitempty"`
	CreatedBy        *int64               `json:"created_by,omitempty"`
	Description      *string              `json:"description,omitempty"`
	Github           *GithubConnectorData `json:"github,omitempty"`
	Identifier       *string              `json:"identifier,omitempty"`
	LastTestAttempt  *int64               `json:"last_test_attempt,omitempty"`
	LastTestErrorMsg *string              `json:"last_test_error_msg,omitempty"`
	LastTestStatus   *ConnectorStatus     `json:"last_test_status,omitempty"`
	SpaceID          *int64               `json:"space_id,omitempty"`
	Type             *ConnectorType       `json:"type,omitempty"`
	Updated          *int64               `json:"updated,omitempty"`
}

// CreateConnectorOptions specifies options for creating a connector based on OpenapiCreateConnectorRequest schema
type CreateConnectorOptions struct {
	Description *string              `json:"description,omitempty"`
	Github      *GithubConnectorData `json:"github,omitempty"`
	Identifier  *string              `json:"identifier,omitempty"`
	SpaceRef    *string              `json:"space_ref,omitempty"`
	Type        *ConnectorType       `json:"type,omitempty"`
}

// UpdateConnectorOptions specifies options for updating a connector
type UpdateConnectorOptions struct {
	Description *string              `json:"description,omitempty"`
	Github      *GithubConnectorData `json:"github,omitempty"`
}

// ListConnectors lists all connectors
func (s *ConnectorsService) ListConnectors(ctx context.Context, opt *ListOptions) ([]*Connector, *Response, error) {
	var connectors []*Connector
	resp, err := s.client.performListRequest(ctx, "connectors", opt, &connectors)
	if err != nil {
		return nil, resp, err
	}
	return connectors, resp, nil
}

// GetConnector retrieves a specific connector by identifier
func (s *ConnectorsService) GetConnector(ctx context.Context, connectorRef string) (*Connector, *Response, error) {
	path := fmt.Sprintf("connectors/%s", connectorRef)
	var connector Connector
	resp, err := s.client.Get(ctx, path, &connector)
	if err != nil {
		return nil, resp, err
	}
	return &connector, resp, nil
}

// CreateConnector creates a new connector
func (s *ConnectorsService) CreateConnector(ctx context.Context, opt *CreateConnectorOptions) (*Connector, *Response, error) {
	var connector Connector
	resp, err := s.client.Post(ctx, "connectors", opt, &connector)
	if err != nil {
		return nil, resp, err
	}
	return &connector, resp, nil
}

// UpdateConnector updates an existing connector
func (s *ConnectorsService) UpdateConnector(ctx context.Context, connectorRef string, opt *UpdateConnectorOptions) (*Connector, *Response, error) {
	path := fmt.Sprintf("connectors/%s", connectorRef)
	var connector Connector
	resp, err := s.client.Patch(ctx, path, opt, &connector)
	if err != nil {
		return nil, resp, err
	}
	return &connector, resp, nil
}

// DeleteConnector deletes a connector
func (s *ConnectorsService) DeleteConnector(ctx context.Context, connectorRef string) (*Response, error) {
	path := fmt.Sprintf("connectors/%s", connectorRef)
	resp, err := s.client.Delete(ctx, path, nil)
	return resp, err
}
