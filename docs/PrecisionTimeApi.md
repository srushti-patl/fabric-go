# {{classname}}

All URIs are relative to *https://api.equinix.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTimeServices**](PrecisionTimeApi.md#CreateTimeServices) | **Post** /fabric/v4/timeServices | Create Time Service
[**DeleteTimeServiceById**](PrecisionTimeApi.md#DeleteTimeServiceById) | **Delete** /fabric/v4/timeServices/{serviceId} | Delete Time Service
[**GetTimeServicesById**](PrecisionTimeApi.md#GetTimeServicesById) | **Get** /fabric/v4/timeServices/{serviceId} | Get Time Service
[**GetTimeServicesConnectionsByServiceId**](PrecisionTimeApi.md#GetTimeServicesConnectionsByServiceId) | **Get** /fabric/v4/timeServices/{serviceId}/connections | Get all Connections
[**GetTimeServicesPackageByCode**](PrecisionTimeApi.md#GetTimeServicesPackageByCode) | **Get** /fabric/v4/timeServicePackages/{packageCode} | Get Package by Code
[**GetTimeServicesPackages**](PrecisionTimeApi.md#GetTimeServicesPackages) | **Get** /fabric/v4/timeServicePackages | Get Packages
[**UpdateTimeServicesById**](PrecisionTimeApi.md#UpdateTimeServicesById) | **Patch** /fabric/v4/timeServices/{serviceId} | Patch Time Service

# **CreateTimeServices**
> PrecisionTimeServiceCreateResponse CreateTimeServices(ctx, body)
Create Time Service

The API provides capability to create timing service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PrecisionTimeServiceRequest**](PrecisionTimeServiceRequest.md)|  | 

### Return type

[**PrecisionTimeServiceCreateResponse**](precisionTimeServiceCreateResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTimeServiceById**
> PrecisionTimeServiceCreateResponse DeleteTimeServiceById(ctx, serviceId)
Delete Time Service

Delete EPT service by it's uuid

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | [**string**](.md)| Service UUID | 

### Return type

[**PrecisionTimeServiceCreateResponse**](precisionTimeServiceCreateResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTimeServicesById**
> PrecisionTimeServiceCreateResponse GetTimeServicesById(ctx, serviceId)
Get Time Service

The API provides capability to get precision timing service's details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | [**string**](.md)| Service UUID | 

### Return type

[**PrecisionTimeServiceCreateResponse**](precisionTimeServiceCreateResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTimeServicesConnectionsByServiceId**
> PrecisionTimeServiceConnectionsResponse GetTimeServicesConnectionsByServiceId(ctx, serviceId)
Get all Connections

The API provides capability to get prevision timing service's details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | [**string**](.md)| Service UUID | 

### Return type

[**PrecisionTimeServiceConnectionsResponse**](precisionTimeServiceConnectionsResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTimeServicesPackageByCode**
> PackageResponse GetTimeServicesPackageByCode(ctx, packageCode)
Get Package by Code

The API provides capability to get timing service's package by code

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **packageCode** | **string**| Package Code | 

### Return type

[**PackageResponse**](PackageResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTimeServicesPackages**
> PrecisionTimeServicePackagesResponse GetTimeServicesPackages(ctx, )
Get Packages

The API provides capability to get timing service's packages

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PrecisionTimeServicePackagesResponse**](precisionTimeServicePackagesResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTimeServicesById**
> PrecisionTimeServiceCreateResponse UpdateTimeServicesById(ctx, body, serviceId)
Patch Time Service

The API provides capability to update timing service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]PrecisionTimeChangeOperation**](precisionTimeChangeOperation.md)|  | 
  **serviceId** | [**string**](.md)| Service UUID | 

### Return type

[**PrecisionTimeServiceCreateResponse**](precisionTimeServiceCreateResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

 - **Content-Type**: application/json-patch+json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

