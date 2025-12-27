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

// RepositoriesService handles communication with repository related methods
type RepositoriesService struct {
	client *Client
}

// Repository represents a Gitness repository
type Repository struct {
	ID             *int64  `json:"id,omitempty"`
	ParentID       *int64  `json:"parent_id,omitempty"`
	Identifier     *string `json:"identifier,omitempty"`
	Path           *string `json:"path,omitempty"`
	Description    *string `json:"description,omitempty"`
	IsPublic       *bool   `json:"is_public,omitempty"`
	CreatedBy      *int64  `json:"created_by,omitempty"`
	Created        *Time   `json:"created,omitempty"`
	Updated        *Time   `json:"updated,omitempty"`
	Size           *int64  `json:"size,omitempty"`
	SizeUpdated    *Time   `json:"size_updated,omitempty"`
	GitURL         *string `json:"git_url,omitempty"`
	DefaultBranch  *string `json:"default_branch,omitempty"`
	ForkID         *int64  `json:"fork_id,omitempty"`
	NumForks       *int    `json:"num_forks,omitempty"`
	NumPulls       *int    `json:"num_pulls,omitempty"`
	NumClosedPulls *int    `json:"num_closed_pulls,omitempty"`
	NumOpenPulls   *int    `json:"num_open_pulls,omitempty"`
	NumMergedPulls *int    `json:"num_merged_pulls,omitempty"`
	Importing      *bool   `json:"importing,omitempty"`
}

// Branch represents a repository branch
type Branch struct {
	Name   *string    `json:"name,omitempty"`
	SHA    *string    `json:"sha,omitempty"`
	Commit *CommitSHA `json:"commit,omitempty"`
}

// CommitSHA represents basic commit information
type CommitSHA struct {
	SHA       *string    `json:"sha,omitempty"`
	Message   *string    `json:"message,omitempty"`
	Author    *Committer `json:"author,omitempty"`
	Committer *Committer `json:"committer,omitempty"`
}

// Committer represents commit author/committer information
type Committer struct {
	Identity *Identity `json:"identity,omitempty"`
	When     *Time     `json:"when,omitempty"`
}

// Identity represents user identity
type Identity struct {
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
}

// CreateRepositoryOptions specifies options for creating a repository
type CreateRepositoryOptions struct {
	Identifier    *string `json:"identifier,omitempty"`
	Description   *string `json:"description,omitempty"`
	IsPublic      *bool   `json:"is_public,omitempty"`
	DefaultBranch *string `json:"default_branch,omitempty"`
	GitIgnore     *string `json:"gitignore,omitempty"`
	License       *string `json:"license,omitempty"`
	Readme        *bool   `json:"readme,omitempty"`
}

// UpdateRepositoryOptions specifies options for updating a repository
type UpdateRepositoryOptions struct {
	Description   *string `json:"description,omitempty"`
	IsPublic      *bool   `json:"is_public,omitempty"`
	DefaultBranch *string `json:"default_branch,omitempty"`
}

// ImportRepositoryOptions specifies options for importing a repository
type ImportRepositoryOptions struct {
	CloneURL   *string `json:"clone_url,omitempty"`
	Username   *string `json:"username,omitempty"`
	Password   *string `json:"password,omitempty"`
	PrivateKey *string `json:"private_key,omitempty"`
	Passphrase *string `json:"passphrase,omitempty"`
	Provider   *string `json:"provider,omitempty"`
	ProviderID *string `json:"provider_id,omitempty"`
}

// ImportRepository imports a repository from external source
func (s *RepositoriesService) ImportRepository(ctx context.Context, spaceRef string, opt *ImportRepositoryOptions) (*Repository, *Response, error) {
	path := fmt.Sprintf("spaces/%s/repos/import", spaceRef)
	var repository Repository
	resp, err := s.client.Post(ctx, path, opt, &repository)
	if err != nil {
		return nil, resp, err
	}
	return &repository, resp, nil
}

// ListRepositoriesOptions specifies options for listing repositories
type ListRepositoriesOptions struct {
	ListOptions
	Recursive *bool `url:"recursive,omitempty"`
}

