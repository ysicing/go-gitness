// Copyright (c) 2025-2025 All rights reserved.
//
// The original source code is licensed under the Apache License 2.0.
//
// You may review the terms of both licenses in the LICENSE file.

package gitness

import (
	"context"
	"fmt"
	"net/url"
)

// UsersService handles communication with user related methods
type UsersService struct {
	client *Client
}

// PublicKey represents a user's public key
type PublicKey struct {
	Identifier  *string `json:"identifier,omitempty"`
	Type        *string `json:"type,omitempty"`
	Content     *string `json:"content,omitempty"`
	Fingerprint *string `json:"fingerprint,omitempty"`
	Usage       *string `json:"usage,omitempty"`
	Created     *Time   `json:"created,omitempty"`
}

// PersonalAccessToken represents a user's personal access token
type PersonalAccessToken struct {
	Identifier *string `json:"identifier,omitempty"`
	Name       *string `json:"name,omitempty"`
	ExpiresAt  *Time   `json:"expires_at,omitempty"`
	IssuedAt   *Time   `json:"issued_at,omitempty"`
	LastUsedAt *Time   `json:"last_used_at,omitempty"`
}

// CreatePublicKeyOptions specifies options for creating a public key
type CreatePublicKeyOptions struct {
	Identifier *string `json:"identifier,omitempty"`
	Content    *string `json:"content,omitempty"`
	Usage      *string `json:"usage,omitempty"`
}

// CreateTokenOptions specifies options for creating a personal access token
type CreateTokenOptions struct {
	Identifier *string `json:"identifier,omitempty"`
	Lifetime   *int64  `json:"lifetime,omitempty"`
}

// UserMembership represents user's membership in spaces
type UserMembership struct {
	SpaceID   *int64  `json:"space_id,omitempty"`
	SpacePath *string `json:"space_path,omitempty"`
	Role      *string `json:"role,omitempty"`
	AddedBy   *int64  `json:"added_by,omitempty"`
	Added     *Time   `json:"added,omitempty"`
}

// ListPublicKeysOptions specifies options for listing public keys
type ListPublicKeysOptions struct {
	ListOptions
	Usage *string `url:"usage,omitempty"`
}

// ListTokensOptions specifies options for listing tokens
type ListTokensOptions struct {
	ListOptions
}

// UserFavorite represents a user's favorite resource
type UserFavorite struct {
	ResourceID   *int64  `json:"resource_id,omitempty"`
	ResourceType *string `json:"resource_type,omitempty"`
	ResourcePath *string `json:"resource_path,omitempty"`
	Added        *Time   `json:"added,omitempty"`
}

// GetCurrentUser retrieves the current authenticated user
func (s *UsersService) GetCurrentUser(ctx context.Context) (*User, *Response, error) {
	var user User
	resp, err := s.client.Get(ctx, "user", &user)
	if err != nil {
		return nil, resp, err
	}
	return &user, resp, nil
}

// GetUser retrieves a user by UID
func (s *UsersService) GetUser(ctx context.Context, userUID string) (*User, *Response, error) {
	path := fmt.Sprintf("users/%s", url.PathEscape(userUID))
	var user User
	resp, err := s.client.Get(ctx, path, &user)
	if err != nil {
		return nil, resp, err
	}
	return &user, resp, nil
}

// ListUserKeys lists user's public keys
func (s *UsersService) ListUserKeys(ctx context.Context, opt *ListPublicKeysOptions) ([]*PublicKey, *Response, error) {
	req := s.client.client.R().SetContext(ctx)

	// Add query parameters if options provided
	if opt != nil {
		buildQueryParams(req, &opt.ListOptions)

		if opt.Usage != nil {
			req.SetQueryParam("usage", *opt.Usage)
		}
	}

	var keys []*PublicKey
	req.SetSuccessResult(&keys)

	resp, err := req.Get("user/keys")
	if err != nil {
		return nil, &Response{Response: resp}, err
	}

	if err := s.client.checkResponse(resp); err != nil {
		return nil, &Response{Response: resp}, err
	}

	response := &Response{Response: resp}
	s.client.parsePaginationHeaders(response)

	return keys, response, nil
}

