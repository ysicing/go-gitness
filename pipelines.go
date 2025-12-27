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

// PipelinesService handles communication with pipeline related methods
type PipelinesService struct {
	client *Client
}

// Pipeline represents a Gitness pipeline
type Pipeline struct {
	ID            *int64  `json:"id,omitempty"`
	Identifier    *string `json:"identifier,omitempty"`
	Description   *string `json:"description,omitempty"`
	Disabled      *bool   `json:"disabled,omitempty"`
	ConfigPath    *string `json:"config_path,omitempty"`
	DefaultBranch *string `json:"default_branch,omitempty"`
	RepoID        *int64  `json:"repo_id,omitempty"`
	Seq           *int64  `json:"seq,omitempty"`
	CreatedBy     *int64  `json:"created_by,omitempty"`
	Created       *int64  `json:"created,omitempty"`
	Updated       *int64  `json:"updated,omitempty"`
	Version       *int64  `json:"version,omitempty"`
}

// PipelineExecution represents a pipeline execution
type PipelineExecution struct {
	Number       *int64            `json:"number,omitempty"`
	PipelineID   *int64            `json:"pipeline_id,omitempty"`
	Status       *string           `json:"status,omitempty"`
	Event        *string           `json:"event,omitempty"`
	Action       *string           `json:"action,omitempty"`
	Ref          *string           `json:"ref,omitempty"`
	Source       *string           `json:"source,omitempty"`
	Target       *string           `json:"target,omitempty"`
	Before       *string           `json:"before,omitempty"`
	After        *string           `json:"after,omitempty"`
	AuthorLogin  *string           `json:"author_login,omitempty"`
	AuthorName   *string           `json:"author_name,omitempty"`
	AuthorEmail  *string           `json:"author_email,omitempty"`
	AuthorAvatar *string           `json:"author_avatar,omitempty"`
	Message      *string           `json:"message,omitempty"`
	Error        *string           `json:"error,omitempty"`
	Started      *int64            `json:"started,omitempty"`
	Finished     *int64            `json:"finished,omitempty"`
	Created      *int64            `json:"created,omitempty"`
	Updated      *int64            `json:"updated,omitempty"`
	Params       map[string]string `json:"params,omitempty"`
}

// TriggerAction defines the different actions on triggers will fire
type TriggerAction string

// Trigger action constants
const (
	TriggerActionBranchCreated        TriggerAction = "branch_created"
	TriggerActionBranchUpdated        TriggerAction = "branch_updated"
	TriggerActionTagCreated           TriggerAction = "tag_created"
	TriggerActionTagUpdated           TriggerAction = "tag_updated"
	TriggerActionPullReqCreated       TriggerAction = "pullreq_created"
	TriggerActionPullReqReopened      TriggerAction = "pullreq_reopened"
	TriggerActionPullReqBranchUpdated TriggerAction = "pullreq_branch_updated"
	TriggerActionPullReqClosed        TriggerAction = "pullreq_closed"
	TriggerActionPullReqMerged        TriggerAction = "pullreq_merged"
)

// TriggerEvent defines the different kinds of events in triggers
type TriggerEvent string

// Trigger event constants
const (
	TriggerEventCron        TriggerEvent = "cron"
	TriggerEventManual      TriggerEvent = "manual"
	TriggerEventPush        TriggerEvent = "push"
	TriggerEventPullRequest TriggerEvent = "pull_request"
	TriggerEventTag         TriggerEvent = "tag"
)

// Trigger type constants
const (
	TriggerTypeHook = "@hook"
	TriggerTypeCron = "@cron"
)

