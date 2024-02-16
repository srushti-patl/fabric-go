# {{classname}}

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachConnectionRouteFilter**](RouteFiltersApi.md#AttachConnectionRouteFilter) | **Put** /fabric/v4/connections/{connectionId}/routeFilters/{routeFilterId} | Attach Route Filter
[**CreateRouteFilter**](RouteFiltersApi.md#CreateRouteFilter) | **Post** /fabric/v4/routeFilters | Create Route Filters
[**DeleteRouteFilterByUuid**](RouteFiltersApi.md#DeleteRouteFilterByUuid) | **Delete** /fabric/v4/routeFilters/{routeFilterId} | Delete Route Filter
[**DetachConnectionRouteFilter**](RouteFiltersApi.md#DetachConnectionRouteFilter) | **Delete** /fabric/v4/connections/{connectionId}/routeFilters/{routeFilterId} | Detach Route Filter
[**GetConnectionRouteFilterByUuid**](RouteFiltersApi.md#GetConnectionRouteFilterByUuid) | **Get** /fabric/v4/connections/{connectionId}/routeFilters/{routeFilterId} | Get Route Filter
[**GetConnectionRouteFilters**](RouteFiltersApi.md#GetConnectionRouteFilters) | **Get** /fabric/v4/connections/{connectionId}/routeFilters | Get All RouteFilters
[**GetRouteFilterByUuid**](RouteFiltersApi.md#GetRouteFilterByUuid) | **Get** /fabric/v4/routeFilters/{routeFilterId} | Get Filter By UUID
[**GetRouteFilterChangeByUuid**](RouteFiltersApi.md#GetRouteFilterChangeByUuid) | **Get** /fabric/v4/routeFilters/{routeFilterId}/changes/{changeId} | Get Change By ID
[**GetRouteFilterChanges**](RouteFiltersApi.md#GetRouteFilterChanges) | **Get** /fabric/v4/routeFilters/{routeFilterId}/changes | Get All Changes
[**GetRouteFilterConnections**](RouteFiltersApi.md#GetRouteFilterConnections) | **Get** /fabric/v4/routeFilters/{routeFilterId}/connections | Get Connections
[**PatchRouteFilterByUuid**](RouteFiltersApi.md#PatchRouteFilterByUuid) | **Patch** /fabric/v4/routeFilters/{routeFilterId} | Patch Route Filter
[**SearchRouteFilters**](RouteFiltersApi.md#SearchRouteFilters) | **Post** /fabric/v4/routeFilters/search | Search Route Filters

# **AttachConnectionRouteFilter**
> ConnectionRouteFilterData AttachConnectionRouteFilter(ctx, body, routeFilterId, connectionId)
Attach Route Filter

This API provides capability to attach a Route Filter to a Connection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConnectionRouteFiltersBase**](ConnectionRouteFiltersBase.md)|  | 
  **routeFilterId** | [**string**](.md)| Route Filters Id | 
  **connectionId** | [**string**](.md)| Connection Id | 

### Return type

[**ConnectionRouteFilterData**](ConnectionRouteFilterData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRouteFilter**
> RouteFiltersData CreateRouteFilter(ctx, body)
Create Route Filters

This API provides capability to create a Route Filter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RouteFiltersBase**](RouteFiltersBase.md)|  | 

### Return type

[**RouteFiltersData**](RouteFiltersData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRouteFilterByUuid**
> RouteFiltersData DeleteRouteFilterByUuid(ctx, routeFilterId)
Delete Route Filter

This API provides capability to delete a Route Filter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **routeFilterId** | [**string**](.md)| Route Filters Id | 

### Return type

[**RouteFiltersData**](RouteFiltersData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DetachConnectionRouteFilter**
> ConnectionRouteFilterData DetachConnectionRouteFilter(ctx, routeFilterId, connectionId)
Detach Route Filter

This API provides capability to detach a Route Filter from a Connection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **routeFilterId** | [**string**](.md)| Route Filters Id | 
  **connectionId** | [**string**](.md)| Connection Id | 

### Return type

[**ConnectionRouteFilterData**](ConnectionRouteFilterData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConnectionRouteFilterByUuid**
> ConnectionRouteFilterData GetConnectionRouteFilterByUuid(ctx, routeFilterId, connectionId)
Get Route Filter

This API provides capability to view a specific Route Filter attached to a Connection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **routeFilterId** | [**string**](.md)| Route Filters Id | 
  **connectionId** | [**string**](.md)| Connection Id | 

### Return type

[**ConnectionRouteFilterData**](ConnectionRouteFilterData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConnectionRouteFilters**
> GetAllConnectionRouteFiltersResponse GetConnectionRouteFilters(ctx, connectionId)
Get All RouteFilters

This API provides capability to view all Route Filters attached to a Connection

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **connectionId** | [**string**](.md)| Connection Id | 

### Return type

[**GetAllConnectionRouteFiltersResponse**](GetAllConnectionRouteFiltersResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRouteFilterByUuid**
> RouteFiltersData GetRouteFilterByUuid(ctx, routeFilterId)
Get Filter By UUID

This API provides capability to view a Route Filter by UUID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **routeFilterId** | [**string**](.md)| Route Filters Id | 

### Return type

[**RouteFiltersData**](RouteFiltersData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRouteFilterChangeByUuid**
> RouteFilterChangeData GetRouteFilterChangeByUuid(ctx, routeFilterId, changeId)
Get Change By ID

This API provides capability to retrieve a specific Route Filter's Changes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **routeFilterId** | [**string**](.md)| Route Filters Id | 
  **changeId** | [**string**](.md)| Routing Protocol Change UUID | 

### Return type

[**RouteFilterChangeData**](RouteFilterChangeData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRouteFilterChanges**
> RouteFilterChangeDataResponse GetRouteFilterChanges(ctx, routeFilterId, optional)
Get All Changes

This API provides capability to retrieve all of a Route Filter's Changes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **routeFilterId** | [**string**](.md)| Route Filters Id | 
 **optional** | ***RouteFiltersApiGetRouteFilterChangesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RouteFiltersApiGetRouteFilterChangesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**| offset | 
 **limit** | **optional.Int32**| number of records to fetch | 

### Return type

[**RouteFilterChangeDataResponse**](RouteFilterChangeDataResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRouteFilterConnections**
> GetRouteFilterGetConnectionsResponse GetRouteFilterConnections(ctx, routeFilterId)
Get Connections

This API provides capability to view all Connections using the Route Filter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **routeFilterId** | [**string**](.md)| Route Filters Id | 

### Return type

[**GetRouteFilterGetConnectionsResponse**](GetRouteFilterGetConnectionsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchRouteFilterByUuid**
> RouteFiltersData PatchRouteFilterByUuid(ctx, body, routeFilterId)
Patch Route Filter

This API provides capability to partially update a Route Filter

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]RouteFiltersPatchRequestItem**](RouteFiltersPatchRequestItem.md)|  | 
  **routeFilterId** | [**string**](.md)| Route Filters Id | 

### Return type

[**RouteFiltersData**](RouteFiltersData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchRouteFilters**
> RouteFiltersSearchResponse SearchRouteFilters(ctx, body)
Search Route Filters

This API provides capability to search Route Filters

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RouteFiltersSearchBase**](RouteFiltersSearchBase.md)|  | 

### Return type

[**RouteFiltersSearchResponse**](RouteFiltersSearchResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

