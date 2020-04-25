# \FetcherApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFetcher**](FetcherApi.md#GetFetcher) | **Get** /fetchers/{fetcher} | 
[**GetFetcherSchema**](FetcherApi.md#GetFetcherSchema) | **Get** /fetchers/{fetcher}/schema | 
[**List14**](FetcherApi.md#List14) | **Get** /fetchers | List fetchers



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