// PipelineTrigger represents a pipeline trigger
type PipelineTrigger struct {
	ID          *int64          `json:"id,omitempty"`
	Identifier  *string         `json:"identifier,omitempty"`
	Type        *string         `json:"trigger_type,omitempty"`
	Description *string         `json:"description,omitempty"`
	Disabled    *bool           `json:"disabled,omitempty"`
	Secret      *string         `json:"secret,omitempty"`
	Actions     []TriggerAction `json:"actions,omitempty"`
	Created     *int64          `json:"created,omitempty"`
	Updated     *int64          `json:"updated,omitempty"`
	Version     *int64          `json:"version,omitempty"`
	PipelineID  *int64          `json:"pipeline_id,omitempty"`
	RepoID      *int64          `json:"repo_id,omitempty"`
	CreatedBy   *int64          `json:"created_by,omitempty"`
}

// CreatePipelineTriggerOptions specifies options for creating a pipeline trigger
type CreatePipelineTriggerOptions struct {
	Identifier  *string         `json:"identifier,omitempty"`
	Type        *string         `json:"trigger_type,omitempty"`
	Description *string         `json:"description,omitempty"`
	Disabled    *bool           `json:"disabled,omitempty"`
	Secret      *string         `json:"secret,omitempty"`
	Actions     []TriggerAction `json:"actions,omitempty"`
}

// UpdatePipelineTriggerOptions specifies options for updating a pipeline trigger
type UpdatePipelineTriggerOptions struct {
	Description *string         `json:"description,omitempty"`
	Disabled    *bool           `json:"disabled,omitempty"`
	Secret      *string         `json:"secret,omitempty"`
	Actions     []TriggerAction `json:"actions,omitempty"`
}

// ListPipelineExecutionsOptions specifies options for listing pipeline executions
type ListPipelineExecutionsOptions struct {
	ListOptions
	Status *string `url:"status,omitempty"`
}

// CreatePipelineOptions specifies options for creating a pipeline
type CreatePipelineOptions struct {
	Identifier    *string `json:"identifier,omitempty"`
	Description   *string `json:"description,omitempty"`
	Disabled      *bool   `json:"disabled,omitempty"`
	ConfigPath    *string `json:"config_path,omitempty"`
	DefaultBranch *string `json:"default_branch,omitempty"`
}

// UpdatePipelineOptions specifies options for updating a pipeline
type UpdatePipelineOptions struct {
	Identifier  *string `json:"identifier,omitempty"`
	Description *string `json:"description,omitempty"`
	Disabled    *bool   `json:"disabled,omitempty"`
	ConfigPath  *string `json:"config_path,omitempty"`
}

// LogLine represents a single log line from execution
type LogLine struct {
	Pos  *int    `json:"pos,omitempty"`
	Out  *string `json:"out,omitempty"`
	Time *int64  `json:"time,omitempty"`
}

// ListPipelines lists pipelines for a repository
func (s *PipelinesService) ListPipelines(ctx context.Context, repoPath string, opt *ListOptions) ([]*Pipeline, *Response, error) {
	path := fmt.Sprintf("repos/%s/pipelines", repoPath)
	var pipelines []*Pipeline
	resp, err := s.client.performListRequest(ctx, path, opt, &pipelines)
	if err != nil {
		return nil, resp, err
	}
	return pipelines, resp, nil
}

// CreatePipeline creates a new pipeline
func (s *PipelinesService) CreatePipeline(ctx context.Context, repoPath string, opt *CreatePipelineOptions) (*Pipeline, *Response, error) {
	path := fmt.Sprintf("repos/%s/pipelines", repoPath)
	var pipeline Pipeline
	resp, err := s.client.Post(ctx, path, opt, &pipeline)
	if err != nil {
		return nil, resp, err
	}
	return &pipeline, resp, nil
}

// GetPipeline retrieves a specific pipeline
func (s *PipelinesService) GetPipeline(ctx context.Context, repoPath, pipelineID string) (*Pipeline, *Response, error) {
	path := fmt.Sprintf("repos/%s/pipelines/%s", repoPath, pipelineID)
	var pipeline Pipeline
	resp, err := s.client.Get(ctx, path, &pipeline)
	if err != nil {
		return nil, resp, err
	}
	return &pipeline, resp, nil
}

