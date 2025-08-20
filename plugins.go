// Copyright (c) 2025-2025 All rights reserved.
//
// The original source code is licensed under the Apache License 2.0.
//
// You may review the terms of both licenses in the LICENSE file.

package gitness

import (
	"context"
)

// PluginsService handles communication with plugins related methods
type PluginsService struct {
	client *Client
}

// Plugin represents a Gitness plugin
type Plugin struct {
	ID          *string `json:"id,omitempty"`
	Identifier  *string `json:"identifier,omitempty"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Type        *string `json:"type,omitempty"`
	Version     *string `json:"version,omitempty"`
	Logo        *string `json:"logo,omitempty"`
	Enabled     *bool   `json:"enabled,omitempty"`
	// Spec is a YAML template to be used for the plugin
	Spec *string `json:"spec,omitempty"`
}

// ListPlugins lists all plugins
func (s *PluginsService) ListPlugins(ctx context.Context) ([]*Plugin, *Response, error) {
	var plugins []*Plugin
	resp, err := s.client.Get(ctx, "plugins", &plugins)
	if err != nil {
		return nil, resp, err
	}
	return plugins, resp, nil
}
