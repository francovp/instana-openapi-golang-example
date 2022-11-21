/*
Introduction to Instana public APIs

Testing MaintenanceConfigurationApiService

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

func Test_instana_MaintenanceConfigurationApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MaintenanceConfigurationApiService DeleteMaintenanceConfig", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.MaintenanceConfigurationApi.DeleteMaintenanceConfig(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MaintenanceConfigurationApiService GetMaintenanceConfig", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.MaintenanceConfigurationApi.GetMaintenanceConfig(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MaintenanceConfigurationApiService GetMaintenanceConfigs", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.MaintenanceConfigurationApi.GetMaintenanceConfigs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MaintenanceConfigurationApiService PutMaintenanceConfig", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.MaintenanceConfigurationApi.PutMaintenanceConfig(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
