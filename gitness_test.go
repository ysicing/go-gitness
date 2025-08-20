// Copyright (c) 2025-2025 All rights reserved.
//
// The original source code is licensed under the Apache License 2.0.
//
// You may review the terms of both licenses in the LICENSE file.

package gitness

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/imroc/req/v3"
)

func TestNewClientWithReqV3(t *testing.T) {
	client, err := NewClient("test-token")
	if err != nil {
		t.Fatalf("NewClient returned error: %v", err)
	}

	if client.token != "test-token" {
		t.Errorf("Expected token %q, got %q", "test-token", client.token)
	}

	if client.baseURL != defaultBaseURL {
		t.Errorf("Expected baseURL %q, got %q", defaultBaseURL, client.baseURL)
	}

	// Check services are initialized
	services := []interface{}{
		client.Admin, client.Audit, client.Checks, client.Connectors,
		client.Gitspaces, client.InfraProviders, client.Pipelines,
		client.PullRequests, client.Repositories, client.Secrets,
		client.Spaces, client.Templates, client.Users, client.Webhooks,
	}

	for i, service := range services {
		if service == nil {
			t.Errorf("Service %d not initialized", i)
		}
	}
}

func TestNewClientWithOptions(t *testing.T) {
	customBaseURL := "https://custom.gitness.com/"
	customTimeout := 30 * time.Second

	client, err := NewClient("test-token",
		WithBaseURL(customBaseURL),
		WithTimeout(customTimeout),
		WithDebug(),
		WithRetry(3),
	)
	if err != nil {
		t.Fatalf("NewClient with options returned error: %v", err)
	}

	if client.baseURL != customBaseURL {
		t.Errorf("Expected baseURL %q, got %q", customBaseURL, client.baseURL)
	}
}

func TestClientHTTPMethods(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		switch r.Method {
		case "GET":
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"method": "GET"})
		case "POST":
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(map[string]string{"method": "POST"})
		case "PUT":
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"method": "PUT"})
		case "PATCH":
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"method": "PATCH"})
		case "DELETE":
			w.WriteHeader(http.StatusNoContent)
		}
	}))
	defer server.Close()

	client, err := NewClient("test-token", WithBaseURL(server.URL+"/"))
	if err != nil {
		t.Fatalf("NewClient returned error: %v", err)
	}

	ctx := context.Background()

	// Test GET
	var getResult map[string]string
	_, err = client.Get(ctx, "test", &getResult)
	if err != nil {
		t.Errorf("GET request failed: %v", err)
	}
	if getResult["method"] != "GET" {
		t.Errorf("Expected method GET, got %s", getResult["method"])
	}

	// Test POST
	var postResult map[string]string
	_, err = client.Post(ctx, "test", map[string]string{"data": "test"}, &postResult)
	if err != nil {
		t.Errorf("POST request failed: %v", err)
	}
	if postResult["method"] != "POST" {
		t.Errorf("Expected method POST, got %s", postResult["method"])
	}

	// Test PUT
	var putResult map[string]string
	_, err = client.Put(ctx, "test", map[string]string{"data": "test"}, &putResult)
	if err != nil {
		t.Errorf("PUT request failed: %v", err)
	}
	if putResult["method"] != "PUT" {
		t.Errorf("Expected method PUT, got %s", putResult["method"])
	}

	// Test PATCH
	var patchResult map[string]string
	_, err = client.Patch(ctx, "test", map[string]string{"data": "test"}, &patchResult)
	if err != nil {
		t.Errorf("PATCH request failed: %v", err)
	}
	if patchResult["method"] != "PATCH" {
		t.Errorf("Expected method PATCH, got %s", patchResult["method"])
	}

	// Test DELETE
	_, err = client.Delete(ctx, "test", nil)
	if err != nil {
		t.Errorf("DELETE request failed: %v", err)
	}
}

