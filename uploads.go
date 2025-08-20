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

// UploadService handles communication with upload related methods
type UploadService struct {
	client *Client
}

// Upload represents an uploaded file
type Upload struct {
	Reference *string `json:"reference,omitempty"`
	FileName  *string `json:"file_name,omitempty"`
	FileSize  *int64  `json:"file_size,omitempty"`
	Checksum  *string `json:"checksum,omitempty"`
	Created   *Time   `json:"created,omitempty"`
}

// CreateUploadRequest represents the request to create an upload session
type CreateUploadRequest struct {
	FileName *string `json:"file_name,omitempty"`
	FileSize *int64  `json:"file_size,omitempty"`
}

// CreateUpload creates an upload session
func (s *UploadService) CreateUpload(ctx context.Context, repoPath string, fileName string, fileSize int64) (*Upload, *Response, error) {
	path := fmt.Sprintf("repos/%s/uploads", repoPath)

	payload := &CreateUploadRequest{
		FileName: &fileName,
		FileSize: &fileSize,
	}

	var upload Upload
	resp, err := s.client.Post(ctx, path, payload, &upload)
	if err != nil {
		return nil, resp, err
	}
	return &upload, resp, nil
}

// GetUpload retrieves upload information
func (s *UploadService) GetUpload(ctx context.Context, repoPath, fileRef string) (*Upload, *Response, error) {
	path := fmt.Sprintf("repos/%s/uploads/%s", repoPath, fileRef)
	var upload Upload
	resp, err := s.client.Get(ctx, path, &upload)
	if err != nil {
		return nil, resp, err
	}
	return &upload, resp, nil
}
