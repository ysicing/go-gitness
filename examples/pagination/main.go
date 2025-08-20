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

	"github.com/ysicing/go-gitness"
)

func main() {
	// Get token from environment variable
	token := os.Getenv("GITNESS_TOKEN")
	if token == "" {
		log.Fatal("GITNESS_TOKEN environment variable is required")
	}

	// Get base URL from environment (default to localhost:3000 for local testing)
	baseURL := os.Getenv("GITNESS_BASE_URL")
	if baseURL == "" {
		baseURL = "http://127.0.0.1:3000/"
	}

	// Create client
	client, err := gitness.NewClient(token, gitness.WithBaseURL(baseURL))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	fmt.Println("=== Gitness SDK Pagination Example ===")
	fmt.Printf("Base URL: %s\n\n", baseURL)

	// Example 1: Admin Users with Pagination
	fmt.Println("1. Admin Users List with Pagination")
	fmt.Println("-----------------------------------")

	page := 1
	limit := 2

	for {
		users, resp, err := client.Admin.ListUsers(ctx, &gitness.ListUsersOptions{
			ListOptions: gitness.ListOptions{
				Page:  gitness.Ptr(page),
				Limit: gitness.Ptr(limit),
			},
		})

		if err != nil {
			log.Printf("Failed to list users on page %d: %v", page, err)
			break
		}

		fmt.Printf("Page %d (limit %d):\n", page, limit)

		// Display pagination info from response headers
		if resp.Total != nil {
			fmt.Printf("  Total users: %d\n", *resp.Total)
		}
		if resp.TotalPages != nil {
			fmt.Printf("  Total pages: %d\n", *resp.TotalPages)
		}
		if resp.Page != nil {
			fmt.Printf("  Current page: %d\n", *resp.Page)
		}
		if resp.PerPage != nil {
			fmt.Printf("  Per page: %d\n", *resp.PerPage)
		}

		// Display users
		fmt.Printf("  Users on this page:\n")
		for i, user := range users {
			fmt.Printf("    %d. %s (%s)", i+1, *user.DisplayName, *user.Email)
			if user.Admin != nil && *user.Admin {
				fmt.Print(" [ADMIN]")
			}
			if user.Blocked != nil && *user.Blocked {
				fmt.Print(" [BLOCKED]")
			}
			fmt.Println()
		}

		// Check if there's a next page
		if resp.NextPage == nil {
			fmt.Println("  No more pages.")
			break
		}

		fmt.Printf("  Next page: %d\n\n", *resp.NextPage)
		page = *resp.NextPage

		// For demo purposes, let's only fetch first 2 pages
		if page > 2 {
			fmt.Println("  Stopping after 2 pages for demo purposes.")
			break
		}
	}

	// Example 2: Filtering Users with Pagination
	fmt.Println("\n2. Filter Admin Users Only")
	fmt.Println("-------------------------")

	adminUsers, resp, err := client.Admin.ListUsers(ctx, &gitness.ListUsersOptions{
		ListOptions: gitness.ListOptions{
			Page:  gitness.Ptr(1),
			Limit: gitness.Ptr(5),
		},
		Admin: gitness.Ptr(true), // Only admin users
	})

	if err != nil {
		log.Printf("Failed to list admin users: %v", err)
	} else {
		fmt.Printf("Found %d admin users:\n", len(adminUsers))
		for i, user := range adminUsers {
			fmt.Printf("  %d. %s (%s)\n", i+1, *user.DisplayName, *user.Email)
		}

		if resp.Total != nil {
			fmt.Printf("Total admin users in system: %d\n", *resp.Total)
		}
	}

	// Example 3: Demonstrate Response Headers
	fmt.Println("\n3. Response Headers Analysis")
	fmt.Println("---------------------------")

	_, resp, err = client.Admin.ListUsers(ctx, &gitness.ListUsersOptions{
		ListOptions: gitness.ListOptions{
			Page:  gitness.Ptr(1),
			Limit: gitness.Ptr(3),
		},
	})

	if err != nil {
		log.Printf("Failed to get response headers: %v", err)
	} else {
		fmt.Println("Gitness API Pagination Headers:")
		fmt.Printf("  x-page: %v\n", resp.Page)
		fmt.Printf("  x-per-page: %v\n", resp.PerPage)
		fmt.Printf("  x-next-page: %v\n", resp.NextPage)
		fmt.Printf("  x-total: %v\n", resp.Total)
		fmt.Printf("  x-total-pages: %v\n", resp.TotalPages)

		// Show the equivalent curl command
		fmt.Printf("\nEquivalent curl command:\n")
		fmt.Printf("curl -X 'GET' \\\n")
		fmt.Printf("  '%sapi/v1/admin/users?page=1&limit=3' \\\n", baseURL)
		fmt.Printf("  -H 'accept: application/json' \\\n")
		fmt.Printf("  -H 'Authorization: Bearer %s'\n", token)
	}

	fmt.Println("\n=== Pagination Example Complete ===")
}