// GetRepository retrieves a repository by its path
func (s *RepositoriesService) GetRepository(ctx context.Context, repoPath string) (*Repository, *Response, error) {
	path := fmt.Sprintf("repos/%s", repoPath)
	var repository Repository
	resp, err := s.client.Get(ctx, path, &repository)
	if err != nil {
		return nil, resp, err
	}
	return &repository, resp, nil
}

// CreateRepository creates a new repository
func (s *RepositoriesService) CreateRepository(ctx context.Context, spaceRef string, opt *CreateRepositoryOptions) (*Repository, *Response, error) {
	path := fmt.Sprintf("spaces/%s/repos", spaceRef)
	var repository Repository
	resp, err := s.client.Post(ctx, path, opt, &repository)
	if err != nil {
		return nil, resp, err
	}
	return &repository, resp, nil
}

// UpdateRepository updates a repository
func (s *RepositoriesService) UpdateRepository(ctx context.Context, repoPath string, opt *UpdateRepositoryOptions) (*Repository, *Response, error) {
	path := fmt.Sprintf("repos/%s", repoPath)
	var repository Repository
	resp, err := s.client.Patch(ctx, path, opt, &repository)
	if err != nil {
		return nil, resp, err
	}
	return &repository, resp, nil
}

// DeleteRepositoryRequest represents options for deleting a repository
type DeleteRepositoryRequest struct {
	DeleteID *string `json:"delete_id,omitempty"`
}

// DeleteRepository deletes a repository
func (s *RepositoriesService) DeleteRepository(ctx context.Context, repoPath string, deleteID *string) (*Response, error) {
	path := fmt.Sprintf("repos/%s", repoPath)

	var payload *DeleteRepositoryRequest
	if deleteID != nil {
		payload = &DeleteRepositoryRequest{
			DeleteID: deleteID,
		}
	}

	resp, err := s.client.Delete(ctx, path, payload)
	return resp, err
}

// ListBranches lists repository branches
func (s *RepositoriesService) ListBranches(ctx context.Context, repoPath string, opt *ListOptions) ([]*Branch, *Response, error) {
	path := fmt.Sprintf("repos/%s/branches", repoPath)
	var branches []*Branch
	resp, err := s.client.performListRequest(ctx, path, opt, &branches)
	if err != nil {
		return nil, resp, err
	}
	return branches, resp, nil
}

// GetBranch retrieves a specific branch
func (s *RepositoriesService) GetBranch(ctx context.Context, repoPath, branchName string) (*Branch, *Response, error) {
	path := fmt.Sprintf("repos/%s/branches/%s", repoPath, branchName)
	var branch Branch
	resp, err := s.client.Get(ctx, path, &branch)
	if err != nil {
		return nil, resp, err
	}
	return &branch, resp, nil
}

// CreateBranch creates a new branch
func (s *RepositoriesService) CreateBranch(ctx context.Context, repoPath string, opt *CreateBranchOptions) (*Branch, *Response, error) {
	path := fmt.Sprintf("repos/%s/branches", repoPath)
	var branch Branch
	resp, err := s.client.Post(ctx, path, opt, &branch)
	if err != nil {
		return nil, resp, err
	}
	return &branch, resp, nil
}

// CreateBranchOptions specifies options for creating a branch
type CreateBranchOptions struct {
	Name   *string `json:"name,omitempty"`
	Target *string `json:"target,omitempty"`
}

// DeleteBranch deletes a branch
func (s *RepositoriesService) DeleteBranch(ctx context.Context, repoPath, branchName string) (*Response, error) {
	path := fmt.Sprintf("repos/%s/branches/%s", repoPath, branchName)
	resp, err := s.client.Delete(ctx, path, nil)
	return resp, err
}

// Commit represents a git commit
type Commit struct {
	SHA       *string    `json:"sha,omitempty"`
	Message   *string    `json:"message,omitempty"`
	Author    *Signature `json:"author,omitempty"`
	Committer *Signature `json:"committer,omitempty"`
	Added     []string   `json:"added,omitempty"`
	Removed   []string   `json:"removed,omitempty"`
	Modified  []string   `json:"modified,omitempty"`
}

// Signature represents a git signature
type Signature struct {
	Identity *Identity `json:"identity,omitempty"`
	When     *Time     `json:"when,omitempty"`
}

