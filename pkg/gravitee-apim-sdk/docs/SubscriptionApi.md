# \SubscriptionApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeSubscriptionStatus**](SubscriptionApi.md#ChangeSubscriptionStatus) | **Post** /apis/{api}/subscriptions/{subscription}/status | Change the status of a subscription
[**CloseSubscription**](SubscriptionApi.md#CloseSubscription) | **Delete** /applications/{application}/subscriptions/{subscription} | Close the subscription
[**CreateSubscription**](SubscriptionApi.md#CreateSubscription) | **Post** /apis/{api}/subscriptions | Subscribe to a plan
[**CreateSubscription1**](SubscriptionApi.md#CreateSubscription1) | **Post** /applications/{application}/subscriptions | Subscribe to a plan
[**GetApiSubscription**](SubscriptionApi.md#GetApiSubscription) | **Get** /apis/{api}/subscriptions/{subscription} | Get a subscription
[**GetSubscription**](SubscriptionApi.md#GetSubscription) | **Get** /applications/{application}/subscriptions/{subscription} | Get subscription information
[**ListApiKeysForSubscription**](SubscriptionApi.md#ListApiKeysForSubscription) | **Get** /apis/{api}/subscriptions/{subscription}/keys | List all API Keys for a subscription
[**ListApiKeysForSubscription1**](SubscriptionApi.md#ListApiKeysForSubscription1) | **Get** /applications/{application}/subscriptions/{subscription}/keys | List all API Keys for a subscription
[**ListApiSubscribed**](SubscriptionApi.md#ListApiSubscribed) | **Get** /applications/{application}/subscribed | List APIs subscribed by the application
[**ListApiSubscribers**](SubscriptionApi.md#ListApiSubscribers) | **Get** /apis/{api}/subscribers | List subscribers for the API
[**ListApiSubscriptions**](SubscriptionApi.md#ListApiSubscriptions) | **Get** /apis/{api}/subscriptions | List subscriptions for the API
[**ListApplicationSubscriptions**](SubscriptionApi.md#ListApplicationSubscriptions) | **Get** /applications/{application}/subscriptions | List subscriptions for the application
[**ListUserSubscriptions**](SubscriptionApi.md#ListUserSubscriptions) | **Get** /subscriptions | List subscriptions for authenticated user
[**ProcessApiSubscription**](SubscriptionApi.md#ProcessApiSubscription) | **Post** /apis/{api}/subscriptions/{subscription}/_process | Update a subscription
[**RenewApiKey**](SubscriptionApi.md#RenewApiKey) | **Post** /apis/{api}/subscriptions/{subscription} | Renew an API key
[**RenewApiKey1**](SubscriptionApi.md#RenewApiKey1) | **Post** /applications/{application}/subscriptions/{subscription} | Renew an API key
[**RevokeApiKey1**](SubscriptionApi.md#RevokeApiKey1) | **Delete** /apis/{api}/subscriptions/{subscription}/keys/{key} | Revoke an API key
[**RevokeApiKey2**](SubscriptionApi.md#RevokeApiKey2) | **Delete** /applications/{application}/subscriptions/{subscription}/keys/{key} | Revoke an API key
[**UpdateApiSubscription**](SubscriptionApi.md#UpdateApiSubscription) | **Put** /apis/{api}/subscriptions/{subscription} | Update a subscription



## ChangeSubscriptionStatus

> Subscription ChangeSubscriptionStatus(ctx, api, subscription, status)

Change the status of a subscription

User must have the MANAGE_PLANS permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
**subscription** | **string**|  | 
**status** | **string**|  | 

### Return type

[**Subscription**](Subscription.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloseSubscription

> Subscription CloseSubscription(ctx, application, subscription)

Close the subscription

User must have the APPLICATION_SUBSCRIPTION[DELETE] permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string**|  | 
**subscription** | **string**|  | 

### Return type

[**Subscription**](Subscription.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSubscription

> Subscription CreateSubscription(ctx, api, application, plan)

Subscribe to a plan

User must have the MANAGE_SUBSCRIPTIONS permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
**application** | **string**|  | 
**plan** | **string**|  | 

### Return type

[**Subscription**](Subscription.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSubscription1

> Subscription CreateSubscription1(ctx, application, plan, optional)

Subscribe to a plan

User must have the MANAGE_SUBSCRIPTIONS permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string**|  | 
**plan** | **string**|  | 
 **optional** | ***CreateSubscription1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateSubscription1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of NewSubscriptionEntity**](NewSubscriptionEntity.md)|  | 

### Return type

[**Subscription**](Subscription.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiSubscription

> Subscription GetApiSubscription(ctx, api, subscription)

Get a subscription

User must have the MANAGE_PLANS permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
**subscription** | **string**|  | 

### Return type

[**Subscription**](Subscription.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubscription

> Subscription GetSubscription(ctx, subscription)

Get subscription information

User must have the READ permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscription** | **string**|  | 

### Return type

[**Subscription**](Subscription.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApiKeysForSubscription

> []ApiKeyEntity ListApiKeysForSubscription(ctx, subscription)

List all API Keys for a subscription

User must have the MANAGE_API_KEYS permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscription** | **string**|  | 

### Return type

[**[]ApiKeyEntity**](ApiKeyEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApiKeysForSubscription1

> []ApiKeyEntity ListApiKeysForSubscription1(ctx, subscription)

List all API Keys for a subscription

User must have the READ permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscription** | **string**|  | 

### Return type

[**[]ApiKeyEntity**](ApiKeyEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApiSubscribed

> []ApplicationEntity ListApiSubscribed(ctx, application)

List APIs subscribed by the application

User must have the APPLICATION_SUBSCRIPTION permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string**|  | 

### Return type

[**[]ApplicationEntity**](ApplicationEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApiSubscribers

> []ApplicationEntity ListApiSubscribers(ctx, api)

List subscribers for the API

User must have the MANAGE_SUBSCRIPTIONS permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 

### Return type

[**[]ApplicationEntity**](ApplicationEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApiSubscriptions

> PagedResult ListApiSubscriptions(ctx, api, plan, application, optional)

List subscriptions for the API

User must have the READ_SUBSCRIPTION permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
**plan** | **string**| plan | 
**application** | **string**| application | 
 **optional** | ***ListApiSubscriptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListApiSubscriptionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **status** | **optional.String**|  | [default to accepted,pending,paused]
 **size** | **optional.Int32**|  | [default to 20]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**PagedResult**](PagedResult.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplicationSubscriptions

> PagedResult ListApplicationSubscriptions(ctx, application, plan, api, optional)

List subscriptions for the application

User must have the READ_SUBSCRIPTION permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string**|  | 
**plan** | **string**| plan | 
**api** | **string**| api | 
 **optional** | ***ListApplicationSubscriptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListApplicationSubscriptionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **status** | **optional.String**|  | [default to accepted,pending,paused]
 **size** | **optional.Int32**|  | [default to 20]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**PagedResult**](PagedResult.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserSubscriptions

> []Subscription ListUserSubscriptions(ctx, optional)

List subscriptions for authenticated user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListUserSubscriptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListUserSubscriptionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **application** | **optional.String**|  | 
 **plan** | **optional.String**|  | 

### Return type

[**[]Subscription**](Subscription.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProcessApiSubscription

> Subscription ProcessApiSubscription(ctx, api, subscription, subscription)

Update a subscription

User must have the MANAGE_PLANS permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
**subscription** | **string**|  | 
**subscription** | [**ProcessSubscriptionEntity**](ProcessSubscriptionEntity.md)|  | 

### Return type

[**Subscription**](Subscription.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenewApiKey

> ApiKeyEntity RenewApiKey(ctx, api, subscription)

Renew an API key

User must have the MANAGE_API_KEYS permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
**subscription** | **string**|  | 

### Return type

[**ApiKeyEntity**](ApiKeyEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenewApiKey1

> ApiKeyEntity RenewApiKey1(ctx, application, subscription)

Renew an API key

User must have the MANAGE_API_KEYS permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string**|  | 
**subscription** | **string**|  | 

### Return type

[**ApiKeyEntity**](ApiKeyEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeApiKey1

> RevokeApiKey1(ctx, api, subscription, key)

Revoke an API key

User must have the MANAGE_API_KEYS permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
**subscription** | **string**|  | 
**key** | **string**|  | 

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


## RevokeApiKey2

> RevokeApiKey2(ctx, application, subscription, key)

Revoke an API key

User must have the MANAGE_API_KEYS permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string**|  | 
**subscription** | **string**|  | 
**key** | **string**|  | 

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


## UpdateApiSubscription

> Subscription UpdateApiSubscription(ctx, api, subscription, subscription)

Update a subscription

User must have the MANAGE_PLANS permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
**subscription** | **string**|  | 
**subscription** | [**UpdateSubscriptionEntity**](UpdateSubscriptionEntity.md)|  | 

### Return type

[**Subscription**](Subscription.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

