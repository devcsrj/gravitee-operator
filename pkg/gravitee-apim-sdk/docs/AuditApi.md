# \AuditApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEvents1**](AuditApi.md#GetEvents1) | **Get** /audit/events | 
[**List4**](AuditApi.md#List4) | **Get** /audit | 



## GetEvents1

> GetEvents1(ctx, )



### Required Parameters

This endpoint does not need any parameter.

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


## List4

> MetadataPageAuditEntity List4(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***List4Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a List4Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mgmt** | **optional.Bool**|  | 
 **api** | **optional.String**|  | 
 **application** | **optional.String**|  | 
 **event** | **optional.String**|  | 
 **from** | **optional.Int64**|  | 
 **to** | **optional.Int64**|  | 
 **size** | **optional.Int32**|  | [default to 20]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**MetadataPageAuditEntity**](MetadataPageAuditEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