// ListCommits lists commits in a repository
func (s *RepositoriesService) ListCommits(ctx context.Context, repoPath string, opt *ListCommitsOptions) ([]*Commit, *Response, error) {
	path := fmt.Sprintf("repos/%s/commits", repoPath)
	req := s.client.client.R().SetContext(ctx)

	// Add query parameters if options provided
	if opt != nil {
		// Add common query parameters
		buildQueryParams(req, &opt.ListOptions)

		// Add specific query parameters
		if opt.GitRef != nil {
			req.SetQueryParam("git_ref", *opt.GitRef)
		}
		if opt.After != nil {
			req.SetQueryParam("after", *opt.After)
		}
		if opt.Since != nil {
			req.SetQueryParam("since", opt.Since.String())
		}
		if opt.Until != nil {
			req.SetQueryParam("until", opt.Until.String())
		}
		if opt.Path != nil {
			req.SetQueryParam("path", *opt.Path)
		}
	}

	var commits []*Commit
	req.SetSuccessResult(&commits)

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

	return commits, response, nil
}

// ListCommitsOptions specifies options for listing commits
type ListCommitsOptions struct {
	ListOptions
	GitRef *string `url:"git_ref,omitempty"`
	After  *string `url:"after,omitempty"`
	Since  *Time   `url:"since,omitempty"`
	Until  *Time   `url:"until,omitempty"`
	Path   *string `url:"path,omitempty"`
}

// GetCommit retrieves a specific commit
func (s *RepositoriesService) GetCommit(ctx context.Context, repoPath, commitSHA string) (*Commit, *Response, error) {
	path := fmt.Sprintf("repos/%s/commits/%s", repoPath, commitSHA)
	var commit Commit
	resp, err := s.client.Get(ctx, path, &commit)
	if err != nil {
		return nil, resp, err
	}
	return &commit, resp, nil
}

// FileContent represents file content information
type FileContent struct {
	Name    *string `json:"name,omitempty"`
	Path    *string `json:"path,omitempty"`
	SHA     *string `json:"sha,omitempty"`
	Size    *int64  `json:"size,omitempty"`
	Type    *string `json:"type,omitempty"`
	Content *string `json:"content,omitempty"`
}

// GetFileContent retrieves file content
func (s *RepositoriesService) GetFileContent(ctx context.Context, repoPath, filePath string, opt *GetFileOptions) (*FileContent, *Response, error) {
	path := fmt.Sprintf("repos/%s/content/%s", repoPath, filePath)
	var fileContent FileContent
	resp, err := s.client.Get(ctx, path, &fileContent)
	if err != nil {
		return nil, resp, err
	}
	return &fileContent, resp, nil
}

// GetFileOptions specifies options for getting file content
type GetFileOptions struct {
	Ref           *string `url:"git_ref,omitempty"`
	IncludeCommit *bool   `url:"include_commit,omitempty"`
}

// TreeNode represents a tree node in a repository
type TreeNode struct {
	Name *string `json:"name,omitempty"`
	Path *string `json:"path,omitempty"`
	Type *string `json:"type,omitempty"`
	Mode *string `json:"mode,omitempty"`
	SHA  *string `json:"sha,omitempty"`
	Size *int64  `json:"size,omitempty"`
}

// ListPaths lists paths in a repository tree
func (s *RepositoriesService) ListPaths(ctx context.Context, repoPath string, opt *ListPathsOptions) ([]*TreeNode, *Response, error) {
	path := fmt.Sprintf("repos/%s/paths", repoPath)
	req := s.client.client.R().SetContext(ctx)

	// Add specific query parameters
	if opt.GitRef != nil {
		req.SetQueryParam("git_ref", *opt.GitRef)
	}
	if opt.Path != nil {
		req.SetQueryParam("path", *opt.Path)
	}
	if opt.IncludeCommit != nil {
		req.SetQueryParam("include_commit", fmt.Sprintf("%t", *opt.IncludeCommit))
	}

	var nodes []*TreeNode
	req.SetSuccessResult(&nodes)

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

	return nodes, response, nil
}

// ListPathsOptions specifies options for listing paths
type ListPathsOptions struct {
	GitRef        *string `url:"git_ref,omitempty"`
	Path          *string `url:"path,omitempty"`
	IncludeCommit *bool   `url:"include_commit,omitempty"`
}

