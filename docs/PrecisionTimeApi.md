# {{classname}}

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTimeServices**](PrecisionTimeApi.md#CreateTimeServices) | **Post** /fabric/v4/timeServices | Create Time Service
[**DeleteTimeServiceById**](PrecisionTimeApi.md#DeleteTimeServiceById) | **Delete** /fabric/v4/timeServices/{serviceId} | Delete time service
[**GetTimeServicesById**](PrecisionTimeApi.md#GetTimeServicesById) | **Get** /fabric/v4/timeServices/{serviceId} | Get Time Service
[**GetTimeServicesConnectionsByServiceId**](PrecisionTimeApi.md#GetTimeServicesConnectionsByServiceId) | **Get** /fabric/v4/timeServices/{serviceId}/connections | Get Connection Links
[**GetTimeServicesPackageByCode**](PrecisionTimeApi.md#GetTimeServicesPackageByCode) | **Get** /fabric/v4/timeServicePackages/{packageCode} | Get Package By Code
[**GetTimeServicesPackages**](PrecisionTimeApi.md#GetTimeServicesPackages) | **Get** /fabric/v4/timeServicePackages | Get Packages
[**UpdateTimeServicesById**](PrecisionTimeApi.md#UpdateTimeServicesById) | **Patch** /fabric/v4/timeServices/{serviceId} | Patch time service

# **CreateTimeServices**
> PrecisionTimeServiceCreateResponse CreateTimeServices(ctx, body, optional)
Create Time Service

The API provides capability to create timing service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PrecisionTimeServiceRequest**](PrecisionTimeServiceRequest.md)|  | 
 **optional** | ***PrecisionTimeApiCreateTimeServicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PrecisionTimeApiCreateTimeServicesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCORRELATIONID** | **optional.**| Correlation identifier | 
 **xAUTHUSERNAME** | **optional.**| User name | 

### Return type

[**PrecisionTimeServiceCreateResponse**](precisionTimeServiceCreateResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTimeServiceById**
> PrecisionTimeServiceCreateResponse DeleteTimeServiceById(ctx, serviceId, optional)
Delete time service

Delete EPT service by it's uuid

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | [**string**](.md)| Service UUID | 
 **optional** | ***PrecisionTimeApiDeleteTimeServiceByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PrecisionTimeApiDeleteTimeServiceByIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCORRELATIONID** | **optional.String**| Correlation identifier | 
 **xAUTHUSERNAME** | **optional.String**| User name | 

### Return type

[**PrecisionTimeServiceCreateResponse**](precisionTimeServiceCreateResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTimeServicesById**
> PrecisionTimeServiceCreateResponse GetTimeServicesById(ctx, serviceId, optional)
Get Time Service

The API provides capability to get precision timing service's details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | [**string**](.md)| Service UUID | 
 **optional** | ***PrecisionTimeApiGetTimeServicesByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PrecisionTimeApiGetTimeServicesByIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCORRELATIONID** | **optional.String**| Correlation identifier | 
 **xAUTHUSERNAME** | **optional.String**| User name | 

### Return type

[**PrecisionTimeServiceCreateResponse**](precisionTimeServiceCreateResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTimeServicesConnectionsByServiceId**
> PrecisionTimeServiceConnectionsResponse GetTimeServicesConnectionsByServiceId(ctx, serviceId, optional)
Get Connection Links

The API provides capability to get prevision timing service's details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | [**string**](.md)| Service UUID | 
 **optional** | ***PrecisionTimeApiGetTimeServicesConnectionsByServiceIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PrecisionTimeApiGetTimeServicesConnectionsByServiceIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCORRELATIONID** | **optional.String**| Correlation identifier | 
 **xAUTHUSERNAME** | **optional.String**| User name | 

### Return type

[**PrecisionTimeServiceConnectionsResponse**](precisionTimeServiceConnectionsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTimeServicesPackageByCode**
> PrecisionTimePackageResponse GetTimeServicesPackageByCode(ctx, packageCode, optional)
Get Package By Code

The API provides capability to get timing service's package by code

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **packageCode** | **string**| Package Code | 
 **optional** | ***PrecisionTimeApiGetTimeServicesPackageByCodeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PrecisionTimeApiGetTimeServicesPackageByCodeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xCORRELATIONID** | **optional.String**| Correlation identifier | 
 **xAUTHUSERNAME** | **optional.String**| User name | 

### Return type

[**PrecisionTimePackageResponse**](precisionTimePackageResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTimeServicesPackages**
> PrecisionTimeServicePackagesResponse GetTimeServicesPackages(ctx, optional)
Get Packages

The API provides capability to get timing service's packages

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PrecisionTimeApiGetTimeServicesPackagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PrecisionTimeApiGetTimeServicesPackagesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xCORRELATIONID** | **optional.String**| Correlation identifier | 
 **xAUTHUSERNAME** | **optional.String**| User name | 

### Return type

[**PrecisionTimeServicePackagesResponse**](precisionTimeServicePackagesResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTimeServicesById**
> PrecisionTimeServiceCreateResponse UpdateTimeServicesById(ctx, body, serviceId, optional)
Patch time service

The API provides capability to update timing service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]PrecisionTimeChangeOperation**](precisionTimeChangeOperation.md)|  | 
  **serviceId** | [**string**](.md)| Service UUID | 
 **optional** | ***PrecisionTimeApiUpdateTimeServicesByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PrecisionTimeApiUpdateTimeServicesByIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xCORRELATIONID** | **optional.**| Correlation identifier | 
 **xAUTHUSERNAME** | **optional.**| User name | 

### Return type

[**PrecisionTimeServiceCreateResponse**](precisionTimeServiceCreateResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json-patch+json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

