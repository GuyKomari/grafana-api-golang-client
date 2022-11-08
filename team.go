package gapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

// SearchTeam represents a search for a Grafana team.
type SearchTeam struct {
	TotalCount int64   `json:"totalCount,omitempty"`
	Teams      []*Team `json:"teams,omitempty"`
	Page       int64   `json:"page,omitempty"`
	PerPage    int64   `json:"perPage,omitempty"`
}

// Team consists of a get response
// It's used in  Add and Update API
type Team struct {
	ID          int64  `json:"id,omitempty"`
	OrgID       int64  `json:"orgId,omitempty"`
	Name        string `json:"name"`
	Email       string `json:"email,omitempty"`
	AvatarURL   string `json:"avatarUrl,omitempty"`
	MemberCount int64  `json:"memberCount,omitempty"`
	Permission  int64  `json:"permission,omitempty"`
}

// TeamMember represents a Grafana team member.
type TeamMember struct {
	OrgID      int64    `json:"orgId,omitempty"`
	TeamID     int64    `json:"teamId,omitempty"`
	UserID     int64    `json:"userID,omitempty"`
	Email      string   `json:"email,omitempty"`
	Login      string   `json:"login,omitempty"`
	AvatarURL  string   `json:"avatarUrl,omitempty"`
	Permission int64    `json:"permission,omitempty"`
	Labels     []string `json:"labels,omitempty"`
}

// SearchTeam searches Grafana teams and returns the results.
func (c *Client) SearchTeam(ctx context.Context, query string) (*SearchTeam, error) {
	var result SearchTeam

	page := "1"
	perPage := "1000"
	path := "/api/teams/search"
	queryValues := url.Values{}
	queryValues.Set("page", page)
	queryValues.Set("perPage", perPage)
	queryValues.Set("query", query)

	err := c.request(ctx, "GET", path, queryValues, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Team fetches and returns the Grafana team whose ID it's passed.
func (c *Client) Team(ctx context.Context, id int64) (*Team, error) {
	team := &Team{}
	err := c.request(ctx, "GET", fmt.Sprintf("/api/teams/%d", id), nil, nil, team)
	if err != nil {
		return nil, err
	}

	return team, nil
}

// AddTeam makes a new team
// email arg is an optional value.
// If you don't want to set email, please set "" (empty string).
// When team creation is successful, returns the team ID.
func (c *Client) AddTeam(ctx context.Context, name string, email string) (int64, error) {
	id := int64(0)
	path := "/api/teams"
	team := Team{
		Name:  name,
		Email: email,
	}
	data, err := json.Marshal(team)
	if err != nil {
		return id, err
	}

	tmp := struct {
		ID int64 `json:"teamId"`
	}{}

	err = c.request(ctx, "POST", path, nil, bytes.NewBuffer(data), &tmp)
	if err != nil {
		return id, err
	}

	return tmp.ID, err
}

// UpdateTeam updates a Grafana team.
func (c *Client) UpdateTeam(ctx context.Context, id int64, name string, email string) error {
	path := fmt.Sprintf("/api/teams/%d", id)
	team := Team{
		Name: name,
	}
	// add param if email exists
	if email != "" {
		team.Email = email
	}
	data, err := json.Marshal(team)
	if err != nil {
		return err
	}

	return c.request(ctx, "PUT", path, nil, bytes.NewBuffer(data), nil)
}

// DeleteTeam deletes the Grafana team whose ID it's passed.
func (c *Client) DeleteTeam(ctx context.Context, id int64) error {
	return c.request(ctx, "DELETE", fmt.Sprintf("/api/teams/%d", id), nil, nil, nil)
}

// TeamMembers fetches and returns the team members for the Grafana team whose ID it's passed.
func (c *Client) TeamMembers(ctx context.Context, id int64) ([]*TeamMember, error) {
	members := make([]*TeamMember, 0)
	err := c.request(ctx, "GET", fmt.Sprintf("/api/teams/%d/members", id), nil, nil, &members)
	if err != nil {
		return members, err
	}

	return members, nil
}

// AddTeamMember adds a user to the Grafana team whose ID it's passed.
func (c *Client) AddTeamMember(ctx context.Context, id int64, userID int64) error {
	path := fmt.Sprintf("/api/teams/%d/members", id)
	member := TeamMember{UserID: userID}
	data, err := json.Marshal(member)
	if err != nil {
		return err
	}

	return c.request(ctx, "POST", path, nil, bytes.NewBuffer(data), nil)
}

// RemoveMemberFromTeam removes a user from the Grafana team whose ID it's passed.
func (c *Client) RemoveMemberFromTeam(ctx context.Context, id int64, userID int64) error {
	path := fmt.Sprintf("/api/teams/%d/members/%d", id, userID)

	return c.request(ctx, "DELETE", path, nil, nil, nil)
}

// TeamPreferences fetches and returns preferences for the Grafana team whose ID it's passed.
func (c *Client) TeamPreferences(ctx context.Context, id int64) (*Preferences, error) {
	preferences := &Preferences{}
	err := c.request(ctx, "GET", fmt.Sprintf("/api/teams/%d/preferences", id), nil, nil, preferences)
	if err != nil {
		return nil, err
	}

	return preferences, nil
}

// UpdateTeamPreferences updates team preferences for the Grafana team whose ID it's passed.
func (c *Client) UpdateTeamPreferences(ctx context.Context, id int64, preferences Preferences) error {
	path := fmt.Sprintf("/api/teams/%d/preferences", id)
	data, err := json.Marshal(preferences)
	if err != nil {
		return err
	}

	return c.request(ctx, "PUT", path, nil, bytes.NewBuffer(data), nil)
}