// UpdatePipeline updates a pipeline
func (s *PipelinesService) UpdatePipeline(ctx context.Context, repoPath, pipelineID string, opt *UpdatePipelineOptions) (*Pipeline, *Response, error) {
	path := fmt.Sprintf("repos/%s/pipelines/%s", repoPath, pipelineID)
	var pipeline Pipeline
	resp, err := s.client.Patch(ctx, path, opt, &pipeline)
	if err != nil {
		return nil, resp, err
	}
	return &pipeline, resp, nil
}

// DeletePipeline deletes a pipeline
func (s *PipelinesService) DeletePipeline(ctx context.Context, repoPath, pipelineID string) (*Response, error) {
	path := fmt.Sprintf("repos/%s/pipelines/%s", repoPath, pipelineID)
	resp, err := s.client.Delete(ctx, path, nil)
	return resp, err
}

// ListPipelineExecutions lists executions for a pipeline
func (s *PipelinesService) ListPipelineExecutions(ctx context.Context, repoPath, pipelineID string, opt *ListPipelineExecutionsOptions) ([]*PipelineExecution, *Response, error) {
	path := fmt.Sprintf("repos/%s/pipelines/%s/executions", repoPath, pipelineID)
	req := s.client.client.R().SetContext(ctx)

	// Add query parameters if options provided
	if opt != nil {
		buildQueryParams(req, &opt.ListOptions)

		if opt.Status != nil {
			req.SetQueryParam("status", *opt.Status)
		}
	}

	var executions []*PipelineExecution
	req.SetSuccessResult(&executions)

	resp, err := req.Get(path)
	if err != nil {
		return nil, &Response{Response: resp}, err
	}

	if err := s.client.checkResponse(resp); err != nil {
		return nil, &Response{Response: resp}, err
	}

	response := &Response{Response: resp}
	s.client.parsePaginationHeaders(response)

	return executions, response, nil
}

// CreateExecution creates/triggers a new pipeline execution
func (s *PipelinesService) CreateExecution(ctx context.Context, repoPath, pipelineID string, branch *string) (*PipelineExecution, *Response, error) {
	path := fmt.Sprintf("repos/%s/pipelines/%s/executions", repoPath, pipelineID)
	req := s.client.client.R().SetContext(ctx)

	if branch != nil {
		req.SetQueryParam("branch", *branch)
	}

	var execution PipelineExecution
	req.SetSuccessResult(&execution)

	resp, err := req.Post(path)
	if err != nil {
		return nil, &Response{Response: resp}, err
	}

	if err := s.client.checkResponse(resp); err != nil {
		return nil, &Response{Response: resp}, err
	}

	return &execution, &Response{Response: resp}, nil
}

// GetPipelineExecution retrieves a specific pipeline execution
func (s *PipelinesService) GetPipelineExecution(ctx context.Context, repoPath, pipelineID string, executionNumber int64) (*PipelineExecution, *Response, error) {
	path := fmt.Sprintf("repos/%s/pipelines/%s/executions/%d", repoPath, pipelineID, executionNumber)
	var execution PipelineExecution
	resp, err := s.client.Get(ctx, path, &execution)
	if err != nil {
		return nil, resp, err
	}
	return &execution, resp, nil
}

// DeleteExecution deletes a pipeline execution
func (s *PipelinesService) DeleteExecution(ctx context.Context, repoPath, pipelineID string, executionNumber int64) (*Response, error) {
	path := fmt.Sprintf("repos/%s/pipelines/%s/executions/%d", repoPath, pipelineID, executionNumber)
	resp, err := s.client.Delete(ctx, path, nil)
	return resp, err
}

// CancelPipelineExecution cancels a pipeline execution
func (s *PipelinesService) CancelPipelineExecution(ctx context.Context, repoPath, pipelineID string, executionNumber int64) (*Response, error) {
	path := fmt.Sprintf("repos/%s/pipelines/%s/executions/%d/cancel", repoPath, pipelineID, executionNumber)
	resp, err := s.client.Post(ctx, path, nil, nil)
	return resp, err
}

