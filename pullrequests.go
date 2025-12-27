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

// PullRequestsService handles communication with pull request related methods
type PullRequestsService struct {
	client *Client
}

// PullRequest represents a Gitness pull request
type PullRequest struct {
	ID               *int64            `json:"id,omitempty"`
	Number           *int64            `json:"number,omitempty"`
	CreatedBy        *int64            `json:"created_by,omitempty"`
	Created          *Time             `json:"created,omitempty"`
	Updated          *Time             `json:"updated,omitempty"`
	Edited           *Time             `json:"edited,omitempty"`
	State            *string           `json:"state,omitempty"`
	IsDraft          *bool             `json:"is_draft,omitempty"`
	Title            *string           `json:"title,omitempty"`
	Description      *string           `json:"description,omitempty"`
	SourceRepoID     *int64            `json:"source_repo_id,omitempty"`
	SourceBranch     *string           `json:"source_branch,omitempty"`
	TargetRepoID     *int64            `json:"target_repo_id,omitempty"`
	TargetBranch     *string           `json:"target_branch,omitempty"`
	MergeMethod      *string           `json:"merge_method,omitempty"`
	MergeCheckStatus *string           `json:"merge_check_status,omitempty"`
	MergeSHA         *string           `json:"merge_sha,omitempty"`
	MergedBy         *int64            `json:"merged_by,omitempty"`
	Merged           *Time             `json:"merged,omitempty"`
	Stats            *PullRequestStats `json:"stats,omitempty"`
	Author           *PrincipalInfo    `json:"author,omitempty"`
	Merger           *PrincipalInfo    `json:"merger,omitempty"`
	Labels           []Label           `json:"labels,omitempty"`
	Reviewers        []Reviewer        `json:"reviewers,omitempty"`
}

// PullRequestStats represents pull request statistics
type PullRequestStats struct {
	Commits         *int `json:"commits,omitempty"`
	FilesChanged    *int `json:"files_changed,omitempty"`
	Additions       *int `json:"additions,omitempty"`
	Deletions       *int `json:"deletions,omitempty"`
	Conversations   *int `json:"conversations,omitempty"`
	UnresolvedCount *int `json:"unresolved_count,omitempty"`
}

