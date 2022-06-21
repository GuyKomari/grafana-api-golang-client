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

import (
	"time"
)

type SyncResultModel struct {
	Elapsed        *DurationModel    `json:"Elapsed,omitempty"`
	FailedUsers    []FailedUserModel `json:"FailedUsers,omitempty"`
	MissingUserIds []int64           `json:"MissingUserIds,omitempty"`
	Started        time.Time         `json:"Started,omitempty"`
	UpdatedUserIds []int64           `json:"UpdatedUserIds,omitempty"`
}