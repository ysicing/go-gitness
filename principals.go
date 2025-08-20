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

// PrincipalsService handles communication with principals related methods
type PrincipalsService struct {
	client *Client
}

// Principal represents a Gitness principal (user or service account)
type Principal struct {
	ID          *int64  `json:"id,omitempty"`
	Type        *string `json:"type,omitempty"`
	UID         *string `json:"uid,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	Email       *string `json:"email,omitempty"`
	Created     *Time   `json:"created,omitempty"`
	Updated     *Time   `json:"updated,omitempty"`
}

// ListPrincipalsOptions specifies options for listing principals
type ListPrincipalsOptions struct {
	ListOptions
	Type *string `url:"type,omitempty"`
}

// ListPrincipals lists all principals
func (s *PrincipalsService) ListPrincipals(ctx context.Context, opt *ListPrincipalsOptions) ([]*Principal, *Response, error) {
	req := s.client.client.R().SetContext(ctx)

	// Add query parameters if options provided
	if opt != nil {
		buildQueryParams(req, &opt.ListOptions)

		if opt.Type != nil {
			req.SetQueryParam("type", *opt.Type)
		}
	}

	var principals []*Principal
	req.SetSuccessResult(&principals)

	resp, err := req.Get("principals")
	if err != nil {
		return nil, &Response{Response: resp}, err
	}

	if err := s.client.checkResponse(resp); err != nil {
		return nil, &Response{Response: resp}, err
	}

	response := &Response{Response: resp}
	s.client.parsePaginationHeaders(response)

	return principals, response, nil
}

// GetPrincipal retrieves a specific principal by ID
func (s *PrincipalsService) GetPrincipal(ctx context.Context, principalID int64) (*Principal, *Response, error) {
	path := fmt.Sprintf("principals/%d", principalID)
	var principal Principal
	resp, err := s.client.Get(ctx, path, &principal)
	if err != nil {
		return nil, resp, err
	}
	return &principal, resp, nil
}
