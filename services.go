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
