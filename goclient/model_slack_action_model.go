/*
 * Grafana HTTP API.
 *
 * The Grafana backend exposes an HTTP API, the same API is used by the frontend to do everything from saving dashboards, creating users and updating data sources.
 *
 * API version: 0.0.1
 * Contact: hello@grafana.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package goclient

// See https://api.slack.com/docs/message-attachments#action_fields and https://api.slack.com/docs/message-buttons for more information.
type SlackActionModel struct {
	Confirm *SlackConfirmationFieldModel `json:"confirm,omitempty"`
	Name    string                       `json:"name,omitempty"`
	Style   string                       `json:"style,omitempty"`
	Text    string                       `json:"text,omitempty"`
	Type_   string                       `json:"type,omitempty"`
	Url     string                       `json:"url,omitempty"`
	Value   string                       `json:"value,omitempty"`
}