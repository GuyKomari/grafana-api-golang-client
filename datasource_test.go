package gapi

import (
	"testing"

	"github.com/gobs/pretty"
	"github.com/grafana/grafana-api-golang-client/goclient/client/datasources"
	"github.com/grafana/grafana-api-golang-client/goclient/models"
	"github.com/stretchr/testify/require"
)

const (
	createdDataSourceJSON = `{"id":1,"uid":"myuid0001","message":"Datasource added", "name": "test_datasource"}`
	getDataSourceJSON     = `{"id":1}`
	getDataSourcesJSON    = `[{"id":1,"name":"foo","type":"cloudwatch","url":"http://some-url.com","access":"access","isDefault":true}]`
)

func TestNewDataSource(t *testing.T) {
	mocksrv, _ := gapiTestTools(t, 200, createdDataSourceJSON)
	defer mocksrv.Close()

	params := datasources.NewAddDatasourceParams().WithBody(
		&models.AddDataSourceCommand{
			Name:      "foo",
			Type:      "cloudwatch",
			URL:       "http://some-url.com",
			Access:    "access",
			IsDefault: true,
			JSONData: map[string]string{
				"assumeRoleArn":           "arn:aws:iam::123:role/some-role",
				"authType":                "keys",
				"customMetricsNamespaces": "SomeNamespace",
				"defaultRegion":           "us-east-1",
				"tlsSkipVerify":           "true",
			},
			SecureJSONData: map[string]string{
				"accessKey": "123",
				"secretKey": "456",
			},
		},
	)

	client, err := GetClient(mocksrv.server.URL)
	require.NoError(t, err)

	res, err := client.Datasources.AddDatasource(params, nil)
	if err != nil {
		t.Fatal(err)
	}

	created := res.Payload.ID

	t.Log(pretty.PrettyFormat(created))

	if *created != int64(1) {
		t.Error("datasource creation response should return the created datasource ID")
	}
}

func TestNewPrometheusDataSource(t *testing.T) {
	mocksrv, _ := gapiTestTools(t, 200, createdDataSourceJSON)
	defer mocksrv.Close()

	params := datasources.NewAddDatasourceParams().WithBody(
		&models.AddDataSourceCommand{
		Name:      "foo_prometheus",
		Type:      "prometheus",
		URL:       "http://some-url.com",
		Access:    "access",
		IsDefault: true,
		JSONData: map[string]string{
			"httpMethod":   "POST",
			"queryTimeout": "60s",
			"timeInterval": "1m",
		},
	})

	client, err := GetClient(mocksrv.server.URL)
	require.NoError(t, err)

	res, err := client.Datasources.AddDatasource(params, nil)
	if err != nil {
		t.Fatal(err)
	}

	created := res.Payload.ID
	t.Log(pretty.PrettyFormat(created))

	if *created != int64(1) {
		t.Error("datasource creation response should return the created datasource ID")
	}
}

func TestNewPrometheusSigV4DataSource(t *testing.T) {
	mocksrv, _ := gapiTestTools(t, 200, createdDataSourceJSON)
	defer mocksrv.Close()

	params := datasources.NewAddDatasourceParams().WithBody(
		&models.AddDataSourceCommand{
			Name:      "sigv4_prometheus",
			Type:      "prometheus",
			URL:       "http://some-url.com",
			Access:    "access",
			IsDefault: true,
			JSONData: map[string]interface{}{
				"httpMethod":    "POST",
				"sigV4Auth":     true,
				"sigV4AuthType": "keys",
				"sigV4Region":   "us-east-1",
			},
			SecureJSONData: map[string]string{
				"sigV4AccessKey": "123",
				"sigV4SecretKey": "456",
			},
		},
	)
	client, err := GetClient(mocksrv.server.URL)
	require.NoError(t, err)

	res, err := client.Datasources.AddDatasource(params, nil)
	if err != nil {
		t.Fatal(err)
	}

	created := res.Payload.ID
	t.Log(pretty.PrettyFormat(created))

	if *created != int64(1) {
		t.Error("datasource creation response should return the created datasource ID")
	}
}

func TestNewElasticsearchDataSource(t *testing.T) {
	mocksrv, _ := gapiTestTools(t, 200, createdDataSourceJSON)
	defer mocksrv.Close()

	params := datasources.NewAddDatasourceParams().WithBody(
		&models.AddDataSourceCommand{
			Name:      "foo_elasticsearch",
			Type:      "elasticsearch",
			URL:       "http://some-url.com",
			IsDefault: true,
			JSONData: map[string]interface{}{
				"esVersion":                  "7.0.0",
				"timeField":                  "time",
				"interval":                   "1m",
				"logMessageField":            "message",
				"logLevelField":              "field",
				"MaxConcurrentShardRequests": 8,
			},
		},
	)

	client, err := GetClient(mocksrv.server.URL)
	require.NoError(t, err)

	res, err := client.Datasources.AddDatasource(params, nil)
	if err != nil {
		t.Fatal(err)
	}

	created := res.Payload.ID
	t.Log(pretty.PrettyFormat(created))

	if *created != int64(1) {
		t.Error("datasource creation response should return the created datasource ID")
	}
}

