// Copyright (c) 2025-2025 All rights reserved.
//
// The original source code is licensed under the Apache License 2.0.
//
// You may review the terms of both licenses in the LICENSE file.

package gitness

import (
	"context"
)

// AuthService handles authentication related methods
type AuthService struct {
	client *Client
}

// LoginRequest represents a login request
type LoginRequest struct {
	LoginIdentifier *string `json:"login_identifier,omitempty"`
	Password        *string `json:"password,omitempty"`
}

// LoginResponse represents a login response
type LoginResponse struct {
	AccessToken *string    `json:"access_token,omitempty"`
	Principal   *Principal `json:"principal,omitempty"`
}

// RegisterRequest represents a registration request
type RegisterRequest struct {
	UID         *string `json:"uid,omitempty"`
	Email       *string `json:"email,omitempty"`
	Password    *string `json:"password,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
}

// Login authenticates a user
func (s *AuthService) Login(ctx context.Context, opt *LoginRequest) (*LoginResponse, *Response, error) {
	var loginResp LoginResponse
	resp, err := s.client.Post(ctx, "login", opt, &loginResp)
	if err != nil {
		return nil, resp, err
	}
	return &loginResp, resp, nil
}

// Logout logs out the current user
func (s *AuthService) Logout(ctx context.Context) (*Response, error) {
	resp, err := s.client.Post(ctx, "logout", nil, nil)
	return resp, err
}

// Register creates a new user account
func (s *AuthService) Register(ctx context.Context, opt *RegisterRequest) (*Principal, *Response, error) {
	var principal Principal
	resp, err := s.client.Post(ctx, "register", opt, &principal)
	if err != nil {
		return nil, resp, err
	}
	return &principal, resp, nil
}