func TestErrorHandling(t *testing.T) {
	// Create a test server that returns errors
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Bad Request",
			"details": "Invalid input provided",
		})
	}))
	defer server.Close()

	client, err := NewClient("test-token", WithBaseURL(server.URL+"/"))
	if err != nil {
		t.Fatalf("NewClient returned error: %v", err)
	}

	ctx := context.Background()

	// Test error handling
	_, err = client.Get(ctx, "test", nil)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	errorResponse, ok := err.(*ErrorResponse)
	if !ok {
		t.Fatalf("Expected ErrorResponse, got %T", err)
	}

	if errorResponse.Message != "Bad Request" {
		t.Errorf("Expected message %q, got %q", "Bad Request", errorResponse.Message)
	}

	if errorResponse.Details != "Invalid input provided" {
		t.Errorf("Expected details %q, got %q", "Invalid input provided", errorResponse.Details)
	}
}

func TestPullRequestOperationsWithReqV3(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		switch r.Method {
		case "GET":
			// List PRs
			json.NewEncoder(w).Encode([]*PullRequest{
				{
					ID:           Ptr(int64(1)),
					Number:       Ptr(int64(1)),
					Title:        Ptr("Test PR"),
					SourceBranch: Ptr("feature"),
					TargetBranch: Ptr("main"),
					State:        Ptr("open"),
				},
			})
		case "POST":
			// Create PR
			json.NewEncoder(w).Encode(&PullRequest{
				ID:           Ptr(int64(2)),
				Number:       Ptr(int64(2)),
				Title:        Ptr("New PR"),
				SourceBranch: Ptr("feature-2"),
				TargetBranch: Ptr("main"),
				State:        Ptr("open"),
			})
		}
	}))
	defer server.Close()

	client, err := NewClient("test-token", WithBaseURL(server.URL+"/"))
	if err != nil {
		t.Fatalf("NewClient returned error: %v", err)
	}

	ctx := context.Background()

	// Test listing pull requests
	prs, _, err := client.PullRequests.ListPullRequests(ctx, "test/repo", nil)
	if err != nil {
		t.Fatalf("ListPullRequests returned error: %v", err)
	}

	if len(prs) != 1 {
		t.Errorf("Expected 1 PR, got %d", len(prs))
	}

	if *prs[0].Title != "Test PR" {
		t.Errorf("Expected PR title 'Test PR', got %q", *prs[0].Title)
	}

	// Test creating a pull request
	pr, _, err := client.PullRequests.CreatePullRequest(ctx, "test/repo", &CreatePullRequestOptions{
		Title:        Ptr("New PR"),
		SourceBranch: Ptr("feature-2"),
		TargetBranch: Ptr("main"),
	})
	if err != nil {
		t.Fatalf("CreatePullRequest returned error: %v", err)
	}

	if *pr.Number != 2 {
		t.Errorf("Expected PR number 2, got %d", *pr.Number)
	}
}