// Tag represents a git tag
type Tag struct {
	Name        *string    `json:"name,omitempty"`
	SHA         *string    `json:"sha,omitempty"`
	IsAnnotated *bool      `json:"is_annotated,omitempty"`
	Title       *string    `json:"title,omitempty"`
	Message     *string    `json:"message,omitempty"`
	Tagger      *Signature `json:"tagger,omitempty"`
	Commit      *Commit    `json:"commit,omitempty"`
}

// ListTagsOptions specifies options for listing tags
type ListTagsOptions struct {
	ListOptions
	Query         *string `url:"query,omitempty"`
	Sort          *string `url:"sort,omitempty"`
	Order         *string `url:"order,omitempty"`
	IncludeCommit *bool   `url:"include_commit,omitempty"`
}

// ListTags lists repository tags
func (s *RepositoriesService) ListTags(ctx context.Context, repoPath string, opt *ListTagsOptions) ([]*Tag, *Response, error) {
	path := fmt.Sprintf("repos/%s/tags", repoPath)
	req := s.client.client.R().SetContext(ctx)

	if opt != nil {
		buildQueryParams(req, &opt.ListOptions)
		if opt.Query != nil {
			req.SetQueryParam("query", *opt.Query)
		}
		if opt.Sort != nil {
			req.SetQueryParam("sort", *opt.Sort)
		}
		if opt.Order != nil {
			req.SetQueryParam("order", *opt.Order)
		}
		if opt.IncludeCommit != nil {
			req.SetQueryParam("include_commit", fmt.Sprintf("%t", *opt.IncludeCommit))
		}
	}

	var tags []*Tag
	req.SetSuccessResult(&tags)

	resp, err := req.Get(path)
	if err != nil {
		return nil, &Response{Response: resp}, err
	}

	if err := s.client.checkResponse(resp); err != nil {
		return nil, &Response{Response: resp}, err
	}

	response := &Response{Response: resp}
	s.client.parsePaginationHeaders(response)

	return tags, response, nil
}

// CreateTagOptions specifies options for creating a tag
type CreateTagOptions struct {
	Name        *string `json:"name,omitempty"`
	Target      *string `json:"target,omitempty"`
	Message     *string `json:"message,omitempty"`
	BypassRules *bool   `json:"bypass_rules,omitempty"`
	DryRunRules *bool   `json:"dry_run_rules,omitempty"`
}

// CreateTagOutput represents the response from creating a tag
type CreateTagOutput struct {
	Tag
	DryRunRules    *bool            `json:"dry_run_rules,omitempty"`
	RuleViolations []*RuleViolation `json:"rule_violations,omitempty"`
}

// RuleViolation represents a rule violation
type RuleViolation struct {
	Rule       *RuleInfo    `json:"rule,omitempty"`
	Bypassable *bool        `json:"bypassable,omitempty"`
	Bypassed   *bool        `json:"bypassed,omitempty"`
	Violations []*Violation `json:"violations,omitempty"`
}

// RuleInfo represents rule information
type RuleInfo struct {
	ID         *int64  `json:"id,omitempty"`
	Identifier *string `json:"identifier,omitempty"`
	Type       *string `json:"type,omitempty"`
	State      *string `json:"state,omitempty"`
}