// RetryPipelineExecution retries a pipeline execution
func (s *PipelinesService) RetryPipelineExecution(ctx context.Context, repoPath, pipelineID string, executionNumber int64) (*PipelineExecution, *Response, error) {
	path := fmt.Sprintf("repos/%s/pipelines/%s/executions/%d/retry", repoPath, pipelineID, executionNumber)
	var execution PipelineExecution
	resp, err := s.client.Post(ctx, path, nil, &execution)
	if err != nil {
		return nil, resp, err
	}
	return &execution, resp, nil
}

// ListPipelineTriggers lists triggers for a pipeline
func (s *PipelinesService) ListPipelineTriggers(ctx context.Context, repoPath, pipelineID string, opt *ListOptions) ([]*PipelineTrigger, *Response, error) {
	path := fmt.Sprintf("repos/%s/pipelines/%s/triggers", repoPath, pipelineID)
	var triggers []*PipelineTrigger
	resp, err := s.client.performListRequest(ctx, path, opt, &triggers)
	if err != nil {
		return nil, resp, err
	}
	return triggers, resp, nil
}

// CreatePipelineTrigger creates a trigger for a pipeline
func (s *PipelinesService) CreatePipelineTrigger(ctx context.Context, repoPath, pipelineID string, opt *CreatePipelineTriggerOptions) (*PipelineTrigger, *Response, error) {
	path := fmt.Sprintf("repos/%s/pipelines/%s/triggers", repoPath, pipelineID)
	var trigger PipelineTrigger
	resp, err := s.client.Post(ctx, path, opt, &trigger)
	if err != nil {
		return nil, resp, err
	}
	return &trigger, resp, nil
}

// GetPipelineTrigger retrieves a specific pipeline trigger
func (s *PipelinesService) GetPipelineTrigger(ctx context.Context, repoPath, pipelineID, triggerID string) (*PipelineTrigger, *Response, error) {
	path := fmt.Sprintf("repos/%s/pipelines/%s/triggers/%s", repoPath, pipelineID, triggerID)
	var trigger PipelineTrigger
	resp, err := s.client.Get(ctx, path, &trigger)
	if err != nil {
		return nil, resp, err
	}
	return &trigger, resp, nil
}

// UpdatePipelineTrigger updates a pipeline trigger
func (s *PipelinesService) UpdatePipelineTrigger(ctx context.Context, repoPath, pipelineID, triggerID string, opt *UpdatePipelineTriggerOptions) (*PipelineTrigger, *Response, error) {
	path := fmt.Sprintf("repos/%s/pipelines/%s/triggers/%s", repoPath, pipelineID, triggerID)
	var trigger PipelineTrigger
	resp, err := s.client.Patch(ctx, path, opt, &trigger)
	if err != nil {
		return nil, resp, err
	}
	return &trigger, resp, nil
}

// DeletePipelineTrigger deletes a pipeline trigger
func (s *PipelinesService) DeletePipelineTrigger(ctx context.Context, repoPath, pipelineID, triggerID string) (*Response, error) {
	path := fmt.Sprintf("repos/%s/pipelines/%s/triggers/%s", repoPath, pipelineID, triggerID)
	resp, err := s.client.Delete(ctx, path, nil)
	return resp, err
}

// ViewExecutionLogs retrieves logs for a specific step in an execution
func (s *PipelinesService) ViewExecutionLogs(ctx context.Context, repoPath, pipelineID string, executionNumber, stageNumber, stepNumber int64) ([]*LogLine, *Response, error) {
	path := fmt.Sprintf("repos/%s/pipelines/%s/executions/%d/logs/%d/%d", repoPath, pipelineID, executionNumber, stageNumber, stepNumber)
	var logs []*LogLine
	resp, err := s.client.Get(ctx, path, &logs)
	if err != nil {
		return nil, resp, err
	}
	return logs, resp, nil
}
