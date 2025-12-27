// Copyright (c) 2025-2025 All rights reserved.
//
// The original source code is licensed under the Apache License 2.0.
//
// You may review the terms of both licenses in the LICENSE file.

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ysicing/go-gitness"
)

func main() {
	// Get token from environment variable
	token := os.Getenv("GITNESS_TOKEN")
	if token == "" {
		log.Fatal("GITNESS_TOKEN environment variable is required")
	}

	// Create client with advanced req/v3 features
	var client *gitness.Client
	var err error

	// Get base URL from environment (optional)
	baseURL := os.Getenv("GITNESS_BASE_URL")

	if baseURL != "" {
		client, err = gitness.NewClient(token,
			gitness.WithBaseURL(baseURL),
			gitness.WithTimeout(30*time.Second),
			gitness.WithRetry(3),
			// gitness.WithDebug(), // Uncomment for debug logging
		)
	} else {
		client, err = gitness.NewClient(token,
			gitness.WithRetry(2), // Enable retry for reliability
		)
	}

	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	// Example repository path
	repoPath := "my-space/my-repo"

	fmt.Println("=== Gitness Go SDK Advanced Examples ===")

	// 1. Repository Operations
	fmt.Println("\n1. Repository Operations")
	repo, _, err := client.Repositories.GetRepository(ctx, repoPath)
	if err != nil {
		log.Printf("Failed to get repository: %v", err)
	} else {
		fmt.Printf("Repository: %s (%s)\n", *repo.Identifier, *repo.Description)
	}

	// List branches
	fmt.Println("\n2. Branch Management")
	branches, _, err := client.Repositories.ListBranches(ctx, repoPath, nil)
	if err != nil {
		log.Printf("Failed to list branches: %v", err)
	} else {
		fmt.Printf("Found %d branches:\n", len(branches))
		for _, branch := range branches {
			fmt.Printf("- %s (%s)\n", *branch.Name, *branch.SHA)
		}
	}

	// 3. Pull Request Operations
	fmt.Println("\n3. Pull Request Management")

	// List open pull requests
	openPRs, _, err := client.PullRequests.ListPullRequests(ctx, repoPath, &gitness.ListPullRequestsOptions{
		State: gitness.Ptr("open"),
		ListOptions: gitness.ListOptions{
			Page:  gitness.Ptr(1),
			Limit: gitness.Ptr(10),
		},
	})
	if err != nil {
		log.Printf("Failed to list pull requests: %v", err)
	} else {
		fmt.Printf("Found %d open pull requests:\n", len(openPRs))
		for _, pr := range openPRs {
			fmt.Printf("- PR #%d: %s\n", *pr.Number, *pr.Title)
			fmt.Printf("  %s -> %s\n", *pr.SourceBranch, *pr.TargetBranch)
			fmt.Printf("  Status: %s\n", *pr.State)
		}
	}

	// Create a pull request (example)
	/*
		newPR, _, err := client.PullRequests.CreatePullRequest(ctx, repoPath, &gitness.CreatePullRequestOptions{
			Title:        gitness.Ptr("Fix: Update documentation"),
			Description:  gitness.Ptr("This PR updates the README with latest information"),
			SourceBranch: gitness.Ptr("feature/update-docs"),
			TargetBranch: gitness.Ptr("main"),
			IsDraft:      gitness.Ptr(false),
		})
		if err != nil {
			log.Printf("Failed to create pull request: %v", err)
		} else {
			fmt.Printf("Created PR #%d: %s\n", *newPR.Number, *newPR.Title)
		}
	*/

	// 4. Commit Operations
	fmt.Println("\n4. Commit History")
	commits, _, err := client.Repositories.ListCommits(ctx, repoPath, &gitness.ListCommitsOptions{
		ListOptions: gitness.ListOptions{
			Page:  gitness.Ptr(1),
			Limit: gitness.Ptr(5),
		},
	})
	if err != nil {
		log.Printf("Failed to list commits: %v", err)
	} else {
		fmt.Printf("Recent commits:\n")
		for _, commit := range commits {
			fmt.Printf("- %s: %s\n", (*commit.SHA)[:8], *commit.Message)
			if commit.Author != nil && commit.Author.Identity != nil {
				fmt.Printf("  by %s\n", *commit.Author.Identity.Name)
			}
		}
	}

	// 5. Pipeline Operations
	fmt.Println("\n5. CI/CD Pipelines")
	pipelines, _, err := client.Pipelines.ListPipelines(ctx, repoPath, &gitness.ListOptions{
		Page:  gitness.Ptr(1),
		Limit: gitness.Ptr(5),
	})
	if err != nil {
		log.Printf("Failed to list pipelines: %v", err)
	} else {
		fmt.Printf("Recent pipelines:\n")
		for _, pipeline := range pipelines {
			fmt.Printf("- Pipeline %s: %v\n", *pipeline.Identifier, *pipeline.Disabled)
			if pipeline.Description != nil {
				fmt.Printf("  Description: %s\n", *pipeline.Description)
			}
		}
	}

	// 6. Secret Management
	fmt.Println("\n6. Secret Management")
	// List secrets (won't show actual values)
	/*
		secrets, _, err := client.Spaces.ListSecrets(ctx, "my-space", nil)
		if err != nil {
			log.Printf("Failed to list secrets: %v", err)
		} else {
			fmt.Printf("Found %d secrets configured\n", len(secrets))
		}
	*/

	// 7. Webhook Management
	fmt.Println("\n7. Webhook Configuration")
	webhooks, _, err := client.Webhooks.ListWebhooks(ctx, repoPath, nil)
	if err != nil {
		log.Printf("Failed to list webhooks: %v", err)
	} else {
		fmt.Printf("Found %d webhooks:\n", len(webhooks))
		for _, webhook := range webhooks {
			fmt.Printf("- %s: %s\n", *webhook.Identifier, *webhook.URL)
			fmt.Printf("  Enabled: %v\n", *webhook.Enabled)
		}
	}

	// 8. Space Operations
	fmt.Println("\n8. Space Management")
	spaces, _, err := client.Spaces.ListSpaces(ctx, &gitness.ListSpacesOptions{
		ListOptions: gitness.ListOptions{
			Page:  gitness.Ptr(1),
			Limit: gitness.Ptr(5),
		},
	})
	if err != nil {
		log.Printf("Failed to list spaces: %v", err)
	} else {
		fmt.Printf("Available spaces:\n")
		for _, space := range spaces {
			fmt.Printf("- %s: %s\n", *space.Identifier, *space.Description)
		}
	}

	// 9. User Information
	fmt.Println("\n9. User Profile")
	user, _, err := client.Users.GetCurrentUser(ctx)
	if err != nil {
		log.Printf("Failed to get current user: %v", err)
	} else {
		fmt.Printf("Current user: %s (%s)\n", *user.DisplayName, *user.Email)
		fmt.Printf("Admin: %v\n", *user.Admin)
	}

	// 10. Template Operations (if available)
	fmt.Println("\n10. Template Management")
	/*
		templates, _, err := client.Templates.ListTemplates(ctx, "my-space", nil)
		if err != nil {
			log.Printf("Failed to list templates: %v", err)
		} else {
			fmt.Printf("Available templates:\n")
			for _, template := range templates {
				fmt.Printf("- %s: %s\n", *template.Identifier, *template.Description)
			}
		}
	*/

	fmt.Println("\n=== Advanced Examples Complete ===")
}
