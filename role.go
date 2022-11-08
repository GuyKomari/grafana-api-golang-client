package gapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

type Role struct {
	Version     int64        `json:"version"`
	UID         string       `json:"uid,omitempty"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Global      bool         `json:"global"`
	Group       string       `json:"group"`
	DisplayName string       `json:"displayName"`
	Hidden      bool         `json:"hidden"`
	Permissions []Permission `json:"permissions,omitempty"`
}

type Permission struct {
	Action string `json:"action"`
	Scope  string `json:"scope"`
}

// GetRole gets a role with permissions for the given UID. Available only in Grafana Enterprise 8.+.
func (c *Client) GetRole(ctx context.Context, uid string) (*Role, error) {
	r := &Role{}
	err := c.request(ctx, "GET", buildURL(uid), nil, nil, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// NewRole creates a new role with permissions. Available only in Grafana Enterprise 8.+.
func (c *Client) NewRole(ctx context.Context, role Role) (*Role, error) {
	data, err := json.Marshal(role)
	if err != nil {
		return nil, err
	}

	r := &Role{}

	err = c.request(ctx, "POST", "/api/access-control/roles", nil, bytes.NewBuffer(data), &r)
	if err != nil {
		return nil, err
	}

	return r, err
}

// UpdateRole updates the role and permissions. Available only in Grafana Enterprise 8.+.
func (c *Client) UpdateRole(ctx context.Context, role Role) error {
	data, err := json.Marshal(role)
	if err != nil {
		return err
	}

	err = c.request(ctx, "PUT", buildURL(role.UID), nil, bytes.NewBuffer(data), nil)

	return err
}

// DeleteRole deletes the role with it's permissions. Available only in Grafana Enterprise 8.+.
func (c *Client) DeleteRole(ctx context.Context, uid string, global bool) error {
	qp := map[string][]string{
		"global": {fmt.Sprint(global)},
	}
	return c.request(ctx, "DELETE", buildURL(uid), qp, nil, nil)
}

func buildURL(uid string) string {
	const rootURL = "/api/access-control/roles"
	return fmt.Sprintf("%s/%s", rootURL, uid)
}
