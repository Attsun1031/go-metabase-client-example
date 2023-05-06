# Go API client for gometabase

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import gometabase "github.com/Attsun1031/go-metabase-client-example"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), gometabase.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), gometabase.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), gometabase.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), gometabase.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost:3000*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DatabaseApi* | [**CreateDatabase**](docs/DatabaseApi.md#createdatabase) | **Post** /api/database | 
*DatabaseApi* | [**ListDatabases**](docs/DatabaseApi.md#listdatabases) | **Get** /api/database | List Databases
*DatasetApi* | [**QueryDatabase**](docs/DatasetApi.md#querydatabase) | **Post** /api/dataset | Execute a query
*GroupApi* | [**CreateGroup**](docs/GroupApi.md#creategroup) | **Post** /api/permissions/group | 
*GroupApi* | [**ListGroups**](docs/GroupApi.md#listgroups) | **Get** /api/permissions/group | 
*SessionApi* | [**CreateSession**](docs/SessionApi.md#createsession) | **Post** /api/session | Create a new session


## Documentation For Models

 - [Database](docs/Database.md)
 - [DatabaseCreate](docs/DatabaseCreate.md)
 - [DatabaseCreateDetails](docs/DatabaseCreateDetails.md)
 - [DatabaseCreateSchedules](docs/DatabaseCreateSchedules.md)
 - [DatabaseCreateSchedulesMetadataSync](docs/DatabaseCreateSchedulesMetadataSync.md)
 - [DatabaseDetails](docs/DatabaseDetails.md)
 - [DatabaseTable](docs/DatabaseTable.md)
 - [DatasetQueryConstraints](docs/DatasetQueryConstraints.md)
 - [DatasetQueryDsl](docs/DatasetQueryDsl.md)
 - [DatasetQueryDslPage](docs/DatasetQueryDslPage.md)
 - [DatasetQueryJsonQuery](docs/DatasetQueryJsonQuery.md)
 - [DatasetQueryNative](docs/DatasetQueryNative.md)
 - [DatasetQueryOpts](docs/DatasetQueryOpts.md)
 - [DatasetQueryResults](docs/DatasetQueryResults.md)
 - [DatasetQueryResultsCol](docs/DatasetQueryResultsCol.md)
 - [DatasetQueryResultsColFingerprint](docs/DatasetQueryResultsColFingerprint.md)
 - [DatasetQueryResultsColFingerprintGlobal](docs/DatasetQueryResultsColFingerprintGlobal.md)
 - [DatasetQueryResultsColFingerprintType](docs/DatasetQueryResultsColFingerprintType.md)
 - [DatasetQueryResultsColTarget](docs/DatasetQueryResultsColTarget.md)
 - [DatasetQueryResultsData](docs/DatasetQueryResultsData.md)
 - [DatasetQueryResultsMetadata](docs/DatasetQueryResultsMetadata.md)
 - [DatasetQueryResultsMetadataColumn](docs/DatasetQueryResultsMetadataColumn.md)
 - [DatasetQueryResultsNativeForm](docs/DatasetQueryResultsNativeForm.md)
 - [Group](docs/Group.md)
 - [GroupCreate](docs/GroupCreate.md)
 - [ListDatabases200Response](docs/ListDatabases200Response.md)
 - [Session](docs/Session.md)
 - [SessionCreate](docs/SessionCreate.md)


## Documentation For Authorization



### ApiKeyAuth

- **Type**: API key
- **API key parameter name**: X-Metabase-Session
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: X-Metabase-Session and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



