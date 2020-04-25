# \PortalApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create13**](PortalApi.md#Create13) | **Post** /configuration/notificationsettings | Create notification settings
[**CreatePage1**](PortalApi.md#CreatePage1) | **Post** /portal/pages | Create a page
[**Delete16**](PortalApi.md#Delete16) | **Delete** /configuration/notificationsettings/{notificationId} | Delete notification settings
[**DeletePage1**](PortalApi.md#DeletePage1) | **Delete** /portal/pages/{page} | Delete a page
[**ExchangeAuthorizationCode**](PortalApi.md#ExchangeAuthorizationCode) | **Post** /auth/oauth2/{identity} | 
[**FetchAllPages1**](PortalApi.md#FetchAllPages1) | **Post** /portal/pages/_fetch | Refresh all pages by calling their associated fetcher
[**FetchPage1**](PortalApi.md#FetchPage1) | **Post** /portal/pages/{page}/_fetch | Refresh page by calling the associated fetcher
[**Get7**](PortalApi.md#Get7) | **Get** /configuration/notificationsettings | Get notification settings
[**GetConfig**](PortalApi.md#GetConfig) | **Get** /portal | 
[**GetImage1**](PortalApi.md#GetImage1) | **Get** /portal/media/{hash} | 
[**GetPage1**](PortalApi.md#GetPage1) | **Get** /portal/pages/{page} | Get a page
[**GetPageContent1**](PortalApi.md#GetPageContent1) | **Get** /portal/pages/{page}/content | Get the page&#39;s content
[**ListPages1**](PortalApi.md#ListPages1) | **Get** /portal/pages | List pages
[**ListSocialIdentityProvider**](PortalApi.md#ListSocialIdentityProvider) | **Get** /portal/identities | Get the list of social identity providers
[**PartialUpdatePage1**](PortalApi.md#PartialUpdatePage1) | **Patch** /portal/pages/{page} | Update a page
[**SavePortalConfig**](PortalApi.md#SavePortalConfig) | **Post** /portal | Save the portal configuration
[**SearchPortalApis**](PortalApi.md#SearchPortalApis) | **Post** /portal/apis/_search | Search for API using the search engine
[**TokenExchange**](PortalApi.md#TokenExchange) | **Post** /auth/oauth2/{identity}/exchange | 
[**Update17**](PortalApi.md#Update17) | **Put** /configuration/notificationsettings/{notificationId} | Update generic notification settings
[**Update18**](PortalApi.md#Update18) | **Put** /configuration/notificationsettings | Update portal notification settings
[**UpdatePage1**](PortalApi.md#UpdatePage1) | **Put** /portal/pages/{page} | Update a page
[**Upload**](PortalApi.md#Upload) | **Post** /portal/media/upload | 



## Create13

> map[string]interface{} Create13(ctx, optional)

Create notification settings

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***Create13Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a Create13Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of GenericNotificationConfigEntity**](GenericNotificationConfigEntity.md)|  | 

### Return type

**map[string]interface{}**

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePage1

> PageEntity CreatePage1(ctx, page)

Create a page

User must be ADMIN to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**page** | [**NewPageEntity**](NewPageEntity.md)|  | 

### Return type

[**PageEntity**](PageEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete16

> Delete16(ctx, notificationId)

Delete notification settings

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationId** | **string**|  | 

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


## DeletePage1

> DeletePage1(ctx, page)

Delete a page

User must be ADMIN to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**page** | **string**|  | 

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


## FetchAllPages1

> PageEntity FetchAllPages1(ctx, )

Refresh all pages by calling their associated fetcher

User must have the MANAGE_PAGES permission to use this service

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**PageEntity**](PageEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchPage1

> PageEntity FetchPage1(ctx, page)

Refresh page by calling the associated fetcher

User must have the MANAGE_PAGES permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**page** | **string**|  | 

### Return type

[**PageEntity**](PageEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get7

> []map[string]interface{} Get7(ctx, )

Get notification settings

### Required Parameters

This endpoint does not need any parameter.

### Return type

**[]map[string]interface{}**

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfig

> PortalConfigEntity GetConfig(ctx, )



### Required Parameters

This endpoint does not need any parameter.

### Return type

[**PortalConfigEntity**](PortalConfigEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImage1

> GetImage1(ctx, hash)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string**|  | 

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


## GetPage1

> GetPage1(ctx, page)

Get a page

Every users can use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**page** | **string**|  | 

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


## GetPageContent1

> GetPageContent1(ctx, page)

Get the page's content

Every users can use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**page** | **string**|  | 

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


## ListPages1

> []PageEntity ListPages1(ctx, optional)

List pages

Every users can use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListPages1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListPages1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **homepage** | **optional.Bool**|  | 
 **type_** | **optional.String**|  | 
 **parent** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **root** | **optional.Bool**|  | 

### Return type

[**[]PageEntity**](PageEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSocialIdentityProvider

> []SocialIdentityProviderEntity ListSocialIdentityProvider(ctx, )

Get the list of social identity providers

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]SocialIdentityProviderEntity**](SocialIdentityProviderEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdatePage1

> PageEntity PartialUpdatePage1(ctx, page, page)

Update a page

User must be ADMIN to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**page** | **string**|  | 
**page** | [**UpdatePageEntity**](UpdatePageEntity.md)|  | 

### Return type

[**PageEntity**](PageEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SavePortalConfig

> SavePortalConfig(ctx, config)

Save the portal configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**config** | [**PortalConfigEntity**](PortalConfigEntity.md)|  | 

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


## SearchPortalApis

> []ApiListItem SearchPortalApis(ctx, q)

Search for API using the search engine

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**q** | **string**|  | 

### Return type

[**[]ApiListItem**](ApiListItem.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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


## Update17

> GenericNotificationConfigEntity Update17(ctx, notificationId, optional)

Update generic notification settings

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationId** | **string**|  | 
 **optional** | ***Update17Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a Update17Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of GenericNotificationConfigEntity**](GenericNotificationConfigEntity.md)|  | 

### Return type

[**GenericNotificationConfigEntity**](GenericNotificationConfigEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update18

> PortalNotificationConfigEntity Update18(ctx, optional)

Update portal notification settings

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***Update18Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a Update18Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of PortalNotificationConfigEntity**](PortalNotificationConfigEntity.md)|  | 

### Return type

[**PortalNotificationConfigEntity**](PortalNotificationConfigEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePage1

> PageEntity UpdatePage1(ctx, page, page)

Update a page

User must be ADMIN to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**page** | **string**|  | 
**page** | [**UpdatePageEntity**](UpdatePageEntity.md)|  | 

### Return type

[**PageEntity**](PageEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Upload

> Upload(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UploadOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UploadOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | **optional.Interface of *os.File****optional.*os.File**|  | 

### Return type

 (empty response body)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

