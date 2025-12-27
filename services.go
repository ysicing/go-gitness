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

// GitspacesService handles communication with gitspace related methods
type GitspacesService struct {
	client *Client
}

// InfraProvidersService handles communication with infrastructure provider related methods
type InfraProvidersService struct {
	client *Client
}

// SecretsService handles communication with secret related methods
type SecretsService struct {
	client *Client
}

// WebhooksService handles communication with webhook related methods
type WebhooksService struct {
	client *Client
}

// Webhook represents a Gitness webhook
type Webhook struct {
	ID          *int64   `json:"id,omitempty"`
	Identifier  *string  `json:"identifier,omitempty"`
	Description *string  `json:"description,omitempty"`
	URL         *string  `json:"url,omitempty"`
	Secret      *string  `json:"secret,omitempty"`
	Triggers    []string `json:"triggers,omitempty"`
	Enabled     *bool    `json:"enabled,omitempty"`
	Insecure    *bool    `json:"insecure,omitempty"`
	Created     *Time    `json:"created,omitempty"`
	Updated     *Time    `json:"updated,omitempty"`
}

// Secret represents a Gitness secret
type Secret struct {
	ID          *int64  `json:"id,omitempty"`
	Identifier  *string `json:"identifier,omitempty"`
	Description *string `json:"description,omitempty"`
	Created     *Time   `json:"created,omitempty"`
	Updated     *Time   `json:"updated,omitempty"`
}

// CreateWebhookOptions specifies options for creating a webhook
type CreateWebhookOptions struct {
	Identifier  *string  `json:"identifier,omitempty"`
	Description *string  `json:"description,omitempty"`
	URL         *string  `json:"url,omitempty"`
	Secret      *string  `json:"secret,omitempty"`
	Triggers    []string `json:"triggers,omitempty"`
	Enabled     *bool    `json:"enabled,omitempty"`
	Insecure    *bool    `json:"insecure,omitempty"`
}

// CreateSecretOptions specifies options for creating a secret
type CreateSecretOptions struct {
	Identifier  *string `json:"identifier,omitempty"`
	Description *string `json:"description,omitempty"`
	Data        *string `json:"data,omitempty"`
}

// CreateWebhook creates a webhook for a repository
func (s *WebhooksService) CreateWebhook(ctx context.Context, repoPath string, opt *CreateWebhookOptions) (*Webhook, *Response, error) {
	path := fmt.Sprintf("repos/%s/webhooks", repoPath)
	var webhook Webhook
	resp, err := s.client.Post(ctx, path, opt, &webhook)
	if err != nil {
		return nil, resp, err
	}
	return &webhook, resp, nil
}

// ListWebhooks lists webhooks for a repository
func (s *WebhooksService) ListWebhooks(ctx context.Context, repoPath string, opt *ListOptions) ([]*Webhook, *Response, error) {
	path := fmt.Sprintf("repos/%s/webhooks", repoPath)
	var webhooks []*Webhook
	resp, err := s.client.performListRequest(ctx, path, opt, &webhooks)
	if err != nil {
		return nil, resp, err
	}
	return webhooks, resp, nil
}

// CreateSecret creates a secret for a repository
func (s *SecretsService) CreateSecret(ctx context.Context, repoPath string, opt *CreateSecretOptions) (*Secret, *Response, error) {
	path := fmt.Sprintf("repos/%s/secrets", repoPath)
	var secret Secret
	resp, err := s.client.Post(ctx, path, opt, &secret)
	if err != nil {
		return nil, resp, err
	}
	return &secret, resp, nil
}

// ListRepoSecrets lists secrets for a repository
func (s *SecretsService) ListRepoSecrets(ctx context.Context, repoPath string, opt *ListOptions) ([]*Secret, *Response, error) {
	path := fmt.Sprintf("repos/%s/secrets", repoPath)
	var secrets []*Secret
	resp, err := s.client.performListRequest(ctx, path, opt, &secrets)
	if err != nil {
		return nil, resp, err
	}
	return secrets, resp, nil
}

// CreateRepoSecret creates a secret for a repository
func (s *SecretsService) CreateRepoSecret(ctx context.Context, repoPath string, opt *CreateSecretOptions) (*Secret, *Response, error) {
	path := fmt.Sprintf("repos/%s/secrets", repoPath)
	var secret Secret
	resp, err := s.client.Post(ctx, path, opt, &secret)
	if err != nil {
		return nil, resp, err
	}
	return &secret, resp, nil
}

