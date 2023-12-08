# {{classname}}

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRouteFilterRule**](RouteFilterRulesApi.md#CreateRouteFilterRule) | **Post** /fabric/v4/routeFilters/{routeFilterId}/routeFilterRules | Create RFRule
[**CreateRouteFilterRulesInBulk**](RouteFilterRulesApi.md#CreateRouteFilterRulesInBulk) | **Post** /fabric/v4/routeFilters/{routeFilterId}/routeFilterRules/bulk | Bulk RFRules
[**DeleteRouteFilterRuleByUuid**](RouteFilterRulesApi.md#DeleteRouteFilterRuleByUuid) | **Delete** /fabric/v4/routeFilters/{routeFilterId}/routeFilterRules/{routeFilterRuleId} | DeleteRFRule
[**GetRouteFilterRuleByUuid**](RouteFilterRulesApi.md#GetRouteFilterRuleByUuid) | **Get** /fabric/v4/routeFilters/{routeFilterId}/routeFilterRules/{routeFilterRuleId} | GetRFRule By UUID
[**GetRouteFilterRuleChangeByUuid**](RouteFilterRulesApi.md#GetRouteFilterRuleChangeByUuid) | **Get** /fabric/v4/routeFilters/{routeFilterId}/routeFilterRules/{routeFilterRuleId}/changes/{changeId} | Get Change By ID
[**GetRouteFilterRuleChanges**](RouteFilterRulesApi.md#GetRouteFilterRuleChanges) | **Get** /fabric/v4/routeFilters/{routeFilterId}/routeFilterRules/{routeFilterRuleId}/changes | Get All Changes
[**GetRouteFilterRules**](RouteFilterRulesApi.md#GetRouteFilterRules) | **Get** /fabric/v4/routeFilters/{routeFilterId}/routeFilterRules | GetRFRules
[**PatchRouteFilterRuleByUuid**](RouteFilterRulesApi.md#PatchRouteFilterRuleByUuid) | **Patch** /fabric/v4/routeFilters/{routeFilterId}/routeFilterRules/{routeFilterRuleId} | PatchRFilterRule
[**ReplaceRouteFilterRuleByUuid**](RouteFilterRulesApi.md#ReplaceRouteFilterRuleByUuid) | **Put** /fabric/v4/routeFilters/{routeFilterId}/routeFilterRules/{routeFilterRuleId} | ReplaceRFRule

# **CreateRouteFilterRule**
> RouteFilterRulesData CreateRouteFilterRule(ctx, body, routeFilterId)
Create RFRule

This API provides capability to create a Route Filter Rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RouteFilterRulesBase**](RouteFilterRulesBase.md)|  | 
  **routeFilterId** | [**string**](.md)| Route Filters Id | 

### Return type

[**RouteFilterRulesData**](RouteFilterRulesData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRouteFilterRulesInBulk**
> GetRouteFilterRulesResponse CreateRouteFilterRulesInBulk(ctx, body, routeFilterId)
Bulk RFRules

This API provides capability to create bulk route filter rules

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RouteFilterRulesPostRequest**](RouteFilterRulesPostRequest.md)|  | 
  **routeFilterId** | [**string**](.md)| Route Filters Id | 

### Return type

[**GetRouteFilterRulesResponse**](GetRouteFilterRulesResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRouteFilterRuleByUuid**
> RouteFilterRulesData DeleteRouteFilterRuleByUuid(ctx, routeFilterId, routeFilterRuleId)
DeleteRFRule

This API provides capability to delete a Route Filter Rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **routeFilterId** | [**string**](.md)| Route Filters Id | 
  **routeFilterRuleId** | [**string**](.md)| Route  Filter  Rules Id | 

### Return type

[**RouteFilterRulesData**](RouteFilterRulesData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRouteFilterRuleByUuid**
> RouteFilterRulesData GetRouteFilterRuleByUuid(ctx, routeFilterId, routeFilterRuleId)
GetRFRule By UUID

This API provides capability to view a Route Filter Rule by UUID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **routeFilterId** | [**string**](.md)| Route Filters Id | 
  **routeFilterRuleId** | [**string**](.md)| Route  Filter  Rules Id | 

### Return type

[**RouteFilterRulesData**](RouteFilterRulesData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRouteFilterRuleChangeByUuid**
> RouteFilterRulesChangeData GetRouteFilterRuleChangeByUuid(ctx, routeFilterId, routeFilterRuleId, changeId)
Get Change By ID

This API provides capability to retrieve a specific Route Filter Rule's Changes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **routeFilterId** | [**string**](.md)| Route Filters Id | 
  **routeFilterRuleId** | [**string**](.md)| Route  Filter  Rules Id | 
  **changeId** | [**string**](.md)| Route Filter Rule Change UUID | 

### Return type

[**RouteFilterRulesChangeData**](RouteFilterRulesChangeData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRouteFilterRuleChanges**
> RouteFilterRulesChangeDataResponse GetRouteFilterRuleChanges(ctx, routeFilterId, routeFilterRuleId, optional)
Get All Changes

This API provides capability to retrieve all of a Route Filter Rule's Changes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **routeFilterId** | [**string**](.md)| Route Filters Id | 
  **routeFilterRuleId** | [**string**](.md)| Route  Filter  Rules Id | 
 **optional** | ***RouteFilterRulesApiGetRouteFilterRuleChangesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RouteFilterRulesApiGetRouteFilterRuleChangesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **offset** | **optional.Int32**| offset | 
 **limit** | **optional.Int32**| number of records to fetch | 

### Return type

[**RouteFilterRulesChangeDataResponse**](RouteFilterRulesChangeDataResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRouteFilterRules**
> GetRouteFilterRulesResponse GetRouteFilterRules(ctx, routeFilterId, optional)
GetRFRules

This API provides capability to get all Route Filters Rules for Fabric

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **routeFilterId** | [**string**](.md)| Route Filters Id | 
 **optional** | ***RouteFilterRulesApiGetRouteFilterRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RouteFilterRulesApiGetRouteFilterRulesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **optional.Int32**| offset | 
 **limit** | **optional.Int32**| number of records to fetch | 

### Return type

[**GetRouteFilterRulesResponse**](GetRouteFilterRulesResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchRouteFilterRuleByUuid**
> RouteFiltersData PatchRouteFilterRuleByUuid(ctx, body, routeFilterId, routeFilterRuleId)
PatchRFilterRule

This API provides capability to partially update a Route Filter Rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]RouteFilterChangePrefixMatch**](RouteFilterChangePrefixMatch.md)|  | 
  **routeFilterId** | [**string**](.md)| Route Filters Id | 
  **routeFilterRuleId** | [**string**](.md)| Route  Filter  Rules Id | 

### Return type

[**RouteFiltersData**](RouteFiltersData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceRouteFilterRuleByUuid**
> RouteFilterRulesData ReplaceRouteFilterRuleByUuid(ctx, body, routeFilterId, routeFilterRuleId)
ReplaceRFRule

This API provides capability to replace a Route Filter Rule completely

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RouteFilterRulesBase**](RouteFilterRulesBase.md)|  | 
  **routeFilterId** | [**string**](.md)| Route Filters Id | 
  **routeFilterRuleId** | [**string**](.md)| Route  Filter  Rules Id | 

### Return type

[**RouteFilterRulesData**](RouteFilterRulesData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