// CreateUserKey creates a public key for the user
func (s *UsersService) CreateUserKey(ctx context.Context, opt *CreatePublicKeyOptions) (*PublicKey, *Response, error) {
	var key PublicKey
	resp, err := s.client.Post(ctx, "user/keys", opt, &key)
	if err != nil {
		return nil, resp, err
	}
	return &key, resp, nil
}

// GetUserKey retrieves a specific public key
func (s *UsersService) GetUserKey(ctx context.Context, keyID string) (*PublicKey, *Response, error) {
	path := fmt.Sprintf("user/keys/%s", url.PathEscape(keyID))
	var key PublicKey
	resp, err := s.client.Get(ctx, path, &key)
	if err != nil {
		return nil, resp, err
	}
	return &key, resp, nil
}

// DeleteUserKey deletes a public key
func (s *UsersService) DeleteUserKey(ctx context.Context, keyID string) (*Response, error) {
	path := fmt.Sprintf("user/keys/%s", url.PathEscape(keyID))
	resp, err := s.client.Delete(ctx, path, nil)
	return resp, err
}

// ListUserTokens lists user's personal access tokens
func (s *UsersService) ListUserTokens(ctx context.Context, opt *ListTokensOptions) ([]*PersonalAccessToken, *Response, error) {
	req := s.client.client.R().SetContext(ctx)

	// Add query parameters if options provided
	if opt != nil {
		buildQueryParams(req, &opt.ListOptions)
	}

	var tokens []*PersonalAccessToken
	req.SetSuccessResult(&tokens)

	resp, err := req.Get("user/tokens")
	if err != nil {
		return nil, &Response{Response: resp}, err
	}

	if err := s.client.checkResponse(resp); err != nil {
		return nil, &Response{Response: resp}, err
	}

	response := &Response{Response: resp}
	s.client.parsePaginationHeaders(response)

	return tokens, response, nil
}

// CreateUserToken creates a personal access token
func (s *UsersService) CreateUserToken(ctx context.Context, opt *CreateTokenOptions) (*PersonalAccessToken, *Response, error) {
	var token PersonalAccessToken
	resp, err := s.client.Post(ctx, "user/tokens", opt, &token)
	if err != nil {
		return nil, resp, err
	}
	return &token, resp, nil
}

// DeleteUserToken deletes a personal access token
func (s *UsersService) DeleteUserToken(ctx context.Context, tokenID string) (*Response, error) {
	path := fmt.Sprintf("user/tokens/%s", url.PathEscape(tokenID))
	resp, err := s.client.Delete(ctx, path, nil)
	return resp, err
}

// ListUserMemberships lists user's space memberships
func (s *UsersService) ListUserMemberships(ctx context.Context) ([]*UserMembership, *Response, error) {
	var memberships []*UserMembership
	resp, err := s.client.Get(ctx, "user/memberships", &memberships)
	if err != nil {
		return nil, resp, err
	}
	return memberships, resp, nil
}

// ListUserFavorites lists user's favorite resources
func (s *UsersService) ListUserFavorites(ctx context.Context) ([]*UserFavorite, *Response, error) {
	var favorites []*UserFavorite
	resp, err := s.client.Get(ctx, "user/favorite", &favorites)
	if err != nil {
		return nil, resp, err
	}
	return favorites, resp, nil
}

// AddUserFavorite adds a resource to user's favorites
func (s *UsersService) AddUserFavorite(ctx context.Context, resourceID int64) (*UserFavorite, *Response, error) {
	path := fmt.Sprintf("user/favorite/%d", resourceID)
	var favorite UserFavorite
	resp, err := s.client.Post(ctx, path, nil, &favorite)
	if err != nil {
		return nil, resp, err
	}
	return &favorite, resp, nil
}

// RemoveUserFavorite removes a resource from user's favorites
func (s *UsersService) RemoveUserFavorite(ctx context.Context, resourceID int64) (*Response, error) {
	path := fmt.Sprintf("user/favorite/%d", resourceID)
	resp, err := s.client.Delete(ctx, path, nil)
	return resp, err
}