// ListSpaceSecrets lists secrets for a space
func (s *SecretsService) ListSpaceSecrets(ctx context.Context, spaceRef string, opt *ListOptions) ([]*Secret, *Response, error) {
	path := fmt.Sprintf("spaces/%s/secrets", spaceRef)
	var secrets []*Secret
	resp, err := s.client.performListRequest(ctx, path, opt, &secrets)
	if err != nil {
		return nil, resp, err
	}
	return secrets, resp, nil
}

// CreateSpaceSecret creates a secret for a space
func (s *SecretsService) CreateSpaceSecret(ctx context.Context, spaceRef string, opt *CreateSecretOptions) (*Secret, *Response, error) {
	path := fmt.Sprintf("spaces/%s/secrets", spaceRef)
	var secret Secret
	resp, err := s.client.Post(ctx, path, opt, &secret)
	if err != nil {
		return nil, resp, err
	}
	return &secret, resp, nil
}

// ListGlobalSecrets lists global secrets
func (s *SecretsService) ListGlobalSecrets(ctx context.Context, opt *ListOptions) ([]*Secret, *Response, error) {
	var secrets []*Secret
	resp, err := s.client.performListRequest(ctx, "secrets", opt, &secrets)
	if err != nil {
		return nil, resp, err
	}
	return secrets, resp, nil
}

// CreateGlobalSecret creates a global secret
func (s *SecretsService) CreateGlobalSecret(ctx context.Context, opt *CreateSecretOptions) (*Secret, *Response, error) {
	var secret Secret
	resp, err := s.client.Post(ctx, "secrets", opt, &secret)
	if err != nil {
		return nil, resp, err
	}
	return &secret, resp, nil
}

// GetSecret retrieves a specific secret
func (s *SecretsService) GetSecret(ctx context.Context, secretRef string) (*Secret, *Response, error) {
	path := fmt.Sprintf("secrets/%s", secretRef)
	var secret Secret
	resp, err := s.client.Get(ctx, path, &secret)
	if err != nil {
		return nil, resp, err
	}
	return &secret, resp, nil
}

// UpdateSecret updates a secret
func (s *SecretsService) UpdateSecret(ctx context.Context, secretRef string, opt *CreateSecretOptions) (*Secret, *Response, error) {
	path := fmt.Sprintf("secrets/%s", secretRef)
	var secret Secret
	resp, err := s.client.Patch(ctx, path, opt, &secret)
	if err != nil {
		return nil, resp, err
	}
	return &secret, resp, nil
}

// DeleteSecret deletes a secret
func (s *SecretsService) DeleteSecret(ctx context.Context, secretRef string) (*Response, error) {
	path := fmt.Sprintf("secrets/%s", secretRef)
	resp, err := s.client.Delete(ctx, path, nil)
	return resp, err
}

// Gitspace represents a Gitness gitspace
type Gitspace struct {
	ID                *int64         `json:"id,omitempty"`
	Identifier        *string        `json:"identifier,omitempty"`
	Name              *string        `json:"name,omitempty"`
	Description       *string        `json:"description,omitempty"`
	SpaceID           *int64         `json:"space_id,omitempty"`
	SpacePath         *string        `json:"space_path,omitempty"`
	IDE               *GitspaceIDE   `json:"ide,omitempty"`
	InfraProviderType *string        `json:"infra_provider_type,omitempty"`
	ResourceType      *string        `json:"resource_type,omitempty"`
	UserUID           *string        `json:"user_uid,omitempty"`
	UserDisplayName   *string        `json:"user_display_name,omitempty"`
	State             *GitspaceState `json:"state,omitempty"`
	URL               *string        `json:"url,omitempty"`
	Created           *Time          `json:"created,omitempty"`
	Updated           *Time          `json:"updated,omitempty"`
	Accessed          *Time          `json:"accessed,omitempty"`
	TotalTimeUsed     *int64         `json:"total_time_used,omitempty"`
}

// GitspaceIDE represents IDE configuration for a gitspace
type GitspaceIDE string

const (
	GitspaceIDEVSCode         GitspaceIDE = "vscode"
	GitspaceIDEVSCodeWeb      GitspaceIDE = "vscode-web"
	GitspaceIDEJetBrainsFleet GitspaceIDE = "jetbrains-fleet"
)

// GitspaceState represents the state of a gitspace
type GitspaceState string

const (
	GitspaceStateUnspecified GitspaceState = "unspecified"
	GitspaceStateRunning     GitspaceState = "running"
	GitspaceStateStopped     GitspaceState = "stopped"
	GitspaceStateError       GitspaceState = "error"
	GitspaceStateUnknown     GitspaceState = "unknown"
)

