# \GatewayApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create17**](GatewayApi.md#Create17) | **Post** /platform/tickets | 
[**GetInstance**](GatewayApi.md#GetInstance) | **Get** /instances/{instance} | Get a gateway instance
[**InstanceMonitoring**](GatewayApi.md#InstanceMonitoring) | **Get** /instances/{instance}/monitoring/{gatewayId} | Get monitoring metrics for a gateway instance
[**ListEvents**](GatewayApi.md#ListEvents) | **Get** /platform/events | 
[**ListInstances**](GatewayApi.md#ListInstances) | **Get** /instances | List gateway instances
[**PlatformAnalytics**](GatewayApi.md#PlatformAnalytics) | **Get** /platform/analytics | 



## Create17

> Create17(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***Create17Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a Create17Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of NewTicketEntity**](NewTicketEntity.md)|  | 

### Return type

 (empty response body)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstance

> InstanceEntity GetInstance(ctx, instance)

Get a gateway instance

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instance** | **string**|  | 

### Return type

[**InstanceEntity**](InstanceEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstanceMonitoring

> MonitoringData InstanceMonitoring(ctx, gatewayId, instance)

Get monitoring metrics for a gateway instance

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **string**|  | 
**instance** | **string**|  | 

### Return type

[**MonitoringData**](MonitoringData.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEvents

> PageEventEntity ListEvents(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListEventsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **optional.String**|  | [default to all]
 **from** | **optional.Int64**|  | 
 **to** | **optional.Int64**|  | 
 **page** | **optional.Int32**|  | [default to 0]
 **size** | **optional.Int32**|  | [default to 10]
 **apiIds** | **optional.String**|  | 

### Return type

[**PageEventEntity**](PageEventEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInstances

> []InstanceListItem ListInstances(ctx, optional)

List gateway instances

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListInstancesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListInstancesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeStopped** | **optional.Bool**|  | 

### Return type

[**[]InstanceListItem**](InstanceListItem.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlatformAnalytics

> PlatformAnalytics(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PlatformAnalyticsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PlatformAnalyticsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **optional.Int64**|  | 
 **to** | **optional.Int64**|  | 
 **interval** | **optional.Int64**|  | 
 **query** | **optional.String**|  | 
 **key** | **optional.String**|  | 
 **field** | **optional.String**|  | 
 **size** | **optional.Int32**|  | 
 **type_** | **optional.String**|  | 
 **ranges** | **optional.String**|  | 
 **order** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

