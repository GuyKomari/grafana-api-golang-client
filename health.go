package gapi

import "context"

type HealthResponse struct {
	Commit   string `json:"commit,omitempty"`
	Database string `json:"database,omitempty"`
	Version  string `json:"version,omitempty"`
}

func (c *Client) Health(ctx context.Context) (HealthResponse, error) {
	health := HealthResponse{}
	err := c.request(ctx, "GET", "/api/health", nil, nil, &health)
	if err != nil {
		return health, err
	}
	return health, nil
}
