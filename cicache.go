// Copyright (c) 2025-2025 All rights reserved.
//
// The original source code is licensed under the Apache License 2.0.
//
// You may review the terms of both licenses in the LICENSE file.

package gitness

import (
	"context"
	"fmt"
	"io"
)

// CiCacheService handles communication with CI cache related methods
type CiCacheService struct {
	client *Client
}

// CiCacheEntry represents a CI cache entry
type CiCacheEntry struct {
	Key      *string `json:"key,omitempty"`
	Size     *int64  `json:"size,omitempty"`
	Created  *Time   `json:"created,omitempty"`
	Accessed *Time   `json:"accessed,omitempty"`
	Version  *int    `json:"version,omitempty"`
}

// UploadCiCacheRequest represents a request to upload CI cache
type UploadCiCacheRequest struct {
	Key     *string   `json:"key,omitempty"`
	Version *int      `json:"version,omitempty"`
	Data    io.Reader `json:"-"`
}

// UploadCiCache uploads a CI cache entry
func (s *CiCacheService) UploadCiCache(ctx context.Context, key string, version int, data io.Reader) (*CiCacheEntry, *Response, error) {
	path := fmt.Sprintf("ci/cache/%s", key)
	
	req := s.client.client.R().SetContext(ctx)
	if version > 0 {
		req.SetQueryParam("version", fmt.Sprintf("%d", version))
	}
	
	// Set the body data for upload
	req.SetBody(data)
	req.SetContentType("application/octet-stream")

	var cacheEntry CiCacheEntry
	req.SetSuccessResult(&cacheEntry)

	resp, err := req.Put(path)
	if err != nil {
		return nil, &Response{Response: resp}, err
	}

	if err := s.client.checkResponse(resp); err != nil {
		return nil, &Response{Response: resp}, err
	}

	return &cacheEntry, &Response{Response: resp}, nil
}

// GetCiCacheOptions specifies optional parameters for getting CI cache
type GetCiCacheOptions struct {
	Version *int `url:"version,omitempty"`
}

// GetCiCache retrieves a CI cache entry by key
func (s *CiCacheService) GetCiCache(ctx context.Context, key string, opt *GetCiCacheOptions) (io.ReadCloser, *Response, error) {
	path := fmt.Sprintf("ci/cache/%s", key)
	req := s.client.client.R().SetContext(ctx)

	if opt != nil && opt.Version != nil {
		req.SetQueryParam("version", fmt.Sprintf("%d", *opt.Version))
	}

	resp, err := req.Get(path)
	if err != nil {
		return nil, &Response{Response: resp}, err
	}

	if err := s.client.checkResponse(resp); err != nil {
		return nil, &Response{Response: resp}, err
	}

	return resp.Body, &Response{Response: resp}, nil
}

// ListCiCacheOptions specifies optional parameters for listing CI cache entries
type ListCiCacheOptions struct {
	ListOptions
	KeyPrefix *string `url:"key_prefix,omitempty"`
}

// ListCiCache lists CI cache entries with optional filtering
func (s *CiCacheService) ListCiCache(ctx context.Context, opt *ListCiCacheOptions) ([]*CiCacheEntry, *Response, error) {
	req := s.client.client.R().SetContext(ctx)

	if opt != nil {
		buildQueryParams(req, &opt.ListOptions)
		if opt.KeyPrefix != nil {
			req.SetQueryParam("key_prefix", *opt.KeyPrefix)
		}
	}

	var entries []*CiCacheEntry
	req.SetSuccessResult(&entries)

	resp, err := req.Get("ci/cache")
	if err != nil {
		return nil, &Response{Response: resp}, err
	}

	if err := s.client.checkResponse(resp); err != nil {
		return nil, &Response{Response: resp}, err
	}

	response := &Response{Response: resp}
	s.client.parsePaginationHeaders(response)

	return entries, response, nil
}

// DeleteCiCache deletes a CI cache entry by key
func (s *CiCacheService) DeleteCiCache(ctx context.Context, key string) (*Response, error) {
	path := fmt.Sprintf("ci/cache/%s", key)
	resp, err := s.client.Delete(ctx, path, nil)
	return resp, err
}

// ClearCiCache clears all CI cache entries
func (s *CiCacheService) ClearCiCache(ctx context.Context) (*Response, error) {
	resp, err := s.client.Delete(ctx, "ci/cache", nil)
	return resp, err
}