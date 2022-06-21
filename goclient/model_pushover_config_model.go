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

type PushoverConfigModel struct {
	Expire       *DurationModel         `json:"expire,omitempty"`
	Html         bool                   `json:"html,omitempty"`
	HttpConfig   *HttpClientConfigModel `json:"http_config,omitempty"`
	Message      string                 `json:"message,omitempty"`
	Priority     string                 `json:"priority,omitempty"`
	Retry        *DurationModel         `json:"retry,omitempty"`
	SendResolved bool                   `json:"send_resolved,omitempty"`
	Sound        string                 `json:"sound,omitempty"`
	Title        string                 `json:"title,omitempty"`
	Token        *SecretModel           `json:"token,omitempty"`
	Url          string                 `json:"url,omitempty"`
	UrlTitle     string                 `json:"url_title,omitempty"`
	UserKey      *SecretModel           `json:"user_key,omitempty"`
}