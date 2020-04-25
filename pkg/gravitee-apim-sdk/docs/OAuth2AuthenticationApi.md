# \OAuth2AuthenticationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExchangeAuthorizationCode**](OAuth2AuthenticationApi.md#ExchangeAuthorizationCode) | **Post** /auth/oauth2/{identity} | 
[**TokenExchange**](OAuth2AuthenticationApi.md#TokenExchange) | **Post** /auth/oauth2/{identity}/exchange | 



## ExchangeAuthorizationCode

> ExchangeAuthorizationCode(ctx, identity, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identity** | **string**|  | 
 **optional** | ***ExchangeAuthorizationCodeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ExchangeAuthorizationCodeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Payload**](Payload.md)|  | 

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


## TokenExchange

> TokenExchange(ctx, identity, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identity** | **string**|  | 
 **optional** | ***TokenExchangeOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TokenExchangeOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | **optional.String**|  | 

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

