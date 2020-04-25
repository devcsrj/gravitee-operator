# \MetadataApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create1**](MetadataApi.md#Create1) | **Post** /apis/{api}/metadata | Create an API metadata
[**Create11**](MetadataApi.md#Create11) | **Post** /configuration/metadata | 
[**Delete1**](MetadataApi.md#Delete1) | **Delete** /apis/{api}/metadata/{metadata} | Delete a metadata
[**Delete13**](MetadataApi.md#Delete13) | **Delete** /configuration/metadata/{metadata} | 
[**GetApiMetadata**](MetadataApi.md#GetApiMetadata) | **Get** /apis/{api}/metadata/{metadata} | A metadata for the given API and metadata id
[**List9**](MetadataApi.md#List9) | **Get** /configuration/metadata | 
[**ListApiMetadatas**](MetadataApi.md#ListApiMetadatas) | **Get** /apis/{api}/metadata | List metadata for the given API
[**Update15**](MetadataApi.md#Update15) | **Put** /configuration/metadata | 
[**Update2**](MetadataApi.md#Update2) | **Put** /apis/{api}/metadata/{metadata} | Update an API metadata



## Create1

> ApiMetadataEntity Create1(ctx, api, optional)

Create an API metadata

User must have the MANAGE_API permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
 **optional** | ***Create1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a Create1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of NewApiMetadataEntity**](NewApiMetadataEntity.md)|  | 

### Return type

[**ApiMetadataEntity**](ApiMetadataEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Create11

> MetadataEntity Create11(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***Create11Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a Create11Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of NewMetadataEntity**](NewMetadataEntity.md)|  | 

### Return type

[**MetadataEntity**](MetadataEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete1

> Delete1(ctx, api, metadata)

Delete a metadata

User must have the MANAGE_API permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
**metadata** | **string**|  | 

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


## Delete13

> Delete13(ctx, metadata)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metadata** | **string**|  | 

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


## GetApiMetadata

> ApiMetadataEntity GetApiMetadata(ctx, api, metadata)

A metadata for the given API and metadata id

User must have the MANAGE_API permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
**metadata** | **string**|  | 

### Return type

[**ApiMetadataEntity**](ApiMetadataEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List9

> []MetadataEntity List9(ctx, )



### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]MetadataEntity**](MetadataEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApiMetadatas

> []ApiMetadataEntity ListApiMetadatas(ctx, api)

List metadata for the given API

User must have the MANAGE_API permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 

### Return type

[**[]ApiMetadataEntity**](ApiMetadataEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update15

> MetadataEntity Update15(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***Update15Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a Update15Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of UpdateMetadataEntity**](UpdateMetadataEntity.md)|  | 

### Return type

[**MetadataEntity**](MetadataEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update2

> ApiMetadataEntity Update2(ctx, api, optional)

Update an API metadata

User must have the MANAGE_API permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
 **optional** | ***Update2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a Update2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpdateApiMetadataEntity**](UpdateApiMetadataEntity.md)|  | 

### Return type

[**ApiMetadataEntity**](ApiMetadataEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

