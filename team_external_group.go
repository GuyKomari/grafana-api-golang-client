package gapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

// TeamGroup represents a Grafana TeamGroup.
type TeamGroup struct {
	OrgID   int64  `json:"orgId,omitempty"`
	TeamID  int64  `json:"teamId,omitempty"`
	GroupID string `json:"groupID,omitempty"`
}

// TeamGroups fetches and returns the list of Grafana team group whose Team ID it's passed.
func (c *Client) TeamGroups(ctx context.Context, id int64) ([]TeamGroup, error) {
	teamGroups := make([]TeamGroup, 0)
	err := c.request(ctx, "GET", fmt.Sprintf("/api/teams/%d/groups", id), nil, nil, &teamGroups)
	if err != nil {
		return teamGroups, err
	}

	return teamGroups, nil
}

// NewTeamGroup creates a new Grafana Team Group .
func (c *Client) NewTeamGroup(ctx context.Context, id int64, groupID string) error {
	dataMap := map[string]string{
		"groupId": groupID,
	}
	data, err := json.Marshal(dataMap)
	if err != nil {
		return err
	}

	return c.request(ctx, "POST", fmt.Sprintf("/api/teams/%d/groups", id), nil, bytes.NewBuffer(data), nil)
}

// DeleteTeam deletes the Grafana team whose ID it's passed.
func (c *Client) DeleteTeamGroup(ctx context.Context, id int64, groupID string) error {
	return c.request(ctx, "DELETE", fmt.Sprintf("/api/teams/%d/groups/%s", id, groupID), nil, nil, nil)
}
