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

// AdminService handles communication with admin related methods
type AdminService struct {
	client *Client
}

// AuditService handles communication with audit related methods
type AuditService struct {
	client *Client
}

// AuditLog represents an audit log entry
type AuditLog struct {
	ID                   *int64  `json:"id,omitempty"`
	Created              *Time   `json:"created,omitempty"`
	Action               *string `json:"action,omitempty"`
	ResourceType         *string `json:"resource_type,omitempty"`
	ResourceIdentifier   *string `json:"resource_identifier,omitempty"`
	PrincipalUID         *string `json:"principal_uid,omitempty"`
	PrincipalDisplayName *string `json:"principal_display_name,omitempty"`
	Data                 *string `json:"data,omitempty"`
}

// ListAuditLogsOptions specifies the optional parameters for listing audit logs
type ListAuditLogsOptions struct {
	ListOptions
	UserUID            *string `url:"user_uid,omitempty"`
	Action             *string `url:"action,omitempty"`
	ResourceType       *string `url:"resource_type,omitempty"`
	ResourceIdentifier *string `url:"resource_identifier,omitempty"`
	From               *Time   `url:"from,omitempty"`
	To                 *Time   `url:"to,omitempty"`
}

// ListAuditLogs lists audit logs with optional filtering and pagination
func (s *AuditService) ListAuditLogs(ctx context.Context, opt *ListAuditLogsOptions) ([]*AuditLog, *Response, error) {
	req := s.client.client.R().SetContext(ctx)

	// Add query parameters if options provided
	if opt != nil {
		// Add common query parameters
		buildQueryParams(req, &opt.ListOptions)

		// Add specific query parameters
		if opt.UserUID != nil {
			req.SetQueryParam("user_uid", *opt.UserUID)
		}
		if opt.Action != nil {
			req.SetQueryParam("action", *opt.Action)
		}
		if opt.ResourceType != nil {
			req.SetQueryParam("resource_type", *opt.ResourceType)
		}
		if opt.ResourceIdentifier != nil {
			req.SetQueryParam("resource_identifier", *opt.ResourceIdentifier)
		}
		if opt.From != nil {
			req.SetQueryParam("from", opt.From.String())
		}
		if opt.To != nil {
			req.SetQueryParam("to", opt.To.String())
		}
	}

	var logs []*AuditLog
	req.SetSuccessResult(&logs)

	resp, err := req.Get("admin/audit")
	if err != nil {
		return nil, &Response{Response: resp}, err
	}

	if err := s.client.checkResponse(resp); err != nil {
		return nil, &Response{Response: resp}, err
	}

	// Parse pagination headers
	response := &Response{Response: resp}
	s.client.parsePaginationHeaders(response)

	return logs, response, nil
}

// GetAuditLog retrieves a specific audit log entry by ID
func (s *AuditService) GetAuditLog(ctx context.Context, id int64) (*AuditLog, *Response, error) {
	path := fmt.Sprintf("admin/audit/%d", id)
	var log AuditLog
	resp, err := s.client.Get(ctx, path, &log)
	if err != nil {
		return nil, resp, err
	}
	return &log, resp, nil
}

// CleanupAuditLogs initiates cleanup of audit logs
func (s *AuditService) CleanupAuditLogs(ctx context.Context) (*Response, error) {
	resp, err := s.client.Post(ctx, "admin/audit/cleanup", nil, nil)
	return resp, err
}

