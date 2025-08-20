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

// SpacesService handles communication with space related methods
type SpacesService struct {
	client *Client
}

// Space represents a Gitness space
type Space struct {
	ID          *int64  `json:"id,omitempty"`
	ParentID    *int64  `json:"parent_id,omitempty"`
	Identifier  *string `json:"identifier,omitempty"`
	Path        *string `json:"path,omitempty"`
	Description *string `json:"description,omitempty"`
	IsPublic    *bool   `json:"is_public,omitempty"`
	CreatedBy   *int64  `json:"created_by,omitempty"`
	Created     *Time   `json:"created,omitempty"`
	Updated     *Time   `json:"updated,omitempty"`
}

// CreateSpaceOptions specifies options for creating a space
type CreateSpaceOptions struct {
	Identifier  *string `json:"identifier,omitempty"`
	ParentRef   *string `json:"parent_ref,omitempty"`
	Description *string `json:"description,omitempty"`
	IsPublic    *bool   `json:"is_public,omitempty"`
}

// UpdateSpaceOptions specifies options for updating a space
type UpdateSpaceOptions struct {
	Description *string `json:"description,omitempty"`
	IsPublic    *bool   `json:"is_public,omitempty"`
}

// ListSpacesOptions specifies options for listing spaces
type ListSpacesOptions struct {
	ListOptions
	Recursive *bool `url:"recursive,omitempty"`
}

// GetSpace retrieves a space by its reference
func (s *SpacesService) GetSpace(ctx context.Context, spaceRef string) (*Space, *Response, error) {
	path := fmt.Sprintf("spaces/%s", spaceRef)
	var space Space
	resp, err := s.client.Get(ctx, path, &space)
	if err != nil {
		return nil, resp, err
	}
	return &space, resp, nil
}

// ListSpaces lists spaces
func (s *SpacesService) ListSpaces(ctx context.Context, opt *ListSpacesOptions) ([]*Space, *Response, error) {
	var spaces []*Space

	req := s.client.client.R().SetContext(ctx)

	// Add query parameters if options provided
	if opt != nil {
		buildQueryParams(req, &opt.ListOptions)

		if opt.Recursive != nil {
			req.SetQueryParam("recursive", fmt.Sprintf("%t", *opt.Recursive))
		}
	}

	req.SetSuccessResult(&spaces)

	resp, err := req.Get("spaces")
	if err != nil {
		return nil, &Response{Response: resp}, err
	}

	if err := s.client.checkResponse(resp); err != nil {
		return nil, &Response{Response: resp}, err
	}

	response := &Response{Response: resp}
	s.client.parsePaginationHeaders(response)

	return spaces, response, nil
}

// CreateSpace creates a new space
func (s *SpacesService) CreateSpace(ctx context.Context, opt *CreateSpaceOptions) (*Space, *Response, error) {
	var space Space
	resp, err := s.client.Post(ctx, "spaces", opt, &space)
	if err != nil {
		return nil, resp, err
	}
	return &space, resp, nil
}

// UpdateSpace updates a space
func (s *SpacesService) UpdateSpace(ctx context.Context, spaceRef string, opt *UpdateSpaceOptions) (*Space, *Response, error) {
	path := fmt.Sprintf("spaces/%s", spaceRef)
	var space Space
	resp, err := s.client.Patch(ctx, path, opt, &space)
	if err != nil {
		return nil, resp, err
	}
	return &space, resp, nil
}

// DeleteSpaceRequest represents options for deleting a space
type DeleteSpaceRequest struct {
	DeleteID *string `json:"delete_id,omitempty"`
}

// DeleteSpace deletes a space
func (s *SpacesService) DeleteSpace(ctx context.Context, spaceRef string, deleteID *string) (*Response, error) {
	path := fmt.Sprintf("spaces/%s", spaceRef)

	var payload *DeleteSpaceRequest
	if deleteID != nil {
		payload = &DeleteSpaceRequest{
			DeleteID: deleteID,
		}
	}

	resp, err := s.client.Delete(ctx, path, payload)
	return resp, err
}

// ListRepositories lists repositories in a space
func (s *SpacesService) ListRepositories(ctx context.Context, spaceRef string, opt *ListRepositoriesOptions) ([]*Repository, *Response, error) {
	path := fmt.Sprintf("spaces/%s/repos", spaceRef)
	var repositories []*Repository

	req := s.client.client.R().SetContext(ctx)

	// Add query parameters if options provided
	if opt != nil {
		buildQueryParams(req, &opt.ListOptions)

		if opt.Sort != nil {
			req.SetQueryParam("sort", *opt.Sort)
		}
		if opt.Order != nil {
			req.SetQueryParam("order", *opt.Order)
		}
		if opt.Query != nil {
			req.SetQueryParam("query", *opt.Query)
		}
	}

	req.SetSuccessResult(&repositories)

	resp, err := req.Get(path)
	if err != nil {
		return nil, &Response{Response: resp}, err
	}

	if err := s.client.checkResponse(resp); err != nil {
		return nil, &Response{Response: resp}, err
	}

	response := &Response{Response: resp}
	s.client.parsePaginationHeaders(response)

	return repositories, response, nil
}