// Violation represents a single violation
type Violation struct {
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// CreateTag creates a new tag
func (s *RepositoriesService) CreateTag(ctx context.Context, repoPath string, opt *CreateTagOptions) (*CreateTagOutput, *Response, error) {
	path := fmt.Sprintf("repos/%s/tags", repoPath)
	var output CreateTagOutput
	resp, err := s.client.Post(ctx, path, opt, &output)
	if err != nil {
		return nil, resp, err
	}
	return &output, resp, nil
}

// DeleteTagOutput represents the response from deleting a tag
type DeleteTagOutput struct {
	DryRunRules    *bool            `json:"dry_run_rules,omitempty"`
	RuleViolations []*RuleViolation `json:"rule_violations,omitempty"`
}

// DeleteTag deletes a tag
func (s *RepositoriesService) DeleteTag(ctx context.Context, repoPath, tagName string) (*DeleteTagOutput, *Response, error) {
	path := fmt.Sprintf("repos/%s/tags/%s", repoPath, tagName)
	var output DeleteTagOutput
	resp, err := s.client.DeleteWithResponse(ctx, path, nil, &output)
	if err != nil {
		return nil, resp, err
	}
	return &output, resp, nil
}

// CommitFileAction represents a file action in a commit
type CommitFileAction struct {
	Action   *string `json:"action,omitempty"`
	Path     *string `json:"path,omitempty"`
	Payload  *string `json:"payload,omitempty"`
	SHA      *string `json:"sha,omitempty"`
	Encoding *string `json:"encoding,omitempty"`
}

// CommitFilesOptions specifies options for committing files
type CommitFilesOptions struct {
	Actions     []*CommitFileAction `json:"actions,omitempty"`
	Branch      *string             `json:"branch,omitempty"`
	NewBranch   *string             `json:"new_branch,omitempty"`
	Title       *string             `json:"title,omitempty"`
	Message     *string             `json:"message,omitempty"`
	Author      *Identity           `json:"author,omitempty"`
	BypassRules *bool               `json:"bypass_rules,omitempty"`
	DryRunRules *bool               `json:"dry_run_rules,omitempty"`
}

// FileReference represents a file reference
type FileReference struct {
	Path    *string `json:"path,omitempty"`
	BlobSHA *string `json:"blob_sha,omitempty"`
}

// CommitFilesResponse represents the response from committing files
type CommitFilesResponse struct {
	CommitID       *string          `json:"commit_id,omitempty"`
	ChangedFiles   []*FileReference `json:"changed_files,omitempty"`
	DryRunRules    *bool            `json:"dry_run_rules,omitempty"`
	RuleViolations []*RuleViolation `json:"rule_violations,omitempty"`
}

// CommitFiles commits files to a repository
func (s *RepositoriesService) CommitFiles(ctx context.Context, repoPath string, opt *CommitFilesOptions) (*CommitFilesResponse, *Response, error) {
	path := fmt.Sprintf("repos/%s/commits", repoPath)
	var output CommitFilesResponse
	resp, err := s.client.Post(ctx, path, opt, &output)
	if err != nil {
		return nil, resp, err
	}
	return &output, resp, nil
}

// GetCommitDiffOptions specifies options for getting commit diff
type GetCommitDiffOptions struct {
	IgnoreWhitespace *bool `url:"ignore_whitespace,omitempty"`
}

// GetCommitDiff retrieves the diff for a specific commit
func (s *RepositoriesService) GetCommitDiff(ctx context.Context, repoPath, commitSHA string, opt *GetCommitDiffOptions) (string, *Response, error) {
	path := fmt.Sprintf("repos/%s/commits/%s/diff", repoPath, commitSHA)
	req := s.client.client.R().SetContext(ctx)

	if opt != nil && opt.IgnoreWhitespace != nil {
		req.SetQueryParam("ignore_whitespace", fmt.Sprintf("%t", *opt.IgnoreWhitespace))
	}

	resp, err := req.Get(path)
	if err != nil {
		return "", &Response{Response: resp}, err
	}

	if err := s.client.checkResponse(resp); err != nil {
		return "", &Response{Response: resp}, err
	}

	return resp.String(), &Response{Response: resp}, nil
}

// CommitDivergenceRequest represents a divergence calculation request
type CommitDivergenceRequest struct {
	From *string `json:"from,omitempty"`
	To   *string `json:"to,omitempty"`
}

// CommitDivergence represents commit divergence information
type CommitDivergence struct {
	Ahead  *int `json:"ahead,omitempty"`
	Behind *int `json:"behind,omitempty"`
}

// CalculateCommitDivergenceOptions specifies options for calculating commit divergence
type CalculateCommitDivergenceOptions struct {
	MaxCount *int                       `json:"max_count,omitempty"`
	Requests []*CommitDivergenceRequest `json:"requests,omitempty"`
}

// CalculateCommitDivergence calculates the divergence between commits
func (s *RepositoriesService) CalculateCommitDivergence(ctx context.Context, repoPath string, opt *CalculateCommitDivergenceOptions) ([]*CommitDivergence, *Response, error) {
	path := fmt.Sprintf("repos/%s/commits/calculate-divergence", repoPath)
	var divergences []*CommitDivergence
	resp, err := s.client.Post(ctx, path, opt, &divergences)
	if err != nil {
		return nil, resp, err
	}
	return divergences, resp, nil
}
