/*
Introduction to Instana public APIs

Testing GlobalApplicationAlertConfigurationApiService

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

func Test_instana_GlobalApplicationAlertConfigurationApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test GlobalApplicationAlertConfigurationApiService CreateGlobalApplicationAlertConfig", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.GlobalApplicationAlertConfigurationApi.CreateGlobalApplicationAlertConfig(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GlobalApplicationAlertConfigurationApiService DeleteGlobalApplicationAlertConfig", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.GlobalApplicationAlertConfigurationApi.DeleteGlobalApplicationAlertConfig(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GlobalApplicationAlertConfigurationApiService DisableGlobalApplicationAlertConfig", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.GlobalApplicationAlertConfigurationApi.DisableGlobalApplicationAlertConfig(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GlobalApplicationAlertConfigurationApiService EnableGlobalApplicationAlertConfig", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.GlobalApplicationAlertConfigurationApi.EnableGlobalApplicationAlertConfig(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GlobalApplicationAlertConfigurationApiService FindActiveGlobalApplicationAlertConfigs", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.GlobalApplicationAlertConfigurationApi.FindActiveGlobalApplicationAlertConfigs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GlobalApplicationAlertConfigurationApiService FindGlobalApplicationAlertConfig", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.GlobalApplicationAlertConfigurationApi.FindGlobalApplicationAlertConfig(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GlobalApplicationAlertConfigurationApiService FindGlobalApplicationAlertConfigVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.GlobalApplicationAlertConfigurationApi.FindGlobalApplicationAlertConfigVersions(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GlobalApplicationAlertConfigurationApiService RestoreGlobalApplicationAlertConfig", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string
		var created int64

		resp, httpRes, err := apiClient.GlobalApplicationAlertConfigurationApi.RestoreGlobalApplicationAlertConfig(context.Background(), id, created).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GlobalApplicationAlertConfigurationApiService UpdateGlobalApplicationAlertConfig", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.GlobalApplicationAlertConfigurationApi.UpdateGlobalApplicationAlertConfig(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
