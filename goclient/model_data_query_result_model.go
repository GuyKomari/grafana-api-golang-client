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

// Deprecated: DataQueryResult should use backend.QueryDataResponse
type DataQueryResultModel struct {
	Dataframes *DataFramesModel          `json:"dataframes,omitempty"`
	Error_     string                    `json:"error,omitempty"`
	Meta       *JsonModel                `json:"meta,omitempty"`
	RefId      string                    `json:"refId,omitempty"`
	Series     *DataTimeSeriesSliceModel `json:"series,omitempty"`
	Tables     []DataTableModel          `json:"tables,omitempty"`
}