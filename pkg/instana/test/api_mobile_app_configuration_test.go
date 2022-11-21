/*
Introduction to Instana public APIs

Testing MobileAppConfigurationApiService

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

func Test_instana_MobileAppConfigurationApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MobileAppConfigurationApiService DeleteMobileAppConfig", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var mobileAppId string

		resp, httpRes, err := apiClient.MobileAppConfigurationApi.DeleteMobileAppConfig(context.Background(), mobileAppId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MobileAppConfigurationApiService GetMobileAppConfig", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.MobileAppConfigurationApi.GetMobileAppConfig(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MobileAppConfigurationApiService GetMobileAppGeoLocationConfiguration", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var mobileAppId string

		resp, httpRes, err := apiClient.MobileAppConfigurationApi.GetMobileAppGeoLocationConfiguration(context.Background(), mobileAppId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MobileAppConfigurationApiService GetMobileAppIpMaskingConfiguration", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var mobileAppId string

		resp, httpRes, err := apiClient.MobileAppConfigurationApi.GetMobileAppIpMaskingConfiguration(context.Background(), mobileAppId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MobileAppConfigurationApiService PostMobileAppConfig", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.MobileAppConfigurationApi.PostMobileAppConfig(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MobileAppConfigurationApiService RenameMobileAppConfig", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var mobileAppId string

		resp, httpRes, err := apiClient.MobileAppConfigurationApi.RenameMobileAppConfig(context.Background(), mobileAppId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MobileAppConfigurationApiService UpdateMobileAppGeoLocationConfiguration", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var mobileAppId string

		resp, httpRes, err := apiClient.MobileAppConfigurationApi.UpdateMobileAppGeoLocationConfiguration(context.Background(), mobileAppId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MobileAppConfigurationApiService UpdateMobileAppIpMaskingConfiguration", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var mobileAppId string

		resp, httpRes, err := apiClient.MobileAppConfigurationApi.UpdateMobileAppIpMaskingConfiguration(context.Background(), mobileAppId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}