# \PluginApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFetcher**](PluginApi.md#GetFetcher) | **Get** /fetchers/{fetcher} | 
[**GetFetcherSchema**](PluginApi.md#GetFetcherSchema) | **Get** /fetchers/{fetcher}/schema | 
[**GetPolicy**](PluginApi.md#GetPolicy) | **Get** /policies/{policy} | Get a policy
[**GetPolicySchema**](PluginApi.md#GetPolicySchema) | **Get** /policies/{policy}/schema | Get a policy&#39;s schema
[**GetResource**](PluginApi.md#GetResource) | **Get** /resources/{resource} | Get a resource
[**GetResourceSchema**](PluginApi.md#GetResourceSchema) | **Get** /resources/{resource}/schema | Get a resource&#39;s schema
[**GetServiceDiscovery**](PluginApi.md#GetServiceDiscovery) | **Get** /services-discovery/{plugin} | Get a service discovery
[**GetServiceDiscoverySchema**](PluginApi.md#GetServiceDiscoverySchema) | **Get** /services-discovery/{plugin}/schema | Get a service discovery&#39;s schema
[**List14**](PluginApi.md#List14) | **Get** /fetchers | List fetchers
[**ListPolicies**](PluginApi.md#ListPolicies) | **Get** /policies | List policies
[**ListResources**](PluginApi.md#ListResources) | **Get** /resources | List resources
[**ListResources1**](PluginApi.md#ListResources1) | **Get** /services-discovery | List service discovery plugins



## GetFetcher

> FetcherEntity GetFetcher(ctx, fetcher)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fetcher** | **string**|  | 

### Return type

[**FetcherEntity**](FetcherEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFetcherSchema

> string GetFetcherSchema(ctx, fetcher)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fetcher** | **string**|  | 

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


## GetPolicy

> PolicyEntity GetPolicy(ctx, policy)

Get a policy

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policy** | **string**|  | 

### Return type

[**PolicyEntity**](PolicyEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicySchema

> string GetPolicySchema(ctx, policy)

Get a policy's schema

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policy** | **string**|  | 

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


## GetResource

> PluginEntity GetResource(ctx, resource)

Get a resource

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resource** | **string**|  | 

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


## GetResourceSchema

> string GetResourceSchema(ctx, resource)

Get a resource's schema

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resource** | **string**|  | 

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


## List14

> []FetcherListItem List14(ctx, optional)

List fetchers

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***List14Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a List14Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expand** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**[]FetcherListItem**](FetcherListItem.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPolicies

> []PolicyListItem ListPolicies(ctx, optional)

List policies

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListPoliciesOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expand** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**[]PolicyListItem**](PolicyListItem.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListResources

> []ResourceListItem ListResources(ctx, optional)

List resources

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListResourcesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListResourcesOpts struct


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

