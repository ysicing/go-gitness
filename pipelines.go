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
	ID      *int64  `json:"id,omitempty"`
	Number  *int64  `json:"number,omitempty"`
	Status  *string `json:"status,omitempty"`
	Event   *string `json:"event,omitempty"`
	Title   *string `json:"title,omitempty"`
	Message *string `json:"message,omitempty"`
	Ref     *string `json:"ref,omitempty"`
	Source  *string `json:"source,omitempty"`
	Target  *string `json:"target,omitempty"`
	Author  *string `json:"author,omitempty"`
	Started *Time   `json:"started,omitempty"`
	Created *Time   `json:"created,omitempty"`
	Updated *Time   `json:"updated,omitempty"`
}

// PipelineExecution represents a pipeline execution
type PipelineExecution struct {
	Number   *int64  `json:"number,omitempty"`
	Status   *string `json:"status,omitempty"`
	Event    *string `json:"event,omitempty"`
	Ref      *string `json:"ref,omitempty"`
	Source   *string `json:"source,omitempty"`
	Target   *string `json:"target,omitempty"`
	Author   *string `json:"author,omitempty"`
	Message  *string `json:"message,omitempty"`
	Started  *Time   `json:"started,omitempty"`
	Finished *Time   `json:"finished,omitempty"`
	Created  *Time   `json:"created,omitempty"`
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

// GetPipeline retrieves a specific pipeline
func (s *PipelinesService) GetPipeline(ctx context.Context, repoPath string, pipelineNumber int64) (*Pipeline, *Response, error) {
	path := fmt.Sprintf("repos/%s/pipelines/%d", repoPath, pipelineNumber)
	var pipeline Pipeline
	resp, err := s.client.Get(ctx, path, &pipeline)
	if err != nil {
		return nil, resp, err
	}
	return &pipeline, resp, nil
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
