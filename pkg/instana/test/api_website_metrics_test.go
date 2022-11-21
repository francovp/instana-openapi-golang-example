/*
Introduction to Instana public APIs

Testing WebsiteMetricsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package instana

import (
	openapiclient "./openapi"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_instana_WebsiteMetricsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test WebsiteMetricsApiService GetBeaconMetrics", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.WebsiteMetricsApi.GetBeaconMetrics(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebsiteMetricsApiService GetBeaconMetricsV2", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.WebsiteMetricsApi.GetBeaconMetricsV2(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WebsiteMetricsApiService GetPageLoad", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string
		var timestamp int64

		resp, httpRes, err := apiClient.WebsiteMetricsApi.GetPageLoad(context.Background(), id, timestamp).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}