// PrincipalInfo represents basic principal information
type PrincipalInfo struct {
	ID          *int64  `json:"id,omitempty"`
	UID         *string `json:"uid,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	Email       *string `json:"email,omitempty"`
	Type        *string `json:"type,omitempty"`
}

// Label represents a repository label
type Label struct {
	ID    *int64  `json:"id,omitempty"`
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
	Color *string `json:"color,omitempty"`
	Scope *string `json:"scope,omitempty"`
}

// Reviewer represents a pull request reviewer
type Reviewer struct {
	Principal      *PrincipalInfo `json:"principal,omitempty"`
	Type           *string        `json:"type,omitempty"`
	ReviewDecision *string        `json:"review_decision,omitempty"`
	SHA            *string        `json:"sha,omitempty"`
	Created        *Time          `json:"created,omitempty"`
	Updated        *Time          `json:"updated,omitempty"`
}

// CreatePullRequestOptions specifies options for creating a pull request
type CreatePullRequestOptions struct {
	Title        *string `json:"title,omitempty"`
	Description  *string `json:"description,omitempty"`
	SourceBranch *string `json:"source_branch,omitempty"`
	TargetBranch *string `json:"target_branch,omitempty"`
	IsDraft      *bool   `json:"is_draft,omitempty"`
}

// UpdatePullRequestOptions specifies options for updating a pull request
type UpdatePullRequestOptions struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
}

// StatePullRequestOptions specifies options for changing pull request state
type StatePullRequestOptions struct {
	State *string `json:"state,omitempty"`
}

// ListPullRequestsOptions specifies options for listing pull requests
type ListPullRequestsOptions struct {
	ListOptions
	State        *string `url:"state,omitempty"`
	SourceBranch *string `url:"source_branch,omitempty"`
	TargetBranch *string `url:"target_branch,omitempty"`
	CreatedBy    *int64  `url:"created_by,omitempty"`
}

// MergePullRequestOptions specifies options for merging a pull request
type MergePullRequestOptions struct {
	Method        *string `json:"method,omitempty"`
	CommitMessage *string `json:"commit_message,omitempty"`
	SourceSHA     *string `json:"source_sha,omitempty"`
	BypassRules   *bool   `json:"bypass_rules,omitempty"`
	DryRun        *bool   `json:"dry_run,omitempty"`
	DryRunRules   *bool   `json:"dry_run_rules,omitempty"`
}

// PullReqActivitySuggestionsMetadata contains metadata for code comment suggestions
type PullReqActivitySuggestionsMetadata struct {
	CheckSums        []string `json:"check_sums,omitempty"`
	AppliedCheckSum  string   `json:"applied_check_sum,omitempty"`
	AppliedCommitSHA string   `json:"applied_commit_sha,omitempty"`
}

// PullReqActivityMentionsMetadata contains metadata for code comment mentions
type PullReqActivityMentionsMetadata struct {
	IDs []int64 `json:"ids,omitempty"`
}

// PullReqActivityMetadata contains metadata related to pull request activity
type PullReqActivityMetadata struct {
	Suggestions *PullReqActivitySuggestionsMetadata `json:"suggestions,omitempty"`
	Mentions    *PullReqActivityMentionsMetadata    `json:"mentions,omitempty"`
}

// PullRequestActivity represents pull request activity/comment
type PullRequestActivity struct {
	ID         *int64                   `json:"id,omitempty"`
	Type       *string                  `json:"type,omitempty"`
	Kind       *string                  `json:"kind,omitempty"`
	Text       *string                  `json:"text,omitempty"`
	PayloadRaw *string                  `json:"payload,omitempty"`
	ReplyTo    *int64                   `json:"reply_to,omitempty"`
	Order      *int64                   `json:"order,omitempty"`
	SubOrder   *int64                   `json:"sub_order,omitempty"`
	Created    *Time                    `json:"created,omitempty"`
	Updated    *Time                    `json:"updated,omitempty"`
	Edited     *Time                    `json:"edited,omitempty"`
	Author     *PrincipalInfo           `json:"author,omitempty"`
	Metadata   *PullReqActivityMetadata `json:"metadata,omitempty"`
}

// CreatePullRequestCommentOptions specifies options for creating a pull request comment
type CreatePullRequestCommentOptions struct {
	Text    *string `json:"text,omitempty"`
	ReplyTo *int64  `json:"reply_to,omitempty"`
}

// CreatePullRequest creates a new pull request
func (s *PullRequestsService) CreatePullRequest(ctx context.Context, repoPath string, opt *CreatePullRequestOptions) (*PullRequest, *Response, error) {
	path := fmt.Sprintf("repos/%s/pullreq", repoPath)
	var pullRequest PullRequest
	resp, err := s.client.Post(ctx, path, opt, &pullRequest)
	if err != nil {
		return nil, resp, err
	}
	return &pullRequest, resp, nil
}

// ListPullRequests lists pull requests for a repository
func (s *PullRequestsService) ListPullRequests(ctx context.Context, repoPath string, opt *ListPullRequestsOptions) ([]*PullRequest, *Response, error) {
	path := fmt.Sprintf("repos/%s/pullreq", repoPath)
	req := s.client.client.R().SetContext(ctx)

	// Add query parameters if options provided
	if opt != nil {
		// Add common query parameters
		buildQueryParams(req, &opt.ListOptions)

		// Add specific query parameters
		if opt.State != nil {
			req.SetQueryParam("state", *opt.State)
		}
		if opt.SourceBranch != nil {
			req.SetQueryParam("source_branch", *opt.SourceBranch)
		}
		if opt.TargetBranch != nil {
			req.SetQueryParam("target_branch", *opt.TargetBranch)
		}
		if opt.CreatedBy != nil {
			req.SetQueryParam("created_by", fmt.Sprintf("%d", *opt.CreatedBy))
		}
	}

	var pullRequests []*PullRequest
	req.SetSuccessResult(&pullRequests)

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

	return pullRequests, response, nil
}

// GetPullRequest retrieves a specific pull request
func (s *PullRequestsService) GetPullRequest(ctx context.Context, repoPath string, pullRequestNumber int64) (*PullRequest, *Response, error) {
	path := fmt.Sprintf("repos/%s/pullreq/%d", repoPath, pullRequestNumber)
	var pullRequest PullRequest
	resp, err := s.client.Get(ctx, path, &pullRequest)
	if err != nil {
		return nil, resp, err
	}
	return &pullRequest, resp, nil
}

// UpdatePullRequest updates a pull request
func (s *PullRequestsService) UpdatePullRequest(ctx context.Context, repoPath string, pullRequestNumber int64, opt *UpdatePullRequestOptions) (*PullRequest, *Response, error) {
	path := fmt.Sprintf("repos/%s/pullreq/%d", repoPath, pullRequestNumber)
	var pullRequest PullRequest
	resp, err := s.client.Patch(ctx, path, opt, &pullRequest)
	if err != nil {
		return nil, resp, err
	}
	return &pullRequest, resp, nil
}

// SetPullRequestState changes the state of a pull request (open, closed, merged)
func (s *PullRequestsService) SetPullRequestState(ctx context.Context, repoPath string, pullRequestNumber int64, opt *StatePullRequestOptions) (*PullRequest, *Response, error) {
	path := fmt.Sprintf("repos/%s/pullreq/%d/state", repoPath, pullRequestNumber)
	var pullRequest PullRequest
	resp, err := s.client.Post(ctx, path, opt, &pullRequest)
	if err != nil {
		return nil, resp, err
	}
	return &pullRequest, resp, nil
}

// MergePullRequest merges a pull request
func (s *PullRequestsService) MergePullRequest(ctx context.Context, repoPath string, pullRequestNumber int64, opt *MergePullRequestOptions) (*PullRequest, *Response, error) {
	path := fmt.Sprintf("repos/%s/pullreq/%d/merge", repoPath, pullRequestNumber)
	var pullRequest PullRequest
	resp, err := s.client.Post(ctx, path, opt, &pullRequest)
	if err != nil {
		return nil, resp, err
	}
	return &pullRequest, resp, nil
}

// ListPullRequestActivity lists activities/comments for a pull request
func (s *PullRequestsService) ListPullRequestActivity(ctx context.Context, repoPath string, pullRequestNumber int64, opt *ListOptions) ([]*PullRequestActivity, *Response, error) {
	path := fmt.Sprintf("repos/%s/pullreq/%d/activities", repoPath, pullRequestNumber)
	var activities []*PullRequestActivity
	resp, err := s.client.performListRequest(ctx, path, opt, &activities)
	if err != nil {
		return nil, resp, err
	}
	return activities, resp, nil
}

// CreatePullRequestComment creates a comment on a pull request
func (s *PullRequestsService) CreatePullRequestComment(ctx context.Context, repoPath string, pullRequestNumber int64, opt *CreatePullRequestCommentOptions) (*PullRequestActivity, *Response, error) {
	path := fmt.Sprintf("repos/%s/pullreq/%d/comments", repoPath, pullRequestNumber)
	var comment PullRequestActivity
	resp, err := s.client.Post(ctx, path, opt, &comment)
	if err != nil {
		return nil, resp, err
	}
	return &comment, resp, nil
}

// AddPullRequestReviewer adds a reviewer to a pull request
func (s *PullRequestsService) AddPullRequestReviewer(ctx context.Context, repoPath string, pullRequestNumber int64, reviewerUID string) (*Response, error) {
	path := fmt.Sprintf("repos/%s/pullreq/%d/reviewers/%s", repoPath, pullRequestNumber, reviewerUID)
	resp, err := s.client.Put(ctx, path, nil, nil)
	return resp, err
}

// RemovePullRequestReviewer removes a reviewer from a pull request
func (s *PullRequestsService) RemovePullRequestReviewer(ctx context.Context, repoPath string, pullRequestNumber int64, reviewerUID string) (*Response, error) {
	path := fmt.Sprintf("repos/%s/pullreq/%d/reviewers/%s", repoPath, pullRequestNumber, reviewerUID)
	resp, err := s.client.Delete(ctx, path, nil)
	return resp, err
}

// ListPullRequestReviewers lists reviewers for a pull request
func (s *PullRequestsService) ListPullRequestReviewers(ctx context.Context, repoPath string, pullRequestNumber int64) ([]*Reviewer, *Response, error) {
	path := fmt.Sprintf("repos/%s/pullreq/%d/reviewers", repoPath, pullRequestNumber)
	var reviewers []*Reviewer
	resp, err := s.client.Get(ctx, path, &reviewers)
	if err != nil {
		return nil, resp, err
	}
	return reviewers, resp, nil
}

// UserGroupReviewer represents a user group reviewer for a pull request
type UserGroupReviewer struct {
	ID            *int64                 `json:"id,omitempty"`
	UserGroupID   *int64                 `json:"user_group_id,omitempty"`
	AddedBy       *PrincipalInfo         `json:"added_by,omitempty"`
	Created       *Time                  `json:"created,omitempty"`
	Updated       *Time                  `json:"updated,omitempty"`
	Decision      *PullReqReviewDecision `json:"decision,omitempty"`
	UserDecisions []UserReviewDecision   `json:"user_decisions,omitempty"`
}

// UserReviewDecision represents an individual user's review decision within a user group
type UserReviewDecision struct {
	UserID   *int64                 `json:"user_id,omitempty"`
	UserInfo *PrincipalInfo         `json:"user_info,omitempty"`
	Decision *PullReqReviewDecision `json:"decision,omitempty"`
	Created  *Time                  `json:"created,omitempty"`
}

// PullReqReviewDecision represents the review decision enum
type PullReqReviewDecision string

const (
	PullReqReviewDecisionApproved         PullReqReviewDecision = "approved"
	PullReqReviewDecisionRequestedChanges PullReqReviewDecision = "changereq"
	PullReqReviewDecisionPending          PullReqReviewDecision = "pending"
)

// CombinedReviewers represents combined individual and user group reviewers
type CombinedReviewers struct {
	Reviewers          []*Reviewer          `json:"reviewers,omitempty"`
	UserGroupReviewers []*UserGroupReviewer `json:"usergroup_reviewers,omitempty"`
}

// UserGroupReviewerAddRequest represents a request to add a user group reviewer
type UserGroupReviewerAddRequest struct {
	UserGroupID *int64 `json:"usergroup_id,omitempty"`
}

// ListPullRequestCombinedReviewers lists both individual and user group reviewers for a pull request
func (s *PullRequestsService) ListPullRequestCombinedReviewers(ctx context.Context, repoPath string, pullRequestNumber int64) (*CombinedReviewers, *Response, error) {
	path := fmt.Sprintf("repos/%s/pullreq/%d/reviewers/combined", repoPath, pullRequestNumber)
	var combinedReviewers CombinedReviewers
	resp, err := s.client.Get(ctx, path, &combinedReviewers)
	if err != nil {
		return nil, resp, err
	}
	return &combinedReviewers, resp, nil
}

// AddPullRequestUserGroupReviewer adds a user group reviewer to a pull request
func (s *PullRequestsService) AddPullRequestUserGroupReviewer(ctx context.Context, repoPath string, pullRequestNumber int64, userGroupID int64) (*UserGroupReviewer, *Response, error) {
	path := fmt.Sprintf("repos/%s/pullreq/%d/reviewers/usergroups", repoPath, pullRequestNumber)
	req := &UserGroupReviewerAddRequest{
		UserGroupID: &userGroupID,
	}

	var userGroupReviewer UserGroupReviewer
	resp, err := s.client.Put(ctx, path, req, &userGroupReviewer)
	if err != nil {
		return nil, resp, err
	}
	return &userGroupReviewer, resp, nil
}

// RemovePullRequestUserGroupReviewer removes a user group reviewer from a pull request
func (s *PullRequestsService) RemovePullRequestUserGroupReviewer(ctx context.Context, repoPath string, pullRequestNumber int64, userGroupID int64) (*Response, error) {
	path := fmt.Sprintf("repos/%s/pullreq/%d/reviewers/usergroups/%d", repoPath, pullRequestNumber, userGroupID)
	resp, err := s.client.Delete(ctx, path, nil)
	return resp, err
}