// GitspaceAction represents an action to perform on a gitspace
type GitspaceAction string

const (
	GitspaceActionStart GitspaceAction = "start"
	GitspaceActionStop  GitspaceAction = "stop"
)

// ListGitspacesOptions specifies the optional parameters for listing gitspaces
type ListGitspacesOptions struct {
	ListOptions
	SpaceRef *string `url:"space_ref,omitempty"`
}

// ListGitspaces lists gitspaces with optional filtering
func (s *GitspacesService) ListGitspaces(ctx context.Context, opt *ListGitspacesOptions) ([]*Gitspace, *Response, error) {
	req := s.client.client.R().SetContext(ctx)

	if opt != nil {
		buildQueryParams(req, &opt.ListOptions)
		if opt.SpaceRef != nil {
			req.SetQueryParam("space_ref", *opt.SpaceRef)
		}
	}

	var gitspaces []*Gitspace
	req.SetSuccessResult(&gitspaces)

	resp, err := req.Get("gitspaces")
	if err != nil {
		return nil, &Response{Response: resp}, err
	}

	if err := s.client.checkResponse(resp); err != nil {
		return nil, &Response{Response: resp}, err
	}

	response := &Response{Response: resp}
	s.client.parsePaginationHeaders(response)

	return gitspaces, response, nil
}

// CreateGitspaceRequest represents a request to create a new gitspace
type CreateGitspaceRequest struct {
	Identifier        *string     `json:"identifier,omitempty"`
	Name              *string     `json:"name,omitempty"`
	Description       *string     `json:"description,omitempty"`
	SpaceRef          *string     `json:"space_ref,omitempty"`
	IDE               GitspaceIDE `json:"ide,omitempty"`
	InfraProviderType *string     `json:"infra_provider_type,omitempty"`
	ResourceType      *string     `json:"resource_type,omitempty"`
}

// CreateGitspace creates a new gitspace
func (s *GitspacesService) CreateGitspace(ctx context.Context, gitspace *CreateGitspaceRequest) (*Gitspace, *Response, error) {
	var newGitspace Gitspace
	resp, err := s.client.Post(ctx, "gitspaces", gitspace, &newGitspace)
	if err != nil {
		return nil, resp, err
	}
	return &newGitspace, resp, nil
}

// FindGitspace retrieves a specific gitspace by identifier
func (s *GitspacesService) FindGitspace(ctx context.Context, identifier string) (*Gitspace, *Response, error) {
	path := fmt.Sprintf("gitspaces/%s", identifier)
	var gitspace Gitspace
	resp, err := s.client.Get(ctx, path, &gitspace)
	if err != nil {
		return nil, resp, err
	}
	return &gitspace, resp, nil
}

// DeleteGitspace deletes a gitspace by identifier
func (s *GitspacesService) DeleteGitspace(ctx context.Context, identifier string) (*Response, error) {
	path := fmt.Sprintf("gitspaces/%s", identifier)
	resp, err := s.client.Delete(ctx, path, nil)
	return resp, err
}

// GitspaceActionRequest represents a request to perform an action on a gitspace
type GitspaceActionRequest struct {
	Action GitspaceAction `json:"action,omitempty"`
}

// ActionOnGitspace performs an action on a gitspace (start/stop)
func (s *GitspacesService) ActionOnGitspace(ctx context.Context, identifier string, action GitspaceAction) (*Gitspace, *Response, error) {
	path := fmt.Sprintf("gitspaces/%s/actions", identifier)
	req := &GitspaceActionRequest{Action: action}

	var gitspace Gitspace
	resp, err := s.client.Post(ctx, path, req, &gitspace)
	if err != nil {
		return nil, resp, err
	}
	return &gitspace, resp, nil
}

// GitspaceEvent represents an event in gitspace lifecycle
type GitspaceEvent struct {
	ID        *int64  `json:"id,omitempty"`
	Type      *string `json:"type,omitempty"`
	Message   *string `json:"message,omitempty"`
	Created   *Time   `json:"created,omitempty"`
	Timestamp *Time   `json:"timestamp,omitempty"`
}

// ListGitspaceEventsOptions specifies the optional parameters for listing gitspace events
type ListGitspaceEventsOptions struct {
	ListOptions
}