func TestNewInfluxDBDataSource(t *testing.T) {
	mocksrv, _ := gapiTestTools(t, 200, createdDataSourceJSON)
	defer mocksrv.Close()

	params := datasources.NewAddDatasourceParams().WithBody(
		&models.AddDataSourceCommand{
			Name:      "foo_influxdb",
			Type:      "influxdb",
			URL:       "http://some-url.com",
			IsDefault: true,
			JSONData: map[string]interface{}{
				"defaultBucket":   "telegraf",
				"httpHeaderNames": []string{"Authorization"},
				"organization":    "acme",
				"version":         "Flux",
			},
			SecureJSONData: map[string]string{
				//"httpHeaderValues": []string{"Token alksdjaslkdjkslajdkj.asdlkjaksdjlkajsdlkjsaldj=="},
			},
		},
	)

	client, err := GetClient(mocksrv.server.URL)
	require.NoError(t, err)

	res, err := client.Datasources.AddDatasource(params, nil)
	if err != nil {
		t.Fatal(err)
	}

	created := res.Payload.ID
	t.Log(pretty.PrettyFormat(created))

	if *created != int64(1) {
		t.Error("datasource creation response should return the created datasource ID")
	}
}

func TestNewOpenTSDBDataSource(t *testing.T) {
	mocksrv, _ := gapiTestTools(t, 200, createdDataSourceJSON)
	defer mocksrv.Close()

	params := datasources.NewAddDatasourceParams().WithBody(
		&models.AddDataSourceCommand{
			Name:      "foo_opentsdb",
			Type:      "opentsdb",
			URL:       "http://some-url.com",
			Access:    "access",
			IsDefault: true,
			JSONData: map[string]interface{}{
				"tsdbResolution": 1,
				"tsdbVersion":    3,
			},
		},
	)

	client, err := GetClient(mocksrv.server.URL)
	require.NoError(t, err)
	
	res, err := client.Datasources.AddDatasource(params, nil)
	if err != nil {
		t.Fatal(err)
	}

	created := res.Payload.ID
	t.Log(pretty.PrettyFormat(created))

	if *created != int64(1) {
		t.Error("datasource creation response should return the created datasource ID")
	}
}

func TestNewAzureDataSource(t *testing.T) {
	mocksrv, _ := gapiTestTools(t, 200, createdDataSourceJSON)
	defer mocksrv.Close()

	params := datasources.NewAddDatasourceParams().WithBody(
		&models.AddDataSourceCommand{
			Name:      "foo_azure",
			Type:      "grafana-azure-monitor-datasource",
			URL:       "http://some-url.com",
			Access:    "access",
			IsDefault: true,
			JSONData: map[string]interface{}{
				"azureLogAnalyticsSameAs":      true,
				"clientId":                     "lorem-ipsum",
				"cloudName":                    "azuremonitor",
				"logAnalyticsClientId":         "lorem-ipsum",
				"logAnalyticsDefaultWorkspace": "lorem-ipsum",
				"logAnalyticsTenantId":         "lorem-ipsum",
				"subscriptionId":               "lorem-ipsum",
				"tenantID":                     "lorem-ipsum",
			},
			SecureJSONData: map[string]string{
				"clientSecret": "alksdjaslkdjkslajdkj.asdlkjaksdjlkajsdlkjsaldj==",
			},
		},
	)

	client, err := GetClient(mocksrv.server.URL)
	require.NoError(t, err)
	
	res, err := client.Datasources.AddDatasource(params, nil)
	if err != nil {
		t.Fatal(err)
	}

	created := res.Payload.ID
	t.Log(pretty.PrettyFormat(created))

	if *created != int64(1) {
		t.Error("datasource creation response should return the created datasource ID")
	}
}

func TestDataSources(t *testing.T) {
	server, client := gapiTestTools(t, 200, getDataSourcesJSON)
	defer server.Close()

	datasources, err := client.DataSources()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(datasources))

	if len(datasources) != 1 {
		t.Error("Length of returned datasources should be 1")
	}
	if datasources[0].ID != 1 || datasources[0].Name != "foo" {
		t.Error("Not correctly parsing returned datasources.")
	}
}

func TestDataSourceIDByName(t *testing.T) {
	server, client := gapiTestTools(t, 200, getDataSourceJSON)
	defer server.Close()

	datasourceID, err := client.DataSourceIDByName("foo")
	if err != nil {
		t.Fatal(err)
	}

	if datasourceID != 1 {
		t.Error("Not correctly parsing returned datasources.")
	}
}
