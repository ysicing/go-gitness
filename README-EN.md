# Gitness Go SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/ysicing/go-gitness.svg)](https://pkg.go.dev/github.com/ysicing/go-gitness)
[![Go Report Card](https://goreportcard.com/badge/github.com/ysicing/go-gitness)](https://goreportcard.com/report/github.com/ysicing/go-gitness)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

**Languages**: [English](README-EN.md) | [中文](README.md)

A comprehensive Go client library for accessing the [Gitness](https://gitness.com) API. This library provides a complete Go SDK for interacting with Gitness services, inspired by the design patterns of go-gitlab and enhanced with analysis of the official Gitness OpenAPI specifications.

## Features

- **Complete API Coverage**: Support for all major Gitness API endpoints including Pull Requests, Checks, Templates, and more
- **Type Safe**: Full Go type definitions for all API entities with proper null handling
- **Context Support**: Built-in context support for request cancellation and timeouts
- **Flexible Configuration**: Customizable HTTP clients, timeouts, and base URLs
- **Error Handling**: Structured error responses with detailed information
- **Modular Design**: Service-oriented architecture for clean code organization
- **Production Ready**: Comprehensive testing and example usage

## Installation

```bash
go get github.com/ysicing/go-gitness
```

## Quick Start

### Basic Usage

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/ysicing/go-gitness"
)

func main() {
    // Create a new Gitness client
    client, err := gitness.NewClient("your-api-token")
    if err != nil {
        log.Fatal(err)
    }

    // Get current user
    user, _, err := client.Users.GetCurrentUser(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Hello, %s!\n", *user.DisplayName)
}
```

### Custom Configuration

```go
client, err := gitness.NewClient("your-api-token",
    gitness.WithBaseURL("https://your-gitness-instance.com/"),
    gitness.WithTimeout(30*time.Second),
    gitness.WithHTTPClient(customHTTPClient),
)
```

## API Reference

### Pull Request Management

```go
// List pull requests
prs, _, err := client.PullRequests.ListPullRequests(ctx, "my-space/my-repo", &gitness.ListPullRequestsOptions{
    State: gitness.Ptr("open"),
    ListOptions: gitness.ListOptions{
        Page:  gitness.Ptr(1),
        Limit: gitness.Ptr(20),
    },
})

// Create a pull request
pr, _, err := client.PullRequests.CreatePullRequest(ctx, "my-space/my-repo", &gitness.CreatePullRequestOptions{
    Title:        gitness.Ptr("Fix: Update documentation"),
    Description:  gitness.Ptr("This PR updates the README with latest information"),
    SourceBranch: gitness.Ptr("feature/update-docs"),
    TargetBranch: gitness.Ptr("main"),
    IsDraft:      gitness.Ptr(false),
})

// Merge a pull request
mergedPR, _, err := client.PullRequests.MergePullRequest(ctx, "my-space/my-repo", 123, &gitness.MergePullRequestOptions{
    Method:        gitness.Ptr("merge"),
    CommitMessage: gitness.Ptr("Merge PR: Fix documentation"),
})

// Add reviewer
_, err = client.PullRequests.AddPullRequestReviewer(ctx, "my-space/my-repo", 123, "reviewer-uid")

// List PR activities/comments
activities, _, err := client.PullRequests.ListPullRequestActivity(ctx, "my-space/my-repo", 123, nil)
```

### Repository Operations

```go
// Create repository with advanced options
repo, _, err := client.Repositories.CreateRepository(ctx, "my-space", &gitness.CreateRepositoryOptions{
    Identifier:    gitness.Ptr("my-repo"),
    Description:   gitness.Ptr("My awesome repository"),
    IsPublic:      gitness.Ptr(true),
    DefaultBranch: gitness.Ptr("main"),
    GitIgnore:     gitness.Ptr("Go"),
    License:       gitness.Ptr("Apache-2.0"),
    Readme:        gitness.Ptr(true),
})

// Import from external repository
importedRepo, _, err := client.Repositories.ImportRepository(ctx, "my-space", &gitness.ImportRepositoryOptions{
    CloneURL: gitness.Ptr("https://github.com/user/repo.git"),
    Username: gitness.Ptr("your-username"),
    Password: gitness.Ptr("your-token"),
    Provider: gitness.Ptr("github"),
})

// Branch management
branch, _, err := client.Repositories.CreateBranch(ctx, "my-space/my-repo", &gitness.CreateBranchOptions{
    Name:   gitness.Ptr("feature/new-feature"),
    Target: gitness.Ptr("main"),
})

// List commits with filtering
commits, _, err := client.Repositories.ListCommits(ctx, "my-space/my-repo", &gitness.ListCommitsOptions{
    GitRef: gitness.Ptr("main"),
    Since:  gitness.Ptr(gitness.Time(time.Now().AddDate(0, 0, -7))), // Last 7 days
    ListOptions: gitness.ListOptions{
        Limit: gitness.Ptr(50),
    },
})

// Get file content
fileContent, _, err := client.Repositories.GetFileContent(ctx, "my-space/my-repo", "README.md", &gitness.GetFileOptions{
    Ref: gitness.Ptr("main"),
    IncludeCommit: gitness.Ptr(true),
})
```

### CI/CD Checks

```go
// Create a check for a commit
check, _, err := client.Checks.CreateCheck(ctx, "my-space/my-repo", "commit-sha", &gitness.CreateCheckOptions{
    Identifier: gitness.Ptr("ci/build"),
    Status:     gitness.Ptr("running"),
    Link:       gitness.Ptr("https://ci.example.com/builds/123"),
    Summary:    gitness.Ptr("Building application..."),
})

// Update check status
updatedCheck, _, err := client.Checks.UpdateCheck(ctx, "my-space/my-repo", "commit-sha", "ci/build", &gitness.UpdateCheckOptions{
    Status:  gitness.Ptr("success"),
    Summary: gitness.Ptr("Build completed successfully"),
})

// List all checks for a commit
checks, _, err := client.Checks.ListChecks(ctx, "my-space/my-repo", "commit-sha", &gitness.ListChecksOptions{
    Latest: gitness.Ptr(true),
})
```

### Template Management

```go
// Create a pipeline template
template, _, err := client.Templates.CreateTemplate(ctx, "my-space", &gitness.CreateTemplateOptions{
    Identifier:  gitness.Ptr("node-ci"),
    Description: gitness.Ptr("Node.js CI pipeline template"),
    Type:        gitness.Ptr("pipeline"),
    Data:        gitness.Ptr(pipelineYAML),
})

// List templates
templates, _, err := client.Templates.ListTemplates(ctx, "my-space", nil)

// Get specific template
template, _, err := client.Templates.GetTemplate(ctx, "my-space", "node-ci")
```

### Advanced Space Management

```go
// Create nested spaces
space, _, err := client.Spaces.CreateSpace(ctx, &gitness.CreateSpaceOptions{
    Identifier:  gitness.Ptr("team-frontend"),
    ParentRef:   gitness.Ptr("my-organization"),
    Description: gitness.Ptr("Frontend team workspace"),
    IsPublic:    gitness.Ptr(false),
})

// List repositories in space with recursion
repos, _, err := client.Spaces.ListRepositories(ctx, "my-space", &gitness.ListRepositoriesOptions{
    Recursive: gitness.Ptr(true),
    ListOptions: gitness.ListOptions{
        Sort:  gitness.Ptr("updated"),
        Order: gitness.Ptr("desc"),
    },
})
```

## Complete Service Architecture

The SDK provides comprehensive coverage of Gitness APIs through specialized service modules:

### Core Services
- **Admin**: Administrative operations and user management
- **Audit**: Audit log management and compliance tracking
- **Spaces**: Workspace and organization management
- **Users**: User profile and authentication management

### Repository Services  
- **Repositories**: Git repository management, branches, commits, file operations
- **PullRequests**: Pull request lifecycle, reviews, merging, comments
- **Checks**: CI/CD status checks and build reporting

### DevOps Services
- **Pipelines**: CI/CD pipeline operations and execution
- **Secrets**: Secret and credential management
- **Webhooks**: Event notification management
- **Templates**: Reusable pipeline and configuration templates

### Infrastructure Services
- **Connectors**: External service integrations (GitHub, GitLab, etc.)
- **Gitspaces**: Development environment management
- **InfraProviders**: Infrastructure provider configuration

## Error Handling

The SDK provides comprehensive error handling with detailed information:

```go
repo, _, err := client.Repositories.GetRepository(ctx, "nonexistent/repo")
if err != nil {
    if gitErr, ok := err.(*gitness.ErrorResponse); ok {
        fmt.Printf("API Error: %s (Status: %d)\n", gitErr.Message, gitErr.Response.StatusCode)
        if gitErr.Details != "" {
            fmt.Printf("Details: %s\n", gitErr.Details)
        }
    } else {
        fmt.Printf("Request Error: %v\n", err)
    }
}
```

## Pagination and Filtering

Most list operations support advanced pagination and filtering:

```go
options := &gitness.ListPullRequestsOptions{
    State:        gitness.Ptr("open"),
    SourceBranch: gitness.Ptr("feature/*"),
    CreatedBy:    gitness.Ptr(int64(123)),
    ListOptions: gitness.ListOptions{
        Page:  gitness.Ptr(2),
        Limit: gitness.Ptr(50), // Note: Gitness uses 'limit' not 'per_page'
        Sort:  gitness.Ptr("created"),
        Order: gitness.Ptr("desc"),
        Query: gitness.Ptr("bug fix"),
    },
}
prs, resp, err := client.PullRequests.ListPullRequests(ctx, "my-space/my-repo", options)

// Access pagination information from response headers
if resp.Total != nil {
    fmt.Printf("Total PRs: %d\n", *resp.Total)
}
if resp.TotalPages != nil {
    fmt.Printf("Total Pages: %d\n", *resp.TotalPages)
}
if resp.NextPage != nil {
    fmt.Printf("Next Page: %d\n", *resp.NextPage)
}
```

## Pagination Support

This SDK fully supports Gitness API pagination through response headers. The Gitness API returns pagination information in the following headers:

- `x-page`: Current page number  
- `x-per-page`: Items per page
- `x-next-page`: Next page number (if available)
- `x-total`: Total number of items
- `x-total-pages`: Total number of pages

### Pagination Example

```go
// List users with pagination
users, resp, err := client.Admin.ListUsers(ctx, &gitness.ListUsersOptions{
    ListOptions: gitness.ListOptions{
        Page:  gitness.Ptr(1),
        Limit: gitness.Ptr(2), // Note: Use 'Limit' not 'PerPage'
    },
})

if err != nil {
    log.Fatal(err)
}

fmt.Printf("Page %d of %d\n", *resp.Page, *resp.TotalPages)
fmt.Printf("Showing %d users out of %d total\n", len(users), *resp.Total)

// Check if there's a next page
if resp.NextPage != nil {
    fmt.Printf("Next page available: %d\n", *resp.NextPage)
}
```

### Walking Through All Pages

```go
page := 1
for {
    users, resp, err := client.Admin.ListUsers(ctx, &gitness.ListUsersOptions{
        ListOptions: gitness.ListOptions{
            Page:  gitness.Ptr(page),
            Limit: gitness.Ptr(10),
        },
    })
    
    if err != nil {
        break
    }
    
    // Process users...
    for _, user := range users {
        fmt.Printf("User: %s\n", *user.DisplayName)
    }
    
    // Check if there's a next page
    if resp.NextPage == nil {
        break // No more pages
    }
    
    page = *resp.NextPage
}
```

## Examples

Check out the comprehensive examples:

- **Basic Usage**: `examples/basic/main.go` - Simple operations  
- **Advanced Features**: `examples/advanced/main.go` - Complex workflows with Pull Requests, CI/CD, and more
- **Pagination Demo**: `examples/pagination/main.go` - Demonstrates pagination with Admin Users API

## Testing

Run the complete test suite:

```bash
go test ./...
```

Run tests with coverage:

```bash
go test -v -cover ./...
```

Build examples:

```bash
cd examples/basic && go build
cd examples/advanced && go build
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built for the [Gitness](https://gitness.com) open-source DevOps platform
- Enhanced through analysis of official Gitness OpenAPI specifications