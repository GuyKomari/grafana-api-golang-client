package gapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type RoleAssignments struct {
	RoleUID         string `json:"role_uid"`
	Users           []int  `json:"users,omitempty"`
	Teams           []int  `json:"teams,omitempty"`
	ServiceAccounts []int  `json:"service_accounts,omitempty"`
}

func (c *Client) GetRoleAssignments(ctx context.Context, uid string) (*RoleAssignments, error) {
	assignments := &RoleAssignments{}
	url := fmt.Sprintf("/api/access-control/roles/%s/assignments", uid)
	if err := c.request(ctx, http.MethodGet, url, nil, nil, assignments); err != nil {
		return nil, err
	}

	return assignments, nil
}

func (c *Client) UpdateRoleAssignments(ctx context.Context, ra *RoleAssignments) (*RoleAssignments, error) {
	response := &RoleAssignments{}

	data, err := json.Marshal(ra)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("/api/access-control/roles/%s/assignments", ra.RoleUID)
	err = c.request(ctx, http.MethodPut, url, nil, bytes.NewBuffer(data), &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
