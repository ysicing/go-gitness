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

	"github.com/ysicing/go-gitness"
)

func main() {
	// Create a new client
	token := "your_gitness_token_here"
	client, err := gitness.NewClient(token)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	// Example: Using Auth Service (auth.go)
	fmt.Println("=== Auth Service ===")
	loginReq := &gitness.LoginRequest{
		LoginIdentifier: gitness.Ptr("user@example.com"),
		Password:        gitness.Ptr("password"),
	}
	// Note: This is just for demonstration, don't use in real authentication flow
	_, _, err = client.Auth.Login(ctx, loginReq)
	if err != nil {
		fmt.Printf("Login error (expected): %v\n", err)
	}

	// Example: Using Principals Service (principals.go)
	fmt.Println("\n=== Principals Service ===")
	principals, resp, err := client.Principals.ListPrincipals(ctx, &gitness.ListPrincipalsOptions{
		ListOptions: gitness.ListOptions{
			Page:  gitness.Ptr(1),
			Limit: gitness.Ptr(10),
		},
		Type: gitness.Ptr("user"),
	})
	if err != nil {
		fmt.Printf("Error listing principals: %v\n", err)
	} else {
		fmt.Printf("Found %d principals (Page: %v, Total: %v)\n",
			len(principals), resp.Page, resp.Total)
	}

	// Example: Using Plugins Service (plugins.go)
	fmt.Println("\n=== Plugins Service ===")
	plugins, _, err := client.Plugins.ListPlugins(ctx)
	if err != nil {
		fmt.Printf("Error listing plugins: %v\n", err)
	} else {
		fmt.Printf("Found %d plugins\n", len(plugins))
	}

	// Example: Using Resource Service (resources.go)
	fmt.Println("\n=== Resource Service ===")
	gitIgnoreTemplates, _, err := client.Resource.ListGitIgnoreTemplates(ctx)
	if err != nil {
		fmt.Printf("Error listing gitignore templates: %v\n", err)
	} else {
		fmt.Printf("Found %d gitignore templates\n", len(gitIgnoreTemplates))
	}

	licenseTemplates, _, err := client.Resource.ListLicenseTemplates(ctx)
	if err != nil {
		fmt.Printf("Error listing license templates: %v\n", err)
	} else {
		fmt.Printf("Found %d license templates\n", len(licenseTemplates))
	}

	// Example: Using System Service (system.go)
	fmt.Println("\n=== System Service ===")
	systemConfig, _, err := client.System.GetSystemConfig(ctx)
	if err != nil {
		fmt.Printf("Error getting system config: %v\n", err)
	} else {
		fmt.Printf("System config - User Signup Allowed: %v, SSH Enabled: %v\n",
			systemConfig.UserSignupAllowed, systemConfig.SSHEnabled)
	}

	// Example: Using Upload Service (uploads.go)
	fmt.Println("\n=== Upload Service ===")
	upload, _, err := client.Upload.CreateUpload(ctx, "owner/repo", "test-file.txt", 1024)
	if err != nil {
		fmt.Printf("Error creating upload: %v\n", err)
	} else {
		fmt.Printf("Created upload with reference: %v\n", upload.Reference)
	}

	// Example: Using User Service methods (users.go)
	fmt.Println("\n=== User Service ===")
	keys, _, err := client.Users.ListUserKeys(ctx, &gitness.ListPublicKeysOptions{
		ListOptions: gitness.ListOptions{
			Page:  gitness.Ptr(1),
			Limit: gitness.Ptr(10),
		},
		Usage: gitness.Ptr("auth"),
	})
	if err != nil {
		fmt.Printf("Error listing user keys: %v\n", err)
	} else {
		fmt.Printf("Found %d user keys\n", len(keys))
	}

	tokens, _, err := client.Users.ListUserTokens(ctx, &gitness.ListTokensOptions{
		ListOptions: gitness.ListOptions{
			Page:  gitness.Ptr(1),
			Limit: gitness.Ptr(10),
		},
	})
	if err != nil {
		fmt.Printf("Error listing user tokens: %v\n", err)
	} else {
		fmt.Printf("Found %d user tokens\n", len(tokens))
	}

	// Example: Using Pipeline Service methods (pipelines.go)
	fmt.Println("\n=== Pipeline Service ===")
	executions, _, err := client.Pipelines.ListPipelineExecutions(ctx, "owner/repo", "pipeline-id",
		&gitness.ListPipelineExecutionsOptions{
			ListOptions: gitness.ListOptions{
				Page:  gitness.Ptr(1),
				Limit: gitness.Ptr(10),
			},
			Status: gitness.Ptr("success"),
		})
	if err != nil {
		fmt.Printf("Error listing pipeline executions: %v\n", err)
	} else {
		fmt.Printf("Found %d pipeline executions\n", len(executions))
	}

	triggers, _, err := client.Pipelines.ListPipelineTriggers(ctx, "owner/repo", "pipeline-id",
		&gitness.ListOptions{
			Page:  gitness.Ptr(1),
			Limit: gitness.Ptr(10),
		})
	if err != nil {
		fmt.Printf("Error listing pipeline triggers: %v\n", err)
	} else {
		fmt.Printf("Found %d pipeline triggers\n", len(triggers))
	}

	// Example: Using Secret Service methods (services.go - extended)
	fmt.Println("\n=== Secret Service ===")
	repoSecrets, _, err := client.Secrets.ListRepoSecrets(ctx, "owner/repo",
		&gitness.ListOptions{
			Page:  gitness.Ptr(1),
			Limit: gitness.Ptr(10),
		})
	if err != nil {
		fmt.Printf("Error listing repo secrets: %v\n", err)
	} else {
		fmt.Printf("Found %d repo secrets\n", len(repoSecrets))
	}

	globalSecrets, _, err := client.Secrets.ListGlobalSecrets(ctx,
		&gitness.ListOptions{
			Page:  gitness.Ptr(1),
			Limit: gitness.Ptr(10),
		})
	if err != nil {
		fmt.Printf("Error listing global secrets: %v\n", err)
	} else {
		fmt.Printf("Found %d global secrets\n", len(globalSecrets))
	}

	// Example: Using Connectors Service with type-safe config (connectors.go)
	fmt.Println("\n=== Connectors Service (Type-Safe) ===")
	githubConnector := &gitness.CreateConnectorOptions{
		Identifier:  gitness.Ptr("my-github"),
		Description: gitness.Ptr("GitHub connector for integration"),
		Type:        &[]gitness.ConnectorType{gitness.ConnectorTypeGithub}[0],
		SpaceRef:    gitness.Ptr("root"),
		Github: &gitness.GithubConnectorData{
			APIURL:   gitness.Ptr("https://api.github.com"),
			Insecure: gitness.Ptr(false),
			Auth: &gitness.ConnectorAuth{
				AuthType: gitness.ConnectorAuthTypeBearer,
				Token:    gitness.Ptr("github_pat_xxxx"),
			},
		},
	}
	_, _, err = client.Connectors.CreateConnector(ctx, githubConnector)
	if err != nil {
		fmt.Printf("Error creating GitHub connector: %v\n", err)
	}

	connectors, _, err := client.Connectors.ListConnectors(ctx, &gitness.ListOptions{
		Page:  gitness.Ptr(1),
		Limit: gitness.Ptr(10),
	})
	if err != nil {
		fmt.Printf("Error listing connectors: %v\n", err)
	} else {
		fmt.Printf("Found %d connectors\n", len(connectors))
		for _, conn := range connectors {
			if conn.Type != nil {
				fmt.Printf("  - %v: %v (Type: %v, Status: %v)\n",
					conn.Identifier, conn.Description, *conn.Type, conn.LastTestStatus)
			}
		}
	}

	// Example: Using Pipeline Triggers with type-safe actions (pipelines.go)
	fmt.Println("\n=== Pipeline Triggers (Type-Safe) ===")
	triggerOptions := &gitness.CreatePipelineTriggerOptions{
		Identifier:  gitness.Ptr("pr-trigger"),
		Type:        gitness.Ptr(gitness.TriggerTypeHook),
		Description: gitness.Ptr("Trigger on pull request events"),
		Disabled:    gitness.Ptr(false),
		Actions: []gitness.TriggerAction{
			gitness.TriggerActionPullReqCreated,
			gitness.TriggerActionPullReqBranchUpdated,
		},
	}
	_, _, err = client.Pipelines.CreatePipelineTrigger(ctx, "owner/repo", "pipeline-1", triggerOptions)
	if err != nil {
		fmt.Printf("Error creating pipeline trigger: %v\n", err)
	} else {
		fmt.Println("Created pipeline trigger with typed actions")
	}

	// Example: Using System Config with structured data (system.go)
	fmt.Println("\n=== System Config (Structured) ===")
	systemConfig, _, err = client.System.GetSystemConfig(ctx)
	if err != nil {
		fmt.Printf("Error getting system config: %v\n", err)
	} else {
		fmt.Printf("System Configuration:\n")
		fmt.Printf("  - User Signup Allowed: %v\n", systemConfig.UserSignupAllowed)
		fmt.Printf("  - Public Resource Creation Enabled: %v\n", systemConfig.PublicResourceCreationEnabled)
		fmt.Printf("  - SSH Enabled: %v\n", systemConfig.SSHEnabled)
		fmt.Printf("  - OIDC Enabled: %v\n", systemConfig.OidcEnabled)
		fmt.Printf("  - LDAP Enabled: %v\n", systemConfig.LdapEnabled)
		fmt.Printf("  - Gitspace Enabled: %v\n", systemConfig.GitspaceEnabled)
		fmt.Printf("  - Artifact Registry Enabled: %v\n", systemConfig.ArtifactRegistryEnabled)
		if systemConfig.UI != nil {
			fmt.Printf("  - Show Plugin UI: %v\n", systemConfig.UI.ShowPlugin)
		}
	}

	fmt.Println("\n=== Type-Safe Improvements Completed! ===")
	fmt.Printf("SDK now uses structured types instead of map[string]interface{} for:\n")
	fmt.Println("✅ Plugin specs and configuration")
	fmt.Println("✅ Pipeline trigger actions and conditions")
	fmt.Println("✅ System configuration with nested structures")
	fmt.Println("✅ Connector configurations with type-specific data")
	fmt.Println("✅ Pull request activity metadata")
	fmt.Println("✅ Upload and delete request payloads")
	fmt.Println("")
	fmt.Printf("Benefits:\n")
	fmt.Println("- Enhanced type safety and compile-time validation")
	fmt.Println("- Better IDE autocompletion and documentation")
	fmt.Println("- Reduced runtime errors from type assertions")
	fmt.Println("- More maintainable and self-documenting code")
}