// User represents a Gitness user
type User struct {
	UID         *string `json:"uid,omitempty"`
	Email       *string `json:"email,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	Admin       *bool   `json:"admin,omitempty"`
	Blocked     *bool   `json:"blocked,omitempty"`
	Created     *Time   `json:"created,omitempty"`
	Updated     *Time   `json:"updated,omitempty"`
}

// ListUsersOptions specifies the optional parameters for listing users
type ListUsersOptions struct {
	ListOptions
	Admin   *bool `url:"admin,omitempty"`
	Blocked *bool `url:"blocked,omitempty"`
}

// ListUsers lists users with optional filtering
func (s *AdminService) ListUsers(ctx context.Context, opt *ListUsersOptions) ([]*User, *Response, error) {
	req := s.client.client.R().SetContext(ctx)

	// Add query parameters if options provided
	if opt != nil {
		if opt.Page != nil {
			req.SetQueryParam("page", fmt.Sprintf("%d", *opt.Page))
		}
		if opt.Limit != nil {
			req.SetQueryParam("limit", fmt.Sprintf("%d", *opt.Limit))
		}
		if opt.Sort != nil {
			req.SetQueryParam("sort", *opt.Sort)
		}
		if opt.Order != nil {
			req.SetQueryParam("order", *opt.Order)
		}
		if opt.Query != nil {
			req.SetQueryParam("query", *opt.Query)
		}
		if opt.Admin != nil {
			req.SetQueryParam("admin", fmt.Sprintf("%t", *opt.Admin))
		}
		if opt.Blocked != nil {
			req.SetQueryParam("blocked", fmt.Sprintf("%t", *opt.Blocked))
		}
	}

	var users []*User
	req.SetSuccessResult(&users)

	resp, err := req.Get("admin/users")
	if err != nil {
		return nil, &Response{Response: resp}, err
	}

	if err := s.client.checkResponse(resp); err != nil {
		return nil, &Response{Response: resp}, err
	}

	// Parse pagination headers
	response := &Response{Response: resp}
	s.client.parsePaginationHeaders(response)

	return users, response, nil
}

// GetUser retrieves a specific user by UID
func (s *AdminService) GetUser(ctx context.Context, userUID string) (*User, *Response, error) {
	path := fmt.Sprintf("admin/users/%s", userUID)
	var user User
	resp, err := s.client.Get(ctx, path, &user)
	if err != nil {
		return nil, resp, err
	}
	return &user, resp, nil
}

// UpdateUserAdminStatus updates user's admin status
func (s *AdminService) UpdateUserAdminStatus(ctx context.Context, userUID string, admin bool) (*User, *Response, error) {
	path := fmt.Sprintf("admin/users/%s/admin", userUID)
	payload := map[string]bool{"admin": admin}

	var user User
	resp, err := s.client.Patch(ctx, path, payload, &user)
	if err != nil {
		return nil, resp, err
	}
	return &user, resp, nil
}

// UpdateUserBlockedStatus updates user's blocked status
func (s *AdminService) UpdateUserBlockedStatus(ctx context.Context, userUID string, blocked bool) (*User, *Response, error) {
	path := fmt.Sprintf("admin/users/%s/blocked", userUID)
	payload := map[string]bool{"blocked": blocked}

	var user User
	resp, err := s.client.Patch(ctx, path, payload, &user)
	if err != nil {
		return nil, resp, err
	}
	return &user, resp, nil
}

// CreateUserRequest represents a request to create a new user
type CreateUserRequest struct {
	UID         *string `json:"uid,omitempty"`
	Email       *string `json:"email,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	Password    *string `json:"password,omitempty"`
	Admin       *bool   `json:"admin,omitempty"`
}

// CreateUser creates a new user
func (s *AdminService) CreateUser(ctx context.Context, user *CreateUserRequest) (*User, *Response, error) {
	var newUser User
	resp, err := s.client.Post(ctx, "admin/users", user, &newUser)
	if err != nil {
		return nil, resp, err
	}
	return &newUser, resp, nil
}

// UpdateUserRequest represents a request to update user information
type UpdateUserRequest struct {
	Email       *string `json:"email,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
}

// UpdateUser updates user information
func (s *AdminService) UpdateUser(ctx context.Context, userUID string, user *UpdateUserRequest) (*User, *Response, error) {
	path := fmt.Sprintf("admin/users/%s", userUID)
	var updatedUser User
	resp, err := s.client.Patch(ctx, path, user, &updatedUser)
	if err != nil {
		return nil, resp, err
	}
	return &updatedUser, resp, nil
}

// DeleteUser deletes a user by UID
func (s *AdminService) DeleteUser(ctx context.Context, userUID string) (*Response, error) {
	path := fmt.Sprintf("admin/users/%s", userUID)
	resp, err := s.client.Delete(ctx, path, nil)
	return resp, err
}

// LDAPUser represents an LDAP user search result
type LDAPUser struct {
	UID         *string `json:"uid,omitempty"`
	Email       *string `json:"email,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
}

// SearchLDAPUsersOptions specifies the optional parameters for searching LDAP users
type SearchLDAPUsersOptions struct {
	ListOptions
	Query *string `url:"query,omitempty"`
}

// SearchLDAPUsers searches for LDAP users
func (s *AdminService) SearchLDAPUsers(ctx context.Context, opt *SearchLDAPUsersOptions) ([]*LDAPUser, *Response, error) {
	req := s.client.client.R().SetContext(ctx)

	if opt != nil {
		buildQueryParams(req, &opt.ListOptions)
		if opt.Query != nil {
			req.SetQueryParam("query", *opt.Query)
		}
	}

	var users []*LDAPUser
	req.SetSuccessResult(&users)

	resp, err := req.Get("admin/ldap/users")
	if err != nil {
		return nil, &Response{Response: resp}, err
	}

	if err := s.client.checkResponse(resp); err != nil {
		return nil, &Response{Response: resp}, err
	}

	response := &Response{Response: resp}
	s.client.parsePaginationHeaders(response)

	return users, response, nil
}

// SyncLDAPUsersRequest represents a request to sync LDAP users
type SyncLDAPUsersRequest struct {
	UserUIDs []string `json:"user_uids,omitempty"`
}

// SyncLDAPUsersResponse represents the response from LDAP sync operation
type SyncLDAPUsersResponse struct {
	Synchronized *int `json:"synchronized,omitempty"`
	Failed       *int `json:"failed,omitempty"`
}

// SyncLDAPUsers synchronizes LDAP users
func (s *AdminService) SyncLDAPUsers(ctx context.Context, req *SyncLDAPUsersRequest) (*SyncLDAPUsersResponse, *Response, error) {
	var syncResp SyncLDAPUsersResponse
	resp, err := s.client.Post(ctx, "admin/ldap/users/sync", req, &syncResp)
	if err != nil {
		return nil, resp, err
	}
	return &syncResp, resp, nil
}
