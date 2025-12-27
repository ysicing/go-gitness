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

// ChecksService handles communication with check related methods
type ChecksService struct {
	client *Client
}

// Check represents a Gitness check
type Check struct {
	ID         *int64         `json:"id,omitempty"`
	Created    *Time          `json:"created,omitempty"`
	Updated    *Time          `json:"updated,omitempty"`
	RepoID     *int64         `json:"repo_id,omitempty"`
	CommitSHA  *string        `json:"commit_sha,omitempty"`
	Identifier *string        `json:"identifier,omitempty"`
	Status     *string        `json:"status,omitempty"`
	Started    *Time          `json:"started,omitempty"`
	Ended      *Time          `json:"ended,omitempty"`
	Link       *string        `json:"link,omitempty"`
	Summary    *string        `json:"summary,omitempty"`
	Payload    map[string]any `json:"payload,omitempty"`
	UID        *string        `json:"uid,omitempty"`
}

// CreateCheckOptions specifies options for creating a check
type CreateCheckOptions struct {
	Identifier *string        `json:"identifier,omitempty"`
	Status     *string        `json:"status,omitempty"`
	Started    *Time          `json:"started,omitempty"`
	Ended      *Time          `json:"ended,omitempty"`
	Link       *string        `json:"link,omitempty"`
	Summary    *string        `json:"summary,omitempty"`
	Payload    map[string]any `json:"payload,omitempty"`
}

// UpdateCheckOptions specifies options for updating a check
type UpdateCheckOptions struct {
	Status  *string        `json:"status,omitempty"`
	Started *Time          `json:"started,omitempty"`
	Ended   *Time          `json:"ended,omitempty"`
	Link    *string        `json:"link,omitempty"`
	Summary *string        `json:"summary,omitempty"`
	Payload map[string]any `json:"payload,omitempty"`
}

// ListChecksOptions specifies options for listing checks
type ListChecksOptions struct {
	Latest *bool `url:"latest,omitempty"`
}

// CreateCheck creates a check for a commit
func (s *ChecksService) CreateCheck(ctx context.Context, repoPath, commitSHA string, opt *CreateCheckOptions) (*Check, *Response, error) {
	path := fmt.Sprintf("repos/%s/commits/%s/checks", repoPath, commitSHA)
	var check Check
	resp, err := s.client.Post(ctx, path, opt, &check)
	if err != nil {
		return nil, resp, err
	}
	return &check, resp, nil
}

// UpdateCheck updates a check
func (s *ChecksService) UpdateCheck(ctx context.Context, repoPath, commitSHA, checkIdentifier string, opt *UpdateCheckOptions) (*Check, *Response, error) {
	path := fmt.Sprintf("repos/%s/commits/%s/checks/%s", repoPath, commitSHA, checkIdentifier)
	var check Check
	resp, err := s.client.Patch(ctx, path, opt, &check)
	if err != nil {
		return nil, resp, err
	}
	return &check, resp, nil
}

// ListChecks lists checks for a commit
func (s *ChecksService) ListChecks(ctx context.Context, repoPath, commitSHA string, opt *ListChecksOptions) ([]*Check, *Response, error) {
	path := fmt.Sprintf("repos/%s/commits/%s/checks", repoPath, commitSHA)
	req := s.client.client.R().SetContext(ctx)

	// Add specific query parameters
	if opt != nil && opt.Latest != nil {
		req.SetQueryParam("latest", fmt.Sprintf("%t", *opt.Latest))
	}

	var checks []*Check
	req.SetSuccessResult(&checks)

	resp, err := req.Get(path)
	if err != nil {
		return nil, &Response{Response: resp}, err
	}

	if err := s.client.checkResponse(resp); err != nil {
		return nil, &Response{Response: resp}, err
	}

	// Parse pagination headers
	response := &Response{Response: resp}
	s.client.parsePaginationHeaders(response)

	return checks, response, nil
}

// GetCheck retrieves a specific check
func (s *ChecksService) GetCheck(ctx context.Context, repoPath, commitSHA, checkIdentifier string) (*Check, *Response, error) {
	path := fmt.Sprintf("repos/%s/commits/%s/checks/%s", repoPath, commitSHA, checkIdentifier)
	var check Check
	resp, err := s.client.Get(ctx, path, &check)
	if err != nil {
		return nil, resp, err
	}
	return &check, resp, nil
}

// TemplatesService handles communication with template related methods
type TemplatesService struct {
	client *Client
}

// Template represents a Gitness template
type Template struct {
	Identifier  *string `json:"identifier,omitempty"`
	Description *string `json:"description,omitempty"`
	Data        *string `json:"data,omitempty"`
	Type        *string `json:"type,omitempty"`
	SpaceID     *int64  `json:"space_id,omitempty"`
	Created     *Time   `json:"created,omitempty"`
	Updated     *Time   `json:"updated,omitempty"`
}

// CreateTemplateOptions specifies options for creating a template
type CreateTemplateOptions struct {
	Identifier  *string `json:"identifier,omitempty"`
	Description *string `json:"description,omitempty"`
	Data        *string `json:"data,omitempty"`
	Type        *string `json:"type,omitempty"`
}

// UpdateTemplateOptions specifies options for updating a template
type UpdateTemplateOptions struct {
	Description *string `json:"description,omitempty"`
	Data        *string `json:"data,omitempty"`
}

// CreateTemplate creates a new template
func (s *TemplatesService) CreateTemplate(ctx context.Context, spaceRef string, opt *CreateTemplateOptions) (*Template, *Response, error) {
	path := fmt.Sprintf("spaces/%s/templates", spaceRef)
	var template Template
	resp, err := s.client.Post(ctx, path, opt, &template)
	if err != nil {
		return nil, resp, err
	}
	return &template, resp, nil
}

// ListTemplates lists templates in a space
func (s *TemplatesService) ListTemplates(ctx context.Context, spaceRef string, opt *ListOptions) ([]*Template, *Response, error) {
	path := fmt.Sprintf("spaces/%s/templates", spaceRef)
	var templates []*Template
	resp, err := s.client.performListRequest(ctx, path, opt, &templates)
	if err != nil {
		return nil, resp, err
	}
	return templates, resp, nil
}

// GetTemplate retrieves a specific template
func (s *TemplatesService) GetTemplate(ctx context.Context, spaceRef, templateIdentifier string) (*Template, *Response, error) {
	path := fmt.Sprintf("spaces/%s/templates/%s", spaceRef, templateIdentifier)
	var template Template
	resp, err := s.client.Get(ctx, path, &template)
	if err != nil {
		return nil, resp, err
	}
	return &template, resp, nil
}

// UpdateTemplate updates a template
func (s *TemplatesService) UpdateTemplate(ctx context.Context, spaceRef, templateIdentifier string, opt *UpdateTemplateOptions) (*Template, *Response, error) {
	path := fmt.Sprintf("spaces/%s/templates/%s", spaceRef, templateIdentifier)
	var template Template
	resp, err := s.client.Patch(ctx, path, opt, &template)
	if err != nil {
		return nil, resp, err
	}
	return &template, resp, nil
}

// DeleteTemplate deletes a template
func (s *TemplatesService) DeleteTemplate(ctx context.Context, spaceRef, templateIdentifier string) (*Response, error) {
	path := fmt.Sprintf("spaces/%s/templates/%s", spaceRef, templateIdentifier)
	resp, err := s.client.Delete(ctx, path, nil)
	return resp, err
}
