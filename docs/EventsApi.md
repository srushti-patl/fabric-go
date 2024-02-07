# {{classname}}

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEvent**](EventsApi.md#GetEvent) | **Get** /fabric/v4/events/{eventId} | Get Event
[**GetEventsByAsset**](EventsApi.md#GetEventsByAsset) | **Get** /fabric/v4/events | Get Events of Asset

# **GetEvent**
> GetEvent(ctx, eventId, type_, optional)
Get Event

The API provides capability to get list of events for user Asset, including pagination

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **eventId** | [**string**](.md)| Event UUID | 
  **type_** | **string**| type of events | 
 **optional** | ***EventsApiGetEventOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EventsApiGetEventOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **connectionId** | [**optional.Interface of string**](.md)| Connection UUID | 
 **portId** | [**optional.Interface of string**](.md)| Port UUID | 
 **xCORRELATIONID** | **optional.String**| Correlation identifier | 
 **xAUTHUSERNAME** | **optional.String**| User name | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventsByAsset**
> InlineResponse200 GetEventsByAsset(ctx, type_, optional)
Get Events of Asset

The API provides capability to get list of events for user Asset, including pagination

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**| type of events | 
 **optional** | ***EventsApiGetEventsByAssetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EventsApiGetEventsByAssetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectionId** | [**optional.Interface of string**](.md)| Connection UUID | 
 **portId** | [**optional.Interface of string**](.md)| Port UUID | 
 **routerId** | [**optional.Interface of string**](.md)| Cloud Router UUID | 
 **routingProtocolId** | [**optional.Interface of string**](.md)| Routing Protocol UUID | 
 **startDateTime** | [**optional.Interface of time.Time**](.md)| Start date and time | 
 **endDateTime** | [**optional.Interface of time.Time**](.md)| End date and time | 
 **xCORRELATIONID** | **optional.String**| Correlation identifier | 
 **xAUTHUSERNAME** | **optional.String**| User name | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

