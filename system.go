// Copyright (c) 2025-2025 All rights reserved.
//
// The original source code is licensed under the Apache License 2.0.
//
// You may review the terms of both licenses in the LICENSE file.

package gitness

import (
	"context"
)

// SystemService handles communication with system related methods
type SystemService struct {
	client *Client
}

// SystemConfig represents system configuration based on SystemConfigOutput schema
type SystemConfig struct {
	ArtifactRegistryEnabled       *bool     `json:"artifact_registry_enabled,omitempty"`
	GitspaceEnabled               *bool     `json:"gitspace_enabled,omitempty"`
	LdapEnabled                   *bool     `json:"ldap_enabled,omitempty"`
	OidcEnabled                   *bool     `json:"oidc_enabled,omitempty"`
	PublicResourceCreationEnabled *bool     `json:"public_resource_creation_enabled,omitempty"`
	SSHEnabled                    *bool     `json:"ssh_enabled,omitempty"`
	UI                            *SystemUI `json:"ui,omitempty"`
	UserSignupAllowed             *bool     `json:"user_signup_allowed,omitempty"`
}

// SystemUI represents UI configuration
type SystemUI struct {
	ShowPlugin *bool `json:"show_plugin,omitempty"`
}

// GetSystemConfig retrieves system configuration
func (s *SystemService) GetSystemConfig(ctx context.Context) (*SystemConfig, *Response, error) {
	var config SystemConfig
	resp, err := s.client.Get(ctx, "system/config", &config)
	if err != nil {
		return nil, resp, err
	}
	return &config, resp, nil
}