// ListGitspaceEvents lists events for a specific gitspace
func (s *GitspacesService) ListGitspaceEvents(ctx context.Context, identifier string, opt *ListGitspaceEventsOptions) ([]*GitspaceEvent, *Response, error) {
	path := fmt.Sprintf("gitspaces/%s/events", identifier)
	req := s.client.client.R().SetContext(ctx)

	if opt != nil {
		buildQueryParams(req, &opt.ListOptions)
	}

	var events []*GitspaceEvent
	req.SetSuccessResult(&events)

	resp, err := req.Get(path)
	if err != nil {
		return nil, &Response{Response: resp}, err
	}

	if err := s.client.checkResponse(resp); err != nil {
		return nil, &Response{Response: resp}, err
	}

	response := &Response{Response: resp}
	s.client.parsePaginationHeaders(response)

	return events, response, nil
}

// InfraProvider represents an infrastructure provider
type InfraProvider struct {
	Identifier  *string                `json:"identifier,omitempty"`
	Name        *string                `json:"name,omitempty"`
	Description *string                `json:"description,omitempty"`
	Type        *InfraProviderType     `json:"type,omitempty"`
	SpaceID     *int64                 `json:"space_id,omitempty"`
	SpacePath   *string                `json:"space_path,omitempty"`
	Metadata    *InfraProviderMetadata `json:"metadata,omitempty"`
	Templates   []*InfraTemplate       `json:"templates,omitempty"`
	Created     *Time                  `json:"created,omitempty"`
	Updated     *Time                  `json:"updated,omitempty"`
}

// InfraProviderType represents the type of infrastructure provider
type InfraProviderType string

const (
	InfraProviderTypeDocker     InfraProviderType = "docker"
	InfraProviderTypeKubernetes InfraProviderType = "kubernetes"
	InfraProviderTypeAWS        InfraProviderType = "aws"
	InfraProviderTypeGCP        InfraProviderType = "gcp"
	InfraProviderTypeAzure      InfraProviderType = "azure"
)

// InfraProviderMetadata represents metadata for an infrastructure provider
type InfraProviderMetadata struct {
	Region       *string           `json:"region,omitempty"`
	Zone         *string           `json:"zone,omitempty"`
	Host         *string           `json:"host,omitempty"`
	Port         *int              `json:"port,omitempty"`
	Namespace    *string           `json:"namespace,omitempty"`
	StorageClass *string           `json:"storage_class,omitempty"`
	Network      *string           `json:"network,omitempty"`
	Subnet       *string           `json:"subnet,omitempty"`
	Credentials  map[string]string `json:"credentials,omitempty"`
	Properties   map[string]any    `json:"properties,omitempty"`
}

// InfraTemplate represents a resource template for an infrastructure provider
type InfraTemplate struct {
	Identifier  *string        `json:"identifier,omitempty"`
	Name        *string        `json:"name,omitempty"`
	Description *string        `json:"description,omitempty"`
	CPU         *string        `json:"cpu,omitempty"`
	Memory      *string        `json:"memory,omitempty"`
	Disk        *string        `json:"disk,omitempty"`
	Properties  map[string]any `json:"properties,omitempty"`
}

// CreateInfraProviderRequest represents a request to create a new infrastructure provider
type CreateInfraProviderRequest struct {
	Identifier  *string                `json:"identifier,omitempty"`
	Name        *string                `json:"name,omitempty"`
	Description *string                `json:"description,omitempty"`
	Type        InfraProviderType      `json:"type,omitempty"`
	SpaceRef    *string                `json:"space_ref,omitempty"`
	Metadata    *InfraProviderMetadata `json:"metadata,omitempty"`
	Templates   []*InfraTemplate       `json:"templates,omitempty"`
}

// CreateInfraProvider creates a new infrastructure provider
func (s *InfraProvidersService) CreateInfraProvider(ctx context.Context, spaceRef string, provider *CreateInfraProviderRequest) (*InfraProvider, *Response, error) {
	path := fmt.Sprintf("spaces/%s/infra-providers", spaceRef)
	var infraProvider InfraProvider
	resp, err := s.client.Post(ctx, path, provider, &infraProvider)
	if err != nil {
		return nil, resp, err
	}
	return &infraProvider, resp, nil
}

// GetInfraProvider retrieves a specific infrastructure provider by identifier
func (s *InfraProvidersService) GetInfraProvider(ctx context.Context, spaceRef, identifier string) (*InfraProvider, *Response, error) {
	path := fmt.Sprintf("spaces/%s/infra-providers/%s", spaceRef, identifier)
	var infraProvider InfraProvider
	resp, err := s.client.Get(ctx, path, &infraProvider)
	if err != nil {
		return nil, resp, err
	}
	return &infraProvider, resp, nil
}