func TestContextCancellation(t *testing.T) {
	// Create a test server with delay
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(200 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "success"}`))
	}))
	defer server.Close()

	client, err := NewClient("test-token", WithBaseURL(server.URL+"/"))
	if err != nil {
		t.Fatalf("NewClient returned error: %v", err)
	}

	// Test with context timeout
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	_, err = client.Get(ctx, "test", nil)
	if err == nil {
		t.Fatal("Expected timeout error, got nil")
	}
}

func TestRetryMechanism(t *testing.T) {
	attemptCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attemptCount++
		if attemptCount < 3 {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"result": "success"})
	}))
	defer server.Close()

	client, err := NewClient("test-token",
		WithBaseURL(server.URL+"/"),
		WithRetry(3),
	)
	if err != nil {
		t.Fatalf("NewClient returned error: %v", err)
	}

	// Configure retry conditions (using correct req/v3 method)
	client.client.SetCommonRetryCondition(func(resp *req.Response, err error) bool {
		return resp.StatusCode >= 500
	})

	ctx := context.Background()
	var result map[string]string
	_, err = client.Get(ctx, "test", &result)
	if err != nil {
		t.Fatalf("Request with retry failed: %v", err)
	}

	if result["result"] != "success" {
		t.Errorf("Expected result 'success', got %s", result["result"])
	}

	if attemptCount != 3 {
		t.Errorf("Expected 3 attempts, got %d", attemptCount)
	}
}

func TestPtr(t *testing.T) {
	str := "test"
	strPtr := Ptr(str)

	if strPtr == nil {
		t.Fatal("Ptr returned nil")
	}

	if *strPtr != str {
		t.Errorf("Expected %q, got %q", str, *strPtr)
	}

	num := 42
	numPtr := Ptr(num)

	if numPtr == nil {
		t.Fatal("Ptr returned nil")
	}

	if *numPtr != num {
		t.Errorf("Expected %d, got %d", num, *numPtr)
	}
}

// TestPaginationHeaders tests the pagination header parsing functionality
func TestPaginationHeaders(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set pagination headers as returned by Gitness API
		w.Header().Set("x-page", "1")
		w.Header().Set("x-per-page", "2")
		w.Header().Set("x-next-page", "2")
		w.Header().Set("x-total", "6")
		w.Header().Set("x-total-pages", "3")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// Return mock user data
		users := []map[string]interface{}{
			{"uid": "user1", "display_name": "User One", "email": "user1@example.com"},
			{"uid": "user2", "display_name": "User Two", "email": "user2@example.com"},
		}
		json.NewEncoder(w).Encode(users)
	}))
	defer server.Close()

	client, err := NewClient("test-token", WithBaseURL(server.URL+"/"))
	if err != nil {
		t.Fatalf("NewClient returned error: %v", err)
	}

	ctx := context.Background()

	// Test Admin ListUsers with pagination
	users, resp, err := client.Admin.ListUsers(ctx, &ListUsersOptions{
		ListOptions: ListOptions{
			Page:  Ptr(1),
			Limit: Ptr(2),
		},
	})

	if err != nil {
		t.Fatalf("ListUsers failed: %v", err)
	}

	if len(users) != 2 {
		t.Errorf("Expected 2 users, got %d", len(users))
	}

	// Verify pagination headers were parsed
	if resp.Page == nil || *resp.Page != 1 {
		t.Errorf("Expected page 1, got %v", resp.Page)
	}

	if resp.PerPage == nil || *resp.PerPage != 2 {
		t.Errorf("Expected per_page 2, got %v", resp.PerPage)
	}

	if resp.NextPage == nil || *resp.NextPage != 2 {
		t.Errorf("Expected next_page 2, got %v", resp.NextPage)
	}

	if resp.Total == nil || *resp.Total != 6 {
		t.Errorf("Expected total 6, got %v", resp.Total)
	}

	if resp.TotalPages == nil || *resp.TotalPages != 3 {
		t.Errorf("Expected total_pages 3, got %v", resp.TotalPages)
	}

	// Verify user data
	if users[0].UID == nil || *users[0].UID != "user1" {
		t.Errorf("Expected first user UID 'user1', got %v", users[0].UID)
	}
}

// TestAllListMethodsPagination tests pagination support across all list methods
func TestAllListMethodsPagination(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set pagination headers
		w.Header().Set("x-page", "1")
		w.Header().Set("x-per-page", "10")
		w.Header().Set("x-total", "50")
		w.Header().Set("x-total-pages", "5")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// Return empty arrays for all list endpoints
		w.Write([]byte("[]"))
	}))
	defer server.Close()

	client, err := NewClient("test-token", WithBaseURL(server.URL+"/"))
	if err != nil {
		t.Fatalf("NewClient returned error: %v", err)
	}

	ctx := context.Background()

	// Test various list methods to ensure they all support pagination
	testCases := []struct {
		name string
		test func() (*Response, error)
	}{
		{
			"Admin.ListUsers",
			func() (*Response, error) {
				_, resp, err := client.Admin.ListUsers(ctx, &ListUsersOptions{
					ListOptions: ListOptions{Page: Ptr(1), Limit: Ptr(10)},
				})
				return resp, err
			},
		},
		{
			"Audit.ListAuditLogs",
			func() (*Response, error) {
				_, resp, err := client.Audit.ListAuditLogs(ctx, &ListAuditLogsOptions{
					ListOptions: ListOptions{Page: Ptr(1), Limit: Ptr(10)},
				})
				return resp, err
			},
		},
		{
			"Spaces.ListSpaces",
			func() (*Response, error) {
				_, resp, err := client.Spaces.ListSpaces(ctx, &ListSpacesOptions{
					ListOptions: ListOptions{Page: Ptr(1), Limit: Ptr(10)},
				})
				return resp, err
			},
		},
		{
			"Repositories.ListBranches",
			func() (*Response, error) {
				_, resp, err := client.Repositories.ListBranches(ctx, "test/repo", &ListOptions{
					Page: Ptr(1), Limit: Ptr(10),
				})
				return resp, err
			},
		},
		{
			"Repositories.ListCommits",
			func() (*Response, error) {
				_, resp, err := client.Repositories.ListCommits(ctx, "test/repo", &ListCommitsOptions{
					ListOptions: ListOptions{Page: Ptr(1), Limit: Ptr(10)},
				})
				return resp, err
			},
		},
		{
			"PullRequests.ListPullRequests",
			func() (*Response, error) {
				_, resp, err := client.PullRequests.ListPullRequests(ctx, "test/repo", &ListPullRequestsOptions{
					ListOptions: ListOptions{Page: Ptr(1), Limit: Ptr(10)},
				})
				return resp, err
			},
		},
		{
			"PullRequests.ListPullRequestActivity",
			func() (*Response, error) {
				_, resp, err := client.PullRequests.ListPullRequestActivity(ctx, "test/repo", 1, &ListOptions{
					Page: Ptr(1), Limit: Ptr(10),
				})
				return resp, err
			},
		},
		{
			"Pipelines.ListPipelines",
			func() (*Response, error) {
				_, resp, err := client.Pipelines.ListPipelines(ctx, "test/repo", &ListOptions{
					Page: Ptr(1), Limit: Ptr(10),
				})
				return resp, err
			},
		},
		{
			"Webhooks.ListWebhooks",
			func() (*Response, error) {
				_, resp, err := client.Webhooks.ListWebhooks(ctx, "test/repo", &ListOptions{
					Page: Ptr(1), Limit: Ptr(10),
				})
				return resp, err
			},
		},
		{
			"Templates.ListTemplates",
			func() (*Response, error) {
				_, resp, err := client.Templates.ListTemplates(ctx, "test/space", &ListOptions{
					Page: Ptr(1), Limit: Ptr(10),
				})
				return resp, err
			},
		},
		{
			"Connectors.ListConnectors",
			func() (*Response, error) {
				_, resp, err := client.Connectors.ListConnectors(ctx, &ListOptions{
					Page: Ptr(1), Limit: Ptr(10),
				})
				return resp, err
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := tc.test()
			if err != nil {
				t.Fatalf("%s failed: %v", tc.name, err)
			}

			// Verify pagination headers were parsed
			if resp.Page == nil || *resp.Page != 1 {
				t.Errorf("%s: Expected page 1, got %v", tc.name, resp.Page)
			}

			if resp.Total == nil || *resp.Total != 50 {
				t.Errorf("%s: Expected total 50, got %v", tc.name, resp.Total)
			}

			if resp.TotalPages == nil || *resp.TotalPages != 5 {
				t.Errorf("%s: Expected total_pages 5, got %v", tc.name, resp.TotalPages)
			}
		})
	}
}
