# \RatingApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create2**](RatingApi.md#Create2) | **Post** /apis/{api}/ratings | 
[**CreateAnswer**](RatingApi.md#CreateAnswer) | **Post** /apis/{api}/ratings/{rating}/answers | 
[**Delete2**](RatingApi.md#Delete2) | **Delete** /apis/{api}/ratings/{rating} | 
[**Delete3**](RatingApi.md#Delete3) | **Delete** /apis/{api}/ratings/{rating}/answers/{answer} | 
[**GetByApiAndUser**](RatingApi.md#GetByApiAndUser) | **Get** /apis/{api}/ratings/current | 
[**GetSummaryByApi**](RatingApi.md#GetSummaryByApi) | **Get** /apis/{api}/ratings/summary | 
[**List**](RatingApi.md#List) | **Get** /apis/{api}/ratings | 
[**Update3**](RatingApi.md#Update3) | **Put** /apis/{api}/ratings/{rating} | 



## Create2

> RatingEntity Create2(ctx, api, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
 **optional** | ***Create2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a Create2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of NewRatingEntity**](NewRatingEntity.md)|  | 

### Return type

[**RatingEntity**](RatingEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAnswer

> RatingEntity CreateAnswer(ctx, api, rating, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
**rating** | **string**|  | 
 **optional** | ***CreateAnswerOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateAnswerOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of NewRatingAnswerEntity**](NewRatingAnswerEntity.md)|  | 

### Return type

[**RatingEntity**](RatingEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete2

> Delete2(ctx, rating)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rating** | **string**|  | 

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


## Delete3

> Delete3(ctx, rating, answer)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**rating** | **string**|  | 
**answer** | **string**|  | 

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


## GetByApiAndUser

> RatingEntity GetByApiAndUser(ctx, api)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 

### Return type

[**RatingEntity**](RatingEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSummaryByApi

> RatingSummaryEntity GetSummaryByApi(ctx, api)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 

### Return type

[**RatingSummaryEntity**](RatingSummaryEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> PageRatingEntity List(ctx, api, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
 **optional** | ***ListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNumber** | **optional.Int32**|  | 
 **pageSize** | **optional.Int32**|  | 

### Return type

[**PageRatingEntity**](PageRatingEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update3

> RatingEntity Update3(ctx, api, rating, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
**rating** | **string**|  | 
 **optional** | ***Update3Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a Update3Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of UpdateRatingEntity**](UpdateRatingEntity.md)|  | 

### Return type

[**RatingEntity**](RatingEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

