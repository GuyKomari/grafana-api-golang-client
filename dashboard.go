package gapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

// DashboardMeta represents Grafana dashboard meta.
type DashboardMeta struct {
	IsStarred bool   `json:"isStarred"`
	Slug      string `json:"slug"`
	Folder    int64  `json:"folderId"`
	URL       string `json:"url"`
}

// DashboardSaveResponse represents the Grafana API response to creating or saving a dashboard.
type DashboardSaveResponse struct {
	Slug    string `json:"slug"`
	ID      int64  `json:"id"`
	UID     string `json:"uid"`
	Status  string `json:"status"`
	Version int64  `json:"version"`
}

// Dashboard represents a Grafana dashboard.
type Dashboard struct {
	Meta      DashboardMeta          `json:"meta"`
	Model     map[string]interface{} `json:"dashboard"`
	FolderID  int64                  `json:"folderId"`
	FolderUID string                 `json:"folderUid"`
	Overwrite bool                   `json:"overwrite"`

	// This is only used when creating a new dashboard, it will always be empty when getting a dashboard.
	Message string `json:"message"`
}

// SaveDashboard is a deprecated method for saving a Grafana dashboard. Use NewDashboard.
// Deprecated: Use NewDashboard instead.
func (c *Client) SaveDashboard(ctx context.Context, model map[string]interface{}, overwrite bool) (*DashboardSaveResponse, error) {
	wrapper := map[string]interface{}{
		"dashboard": model,
		"overwrite": overwrite,
	}
	data, err := json.Marshal(wrapper)
	if err != nil {
		return nil, err
	}

	result := &DashboardSaveResponse{}
	err = c.request(ctx, "POST", "/api/dashboards/db", nil, bytes.NewBuffer(data), &result)
	if err != nil {
		return nil, err
	}

	return result, err
}

// NewDashboard creates a new Grafana dashboard.
func (c *Client) NewDashboard(ctx context.Context, dashboard Dashboard) (*DashboardSaveResponse, error) {
	data, err := json.Marshal(dashboard)
	if err != nil {
		return nil, err
	}

	result := &DashboardSaveResponse{}
	err = c.request(ctx, "POST", "/api/dashboards/db", nil, bytes.NewBuffer(data), &result)
	if err != nil {
		return nil, err
	}

	return result, err
}

// Dashboards fetches and returns all dashboards.
func (c *Client) Dashboards(ctx context.Context) ([]FolderDashboardSearchResponse, error) {
	params := url.Values{
		"type": {"dash-db"},
	}
	return c.FolderDashboardSearch(ctx, params)
}

// Dashboard will be removed.
// Deprecated: Starting from Grafana v5.0. Use DashboardByUID instead.
func (c *Client) Dashboard(ctx context.Context, slug string) (*Dashboard, error) {
	return c.dashboard(ctx, fmt.Sprintf("/api/dashboards/db/%s", slug))
}

// DashboardByUID gets a dashboard by UID.
func (c *Client) DashboardByUID(ctx context.Context, uid string) (*Dashboard, error) {
	return c.dashboard(ctx, fmt.Sprintf("/api/dashboards/uid/%s", uid))
}

// DashboardsByIDs uses the folder and dashboard search endpoint to find
// dashboards by list of dashboard IDs.
func (c *Client) DashboardsByIDs(ctx context.Context, ids []int64) ([]FolderDashboardSearchResponse, error) {
	dashboardIdsJSON, err := json.Marshal(ids)
	if err != nil {
		return nil, err
	}

	params := url.Values{
		"type":         {"dash-db"},
		"dashboardIds": {string(dashboardIdsJSON)},
	}
	return c.FolderDashboardSearch(ctx, params)
}

func (c *Client) dashboard(ctx context.Context, path string) (*Dashboard, error) {
	result := &Dashboard{}
	err := c.request(ctx, "GET", path, nil, nil, &result)
	if err != nil {
		return nil, err
	}
	result.FolderID = result.Meta.Folder

	return result, err
}

// DeleteDashboard will be removed.
// Deprecated: Starting from Grafana v5.0. Use DeleteDashboardByUID instead.
func (c *Client) DeleteDashboard(ctx context.Context, slug string) error {
	return c.deleteDashboard(ctx, fmt.Sprintf("/api/dashboards/db/%s", slug))
}

// DeleteDashboardByUID deletes a dashboard by UID.
func (c *Client) DeleteDashboardByUID(ctx context.Context, uid string) error {
	return c.deleteDashboard(ctx, fmt.Sprintf("/api/dashboards/uid/%s", uid))
}

func (c *Client) deleteDashboard(ctx context.Context, path string) error {
	return c.request(ctx, "DELETE", path, nil, nil, nil)
}
