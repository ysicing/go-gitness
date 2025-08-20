// Copyright (c) 2025-2025 All rights reserved.
//
// The original source code is licensed under the Apache License 2.0.
//
// You may review the terms of both licenses in the LICENSE file.

package gitness

import (
	"context"
)

// ResourceService handles communication with resource related methods
type ResourceService struct {
	client *Client
}

// GitIgnoreTemplate represents a gitignore template
type GitIgnoreTemplate struct {
	Name    *string `json:"name,omitempty"`
	Content *string `json:"content,omitempty"`
}

// LicenseTemplate represents a license template
type LicenseTemplate struct {
	Key         *string `json:"key,omitempty"`
	Name        *string `json:"name,omitempty"`
	SPDXID      *string `json:"spdx_id,omitempty"`
	Description *string `json:"description,omitempty"`
	Content     *string `json:"content,omitempty"`
}

// ListGitIgnoreTemplates lists available gitignore templates
func (s *ResourceService) ListGitIgnoreTemplates(ctx context.Context) ([]*GitIgnoreTemplate, *Response, error) {
	var templates []*GitIgnoreTemplate
	resp, err := s.client.Get(ctx, "resources/gitignore", &templates)
	if err != nil {
		return nil, resp, err
	}
	return templates, resp, nil
}

// ListLicenseTemplates lists available license templates
func (s *ResourceService) ListLicenseTemplates(ctx context.Context) ([]*LicenseTemplate, *Response, error) {
	var templates []*LicenseTemplate
	resp, err := s.client.Get(ctx, "resources/license", &templates)
	if err != nil {
		return nil, resp, err
	}
	return templates, resp, nil
}
