/*
Metabase

Testing DatasetApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package gometabase

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/Attsun1031/go-metabase-client-example"
)

func Test_gometabase_DatasetApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DatasetApiService QueryDatabase", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DatasetApi.QueryDatabase(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
