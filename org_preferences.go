package gapi

import (
	"bytes"
	"context"
	"encoding/json"
)

// UpdateOrgPreferencesResponse represents the response to a request
// updating Grafana org preferences.
type UpdateOrgPreferencesResponse struct {
	Message string `json:"message"`
}

// OrgPreferences fetches org preferences.
func (c *Client) OrgPreferences(ctx context.Context) (Preferences, error) {
	var prefs Preferences
	err := c.request(ctx, "GET", "/api/org/preferences", nil, nil, &prefs)
	return prefs, err
}

// UpdateOrgPreferences updates only those org preferences specified in the passed Preferences, without impacting others.
func (c *Client) UpdateOrgPreferences(ctx context.Context, p Preferences) (UpdateOrgPreferencesResponse, error) {
	var resp UpdateOrgPreferencesResponse
	data, err := json.Marshal(p)
	if err != nil {
		return resp, err
	}

	err = c.request(ctx, "PATCH", "/api/org/preferences", nil, bytes.NewBuffer(data), &resp)
	if err != nil {
		return resp, err
	}

	return resp, err
}

// UpdateAllOrgPreferences overrwrites all org preferences with the passed Preferences.
func (c *Client) UpdateAllOrgPreferences(ctx context.Context, p Preferences) (UpdateOrgPreferencesResponse, error) {
	var resp UpdateOrgPreferencesResponse
	data, err := json.Marshal(p)
	if err != nil {
		return resp, err
	}

	err = c.request(ctx, "PUT", "/api/org/preferences", nil, bytes.NewBuffer(data), &resp)
	if err != nil {
		return resp, err
	}

	return resp, err
}
