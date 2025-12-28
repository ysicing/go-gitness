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
)

// TestURLEncodingForRepoPath tests that repo paths with slashes are correctly URL encoded
func TestURLEncodingForRepoPath(t *testing.T) {
	// Track the actual request path received by the server
	var receivedPath string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		receivedPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode([]*Branch{})
	}))
	defer server.Close()

	client, err := NewClient("test-token", WithBaseURL(server.URL+"/"))
	if err != nil {
		t.Fatalf("NewClient returned error: %v", err)
	}

	ctx := context.Background()

	// Test with repo path containing a slash
	repoPath := "ci/demo"
	_, _, err = client.Repositories.ListBranches(ctx, repoPath, nil)
	if err != nil {
		t.Fatalf("ListBranches returned error: %v", err)
	}

	// The path should be: /api/v1/repos/ci%2Fdemo/branches
	// NOT: /api/v1/repos/ci/demo/branches (which would be interpreted as 4 path segments)
	expectedPath := "/api/v1/repos/ci%2Fdemo/branches"
	t.Logf("Received path: %s", receivedPath)
	t.Logf("Expected path: %s", expectedPath)

	if receivedPath != expectedPath {
		t.Errorf("Path encoding is incorrect.\nExpected: %s\nReceived: %s", expectedPath, receivedPath)
		t.Errorf("This means the repo path 'ci/demo' is not being properly URL encoded.")
	}
}

// TestURLEncodingForBranchName tests that branch names with special characters are correctly URL encoded
func TestURLEncodingForBranchName(t *testing.T) {
	var receivedPath string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		receivedPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&Branch{})
	}))
	defer server.Close()

	client, err := NewClient("test-token", WithBaseURL(server.URL+"/"))
	if err != nil {
		t.Fatalf("NewClient returned error: %v", err)
	}

	ctx := context.Background()

	// Test with repo path and branch name containing special characters
	repoPath := "ci/demo"
	branchName := "feature/test-branch"
	_, _, err = client.Repositories.GetBranch(ctx, repoPath, branchName)
	if err != nil {
		t.Fatalf("GetBranch returned error: %v", err)
	}

	// The path should be: /api/v1/repos/ci%2Fdemo/branches/feature%2Ftest-branch
	expectedPath := "/api/v1/repos/ci%2Fdemo/branches/feature%2Ftest-branch"
	t.Logf("Received path: %s", receivedPath)
	t.Logf("Expected path: %s", expectedPath)

	if receivedPath != expectedPath {
		t.Errorf("Path encoding is incorrect.\nExpected: %s\nReceived: %s", expectedPath, receivedPath)
	}
}
