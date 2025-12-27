# Gitness Go SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/ysicing/go-gitness.svg)](https://pkg.go.dev/github.com/ysicing/go-gitness)
[![Go Report Card](https://goreportcard.com/badge/github.com/ysicing/go-gitness)](https://goreportcard.com/report/github.com/ysicing/go-gitness)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

**语言版本**: [English](README-EN.md) | [中文](README.md)

一个用于访问 [Gitness](https://gitness.com) API 的全面 Go 客户端库。这个库提供了完整的 Go SDK 来与 Gitness 服务交互，采用了 go-gitlab 的设计模式，并通过分析官方 Gitness OpenAPI 规范进行了增强。

## 特性

- **完整 API 覆盖**：支持所有主要的 Gitness API 端点，包括 Pull Requests、Checks、Templates 等
- **类型安全**：为所有 API 实体提供完整的 Go 类型定义，正确处理空值
- **Context 支持**：内置 context 支持，用于请求取消和超时
- **灵活配置**：可自定义 HTTP 客户端、超时和基础 URL
- **错误处理**：结构化错误响应，包含详细信息
- **模块化设计**：面向服务的架构，代码组织清晰
- **生产就绪**：全面测试和示例用法

## 安装

```bash
go get github.com/ysicing/go-gitness
```

## 快速开始

### 基础用法

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/ysicing/go-gitness"
)

func main() {
    // 创建新的 Gitness 客户端
    client, err := gitness.NewClient("your-api-token")
    if err != nil {
        log.Fatal(err)
    }

    // 获取当前用户
    user, _, err := client.Users.GetCurrentUser(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Hello, %s!\n", *user.DisplayName)
}
```

### 自定义配置

```go
client, err := gitness.NewClient("your-api-token",
    gitness.WithBaseURL("https://your-gitness-instance.com/"),
    gitness.WithTimeout(30*time.Second),
    gitness.WithHTTPClient(customHTTPClient),
)
```

## API 参考

### Pull Request 管理

```go
// 列出 pull request
prs, _, err := client.PullRequests.ListPullRequests(ctx, "my-space/my-repo", &gitness.ListPullRequestsOptions{
    State: gitness.Ptr("open"),
    ListOptions: gitness.ListOptions{
        Page:  gitness.Ptr(1),
        Limit: gitness.Ptr(20),
    },
})

// 创建 pull request
pr, _, err := client.PullRequests.CreatePullRequest(ctx, "my-space/my-repo", &gitness.CreatePullRequestOptions{
    Title:        gitness.Ptr("Fix: Update documentation"),
    Description:  gitness.Ptr("This PR updates the README with latest information"),
    SourceBranch: gitness.Ptr("feature/update-docs"),
    TargetBranch: gitness.Ptr("main"),
    IsDraft:      gitness.Ptr(false),
})

// 合并 pull request
mergedPR, _, err := client.PullRequests.MergePullRequest(ctx, "my-space/my-repo", 123, &gitness.MergePullRequestOptions{
    Method:        gitness.Ptr("merge"),
    CommitMessage: gitness.Ptr("Merge PR: Fix documentation"),
})

// 添加审查者
_, err = client.PullRequests.AddPullRequestReviewer(ctx, "my-space/my-repo", 123, "reviewer-uid")

// 列出 PR 活动/评论
activities, _, err := client.PullRequests.ListPullRequestActivity(ctx, "my-space/my-repo", 123, nil)
```

### 仓库操作

```go
// 创建具有高级选项的仓库
repo, _, err := client.Repositories.CreateRepository(ctx, "my-space", &gitness.CreateRepositoryOptions{
    Identifier:    gitness.Ptr("my-repo"),
    Description:   gitness.Ptr("My awesome repository"),
    IsPublic:      gitness.Ptr(true),
    DefaultBranch: gitness.Ptr("main"),
    GitIgnore:     gitness.Ptr("Go"),
    License:       gitness.Ptr("Apache-2.0"),
    Readme:        gitness.Ptr(true),
})

// 从外部仓库导入
importedRepo, _, err := client.Repositories.ImportRepository(ctx, "my-space", &gitness.ImportRepositoryOptions{
    CloneURL: gitness.Ptr("https://github.com/user/repo.git"),
    Username: gitness.Ptr("your-username"),
    Password: gitness.Ptr("your-token"),
    Provider: gitness.Ptr("github"),
})

// 分支管理
branch, _, err := client.Repositories.CreateBranch(ctx, "my-space/my-repo", &gitness.CreateBranchOptions{
    Name:   gitness.Ptr("feature/new-feature"),
    Target: gitness.Ptr("main"),
})

// 使用过滤器列出提交
commits, _, err := client.Repositories.ListCommits(ctx, "my-space/my-repo", &gitness.ListCommitsOptions{
    GitRef: gitness.Ptr("main"),
    Since:  gitness.Ptr(gitness.Time(time.Now().AddDate(0, 0, -7))), // 最近 7 天
    ListOptions: gitness.ListOptions{
        Limit: gitness.Ptr(50),
    },
})

// 获取文件内容
fileContent, _, err := client.Repositories.GetFileContent(ctx, "my-space/my-repo", "README.md", &gitness.GetFileOptions{
    Ref: gitness.Ptr("main"),
    IncludeCommit: gitness.Ptr(true),
})

// 提交文件到仓库
commitResp, _, err := client.Repositories.CommitFiles(ctx, "my-space/my-repo", &gitness.CommitFilesOptions{
    Branch:  gitness.Ptr("main"),
    Title:   gitness.Ptr("Update README"),
    Message: gitness.Ptr("Add new section"),
    Actions: []*gitness.CommitFileAction{
        {
            Action:  gitness.Ptr("UPDATE"),
            Path:    gitness.Ptr("README.md"),
            Payload: gitness.Ptr("# Updated Content"),
        },
    },
})

// 获取提交差异
diff, _, err := client.Repositories.GetCommitDiff(ctx, "my-space/my-repo", "commit-sha", nil)
```

### 标签管理

```go
// 列出标签
tags, _, err := client.Repositories.ListTags(ctx, "my-space/my-repo", &gitness.ListTagsOptions{
    IncludeCommit: gitness.Ptr(true),
    ListOptions: gitness.ListOptions{
        Limit: gitness.Ptr(20),
    },
})

// 创建标签
tag, _, err := client.Repositories.CreateTag(ctx, "my-space/my-repo", &gitness.CreateTagOptions{
    Name:    gitness.Ptr("v1.0.0"),
    Target:  gitness.Ptr("main"),
    Message: gitness.Ptr("Release v1.0.0"),
})

// 删除标签
_, err = client.Repositories.DeleteTag(ctx, "my-space/my-repo", "v1.0.0")
```

### CI/CD 检查

```go
// 为提交创建检查
check, _, err := client.Checks.CreateCheck(ctx, "my-space/my-repo", "commit-sha", &gitness.CreateCheckOptions{
    Identifier: gitness.Ptr("ci/build"),
    Status:     gitness.Ptr("running"),
    Link:       gitness.Ptr("https://ci.example.com/builds/123"),
    Summary:    gitness.Ptr("Building application..."),
})

// 更新检查状态
updatedCheck, _, err := client.Checks.UpdateCheck(ctx, "my-space/my-repo", "commit-sha", "ci/build", &gitness.UpdateCheckOptions{
    Status:  gitness.Ptr("success"),
    Summary: gitness.Ptr("Build completed successfully"),
})

// 列出提交的所有检查
checks, _, err := client.Checks.ListChecks(ctx, "my-space/my-repo", "commit-sha", &gitness.ListChecksOptions{
    Latest: gitness.Ptr(true),
})
```

### 模板管理

```go
// 创建 pipeline 模板
template, _, err := client.Templates.CreateTemplate(ctx, "my-space", &gitness.CreateTemplateOptions{
    Identifier:  gitness.Ptr("node-ci"),
    Description: gitness.Ptr("Node.js CI pipeline template"),
    Type:        gitness.Ptr("pipeline"),
    Data:        gitness.Ptr(pipelineYAML),
})

// 列出模板
templates, _, err := client.Templates.ListTemplates(ctx, "my-space", nil)

// 获取特定模板
template, _, err := client.Templates.GetTemplate(ctx, "my-space", "node-ci")
```

### Pipeline 管理

```go
// 创建 pipeline
pipeline, _, err := client.Pipelines.CreatePipeline(ctx, "my-space/my-repo", &gitness.CreatePipelineOptions{
    Identifier:    gitness.Ptr("build-and-test"),
    Description:   gitness.Ptr("Build and test pipeline"),
    ConfigPath:    gitness.Ptr(".harness/pipeline.yaml"),
    DefaultBranch: gitness.Ptr("main"),
})

// 列出 pipelines
pipelines, _, err := client.Pipelines.ListPipelines(ctx, "my-space/my-repo", nil)

// 更新 pipeline
updatedPipeline, _, err := client.Pipelines.UpdatePipeline(ctx, "my-space/my-repo", "build-and-test", &gitness.UpdatePipelineOptions{
    Description: gitness.Ptr("Updated description"),
    Disabled:    gitness.Ptr(false),
})

// 触发执行
execution, _, err := client.Pipelines.CreateExecution(ctx, "my-space/my-repo", "build-and-test", gitness.Ptr("main"))

// 查看执行日志
logs, _, err := client.Pipelines.ViewExecutionLogs(ctx, "my-space/my-repo", "build-and-test", 1, 0, 0)
for _, line := range logs {
    fmt.Printf("[%d] %s\n", *line.Pos, *line.Out)
}

// 取消执行
_, err = client.Pipelines.CancelPipelineExecution(ctx, "my-space/my-repo", "build-and-test", 1)

// 重试执行
retryExec, _, err := client.Pipelines.RetryPipelineExecution(ctx, "my-space/my-repo", "build-and-test", 1)

// 删除 pipeline
_, err = client.Pipelines.DeletePipeline(ctx, "my-space/my-repo", "build-and-test")
```

### 高级空间管理

```go
// 创建嵌套空间
space, _, err := client.Spaces.CreateSpace(ctx, &gitness.CreateSpaceOptions{
    Identifier:  gitness.Ptr("team-frontend"),
    ParentRef:   gitness.Ptr("my-organization"),
    Description: gitness.Ptr("Frontend team workspace"),
    IsPublic:    gitness.Ptr(false),
})

// 使用递归列出空间中的仓库
repos, _, err := client.Spaces.ListRepositories(ctx, "my-space", &gitness.ListRepositoriesOptions{
    Recursive: gitness.Ptr(true),
    ListOptions: gitness.ListOptions{
        Sort:  gitness.Ptr("updated"),
        Order: gitness.Ptr("desc"),
    },
})
```

## 完整服务架构

SDK 通过专门的服务模块提供对 Gitness API 的全面覆盖：

### 核心服务
- **Admin**：管理操作和用户管理
- **Audit**：审计日志管理和合规追踪
- **Spaces**：工作空间和组织管理
- **Users**：用户档案和身份验证管理

### 仓库服务
- **Repositories**：Git 仓库管理、分支、提交、标签、文件操作
- **PullRequests**：Pull request 生命周期、审查、合并、评论
- **Checks**：CI/CD 状态检查和构建报告

### DevOps 服务
- **Pipelines**：CI/CD pipeline 完整生命周期管理（创建、更新、删除、执行、日志）
- **Secrets**：密钥和凭证管理
- **Webhooks**：事件通知管理
- **Templates**：可重用的 pipeline 和配置模板

### 基础设施服务
- **Connectors**：外部服务集成（GitHub、GitLab 等）
- **Gitspaces**：开发环境管理
- **InfraProviders**：基础设施提供商配置

## 错误处理

SDK 提供全面的错误处理，包含详细信息：

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

## 分页和过滤

大多数列表操作都支持高级分页和过滤：

```go
options := &gitness.ListPullRequestsOptions{
    State:        gitness.Ptr("open"),
    SourceBranch: gitness.Ptr("feature/*"),
    CreatedBy:    gitness.Ptr(int64(123)),
    ListOptions: gitness.ListOptions{
        Page:  gitness.Ptr(2),
        Limit: gitness.Ptr(50), // 注意：Gitness 使用 'limit' 而不是 'per_page'
        Sort:  gitness.Ptr("created"),
        Order: gitness.Ptr("desc"),
        Query: gitness.Ptr("bug fix"),
    },
}
prs, resp, err := client.PullRequests.ListPullRequests(ctx, "my-space/my-repo", options)

// 从响应头获取分页信息
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

## 分页支持

此 SDK 通过响应头完全支持 Gitness API 分页。Gitness API 在以下头部返回分页信息：

- `x-page`：当前页码
- `x-per-page`：每页项目数
- `x-next-page`：下一页码（如果可用）
- `x-total`：总项目数
- `x-total-pages`：总页数

### 分页示例

```go
// 使用分页列出用户
users, resp, err := client.Admin.ListUsers(ctx, &gitness.ListUsersOptions{
    ListOptions: gitness.ListOptions{
        Page:  gitness.Ptr(1),
        Limit: gitness.Ptr(2), // 注意：使用 'Limit' 而不是 'PerPage'
    },
})

if err != nil {
    log.Fatal(err)
}

fmt.Printf("Page %d of %d\n", *resp.Page, *resp.TotalPages)
fmt.Printf("Showing %d users out of %d total\n", len(users), *resp.Total)

// 检查是否有下一页
if resp.NextPage != nil {
    fmt.Printf("Next page available: %d\n", *resp.NextPage)
}
```

### 遍历所有页面

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
    
    // 处理用户...
    for _, user := range users {
        fmt.Printf("User: %s\n", *user.DisplayName)
    }
    
    // 检查是否有下一页
    if resp.NextPage == nil {
        break // 没有更多页面
    }
    
    page = *resp.NextPage
}
```

## 示例

查看全面的示例：

- **基础用法**：`examples/basic/main.go` - 简单操作
- **高级功能**：`examples/advanced/main.go` - 复杂工作流，包含 Pull Requests、CI/CD 等
- **分页演示**：`examples/pagination/main.go` - 演示使用 Admin Users API 的分页

## 测试

运行完整的测试套件：

```bash
go test ./...
```

运行带覆盖率的测试：

```bash
go test -v -cover ./...
```

构建示例：

```bash
cd examples/basic && go build
cd examples/advanced && go build
```

## 贡献

1. Fork 仓库
2. 创建你的功能分支（`git checkout -b feature/amazing-feature`）
3. 提交你的更改（`git commit -m 'Add some amazing feature'`）
4. 推送到分支（`git push origin feature/amazing-feature`）
5. 打开一个 Pull Request

## 许可证

本项目采用 Apache License 2.0 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 致谢

- 为 [Gitness](https://gitness.com) 开源 DevOps 平台构建
- 通过分析官方 Gitness OpenAPI 规范进行增强