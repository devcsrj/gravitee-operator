# \ServiceDiscoveryApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetServiceDiscovery**](ServiceDiscoveryApi.md#GetServiceDiscovery) | **Get** /services-discovery/{plugin} | Get a service discovery
[**GetServiceDiscoverySchema**](ServiceDiscoveryApi.md#GetServiceDiscoverySchema) | **Get** /services-discovery/{plugin}/schema | Get a service discovery&#39;s schema
[**ListResources1**](ServiceDiscoveryApi.md#ListResources1) | **Get** /services-discovery | List service discovery plugins



## GetServiceDiscovery

> PluginEntity GetServiceDiscovery(ctx, plugin)

Get a service discovery

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plugin** | **string**|  | 

### Return type

[**PluginEntity**](PluginEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceDiscoverySchema

> string GetServiceDiscoverySchema(ctx, plugin)

Get a service discovery's schema

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**plugin** | **string**|  | 

### Return type

**string**

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListResources1

> []ResourceListItem ListResources1(ctx, optional)

List service discovery plugins

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListResources1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListResources1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expand** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**[]ResourceListItem**](ResourceListItem.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

