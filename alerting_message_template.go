package gapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

// AlertingMessageTemplate is a re-usable template for Grafana Alerting messages.
type AlertingMessageTemplate struct {
	Name     string `json:"name"`
	Template string `json:"template"`
}

// MessageTemplates fetches all message templates.
func (c *Client) MessageTemplates(ctx context.Context) ([]AlertingMessageTemplate, error) {
	ts := make([]AlertingMessageTemplate, 0)
	err := c.request(ctx, "GET", "/api/v1/provisioning/templates", nil, nil, &ts)
	if err != nil {
		return nil, err
	}
	return ts, nil
}

// MessageTemplate fetches a single message template, identified by its name.
func (c *Client) MessageTemplate(ctx context.Context, name string) (*AlertingMessageTemplate, error) {
	t := AlertingMessageTemplate{}
	uri := fmt.Sprintf("/api/v1/provisioning/templates/%s", name)
	err := c.request(ctx, "GET", uri, nil, nil, &t)
	if err != nil {
		return nil, err
	}
	return &t, err
}

// SetMessageTemplate creates or updates a message template.
func (c *Client) SetMessageTemplate(ctx context.Context, name, content string) error {
	req := struct {
		Template string `json:"template"`
	}{Template: content}
	body, err := json.Marshal(req)
	if err != nil {
		return err
	}

	uri := fmt.Sprintf("/api/v1/provisioning/templates/%s", name)
	return c.request(ctx, "PUT", uri, nil, bytes.NewBuffer(body), nil)
}

// DeleteMessageTemplate deletes a message template.
func (c *Client) DeleteMessageTemplate(ctx context.Context, name string) error {
	uri := fmt.Sprintf("/api/v1/provisioning/templates/%s", name)
	return c.request(ctx, "DELETE", uri, nil, nil, nil)
}
