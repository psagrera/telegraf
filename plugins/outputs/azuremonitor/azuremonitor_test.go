package azuremonitor

import (
	"time"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/metric"
)

// func TestDefaultConnectAndWrite(t *testing.T) {
// 	if testing.Short() {
// 		t.Skip("Skipping integration test in short mode")
// 	}

// 	// Test with all defaults (MSI+IMS)
// 	azmon := &AzureMonitor{}

// 	// Verify that we can connect to Log Analytics
// 	err := azmon.Connect()
// 	require.NoError(t, err)

// 	// Verify that we can write a metric to Log Analytics
// 	err = azmon.Write(testutil.MockMetrics())
// 	require.NoError(t, err)
// }

// MockMetrics returns a mock []telegraf.Metric object for using in unit tests
// of telegraf output sinks.
func getMockMetrics() []telegraf.Metric {
	metrics := make([]telegraf.Metric, 0)
	// Create a new point batch
	metrics = append(metrics, getTestMetric(1.0))
	return metrics
}

// TestMetric Returns a simple test point:
//     measurement -> "test1" or name
//     tags -> "tag1":"value1"
//     value -> value
//     time -> time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
func getTestMetric(value interface{}, name ...string) telegraf.Metric {
	if value == nil {
		panic("Cannot use a nil value")
	}
	measurement := "test1"
	if len(name) > 0 {
		measurement = name[0]
	}
	tags := map[string]string{"tag1": "value1"}
	pt, _ := metric.New(
		measurement,
		tags,
		map[string]interface{}{"value": value},
		time.Now().UTC(),
	)
	return pt
}

// func TestPostData(t *testing.T) {
// 	azmon := &AzureMonitor{
// 		Region: "eastus",
// 	}
// 	err := azmon.Connect()

// 	metrics := getMockMetrics()
// 	t.Logf("mock metrics are %+v\n", metrics)
// 	// metricsList, err := azmon.add(&metrics[0])
// 	for _, m := range metrics {
// 		azmon.Add(m)
// 	}

// 	jsonBytes, err := json.Marshal(azmon.cache)
// 	t.Logf("json content is:\n----------\n%s\n----------\n", string(jsonBytes))

// 	req, err := azmon.postData(&jsonBytes)
// 	if err != nil {
// 		// t.Logf("Error publishing metrics %s", err)
// 		t.Logf("url is %+v\n", req.URL)
// 		// t.Logf("failed request is %+v\n", req)

// 		// raw, err := httputil.DumpRequestOut(req, true)
// 		// if err != nil {
// 		// 	t.Logf("Request detail is \n%s\n", string(raw))
// 		// } else {
// 		// 	t.Logf("could not dump request: %s\n", err)
// 		// }
// 	}
// 	require.NoError(t, err)
// }