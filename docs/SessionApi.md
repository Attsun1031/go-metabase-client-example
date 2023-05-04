# \SessionApi

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSession**](SessionApi.md#CreateSession) | **Post** /api/session | Create a new session



## CreateSession

> Session CreateSession(ctx).SessionCreate(sessionCreate).Execute()

Create a new session



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/Attsun1031/go-metabase-client-example"
)

func main() {
    sessionCreate := *openapiclient.NewSessionCreate("Username_example", "Password_example") // SessionCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionApi.CreateSession(context.Background()).SessionCreate(sessionCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionApi.CreateSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSession`: Session
    fmt.Fprintf(os.Stdout, "Response from `SessionApi.CreateSession`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionCreate** | [**SessionCreate**](SessionCreate.md) |  | 

### Return type

[**Session**](Session.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

