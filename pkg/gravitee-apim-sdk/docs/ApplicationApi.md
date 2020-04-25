# \ApplicationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOrUpdateApplicationMember**](ApplicationApi.md#AddOrUpdateApplicationMember) | **Post** /applications/{application}/members | Add or update an application member
[**ApplicationLog**](ApplicationApi.md#ApplicationLog) | **Get** /applications/{application}/logs/{log} | Get a specific log
[**ApplicationLogs**](ApplicationApi.md#ApplicationLogs) | **Get** /applications/{application}/logs | Get application logs
[**CloseSubscription**](ApplicationApi.md#CloseSubscription) | **Delete** /applications/{application}/subscriptions/{subscription} | Close the subscription
[**Create5**](ApplicationApi.md#Create5) | **Post** /applications/{application}/notificationsettings | Create notification settings
[**Create6**](ApplicationApi.md#Create6) | **Post** /applications/{application}/alerts | 
[**CreateApplication**](ApplicationApi.md#CreateApplication) | **Post** /applications | Create an application
[**CreateSubscription1**](ApplicationApi.md#CreateSubscription1) | **Post** /applications/{application}/subscriptions | Subscribe to a plan
[**Delete6**](ApplicationApi.md#Delete6) | **Delete** /applications/{application}/notificationsettings/{notificationId} | Delete notification settings
[**Delete7**](ApplicationApi.md#Delete7) | **Delete** /applications/{application}/alerts/{alert} | 
[**DeleteApplication**](ApplicationApi.md#DeleteApplication) | **Delete** /applications/{application} | Delete an application
[**DeleteApplicationMember**](ApplicationApi.md#DeleteApplicationMember) | **Delete** /applications/{application}/members | Remove an application member
[**Get2**](ApplicationApi.md#Get2) | **Get** /applications/{application}/notificationsettings | Get notification settings
[**GetApplication**](ApplicationApi.md#GetApplication) | **Get** /applications/{application} | Get an application
[**GetHooks1**](ApplicationApi.md#GetHooks1) | **Get** /applications/hooks | Get the list of available hooks
[**GetNotifiers1**](ApplicationApi.md#GetNotifiers1) | **Get** /applications/{application}/notifiers | 
[**GetPermissions1**](ApplicationApi.md#GetPermissions1) | **Get** /applications/{application}/members/permissions | Get application members
[**GetSubscription**](ApplicationApi.md#GetSubscription) | **Get** /applications/{application}/subscriptions/{subscription} | Get subscription information
[**Hits1**](ApplicationApi.md#Hits1) | **Get** /applications/{application}/analytics | Get Application analytics
[**List3**](ApplicationApi.md#List3) | **Get** /applications/{application}/alerts | List configured alerts of a given APPLICATION
[**ListApiKeysForSubscription1**](ApplicationApi.md#ListApiKeysForSubscription1) | **Get** /applications/{application}/subscriptions/{subscription}/keys | List all API Keys for a subscription
[**ListApiSubscribed**](ApplicationApi.md#ListApiSubscribed) | **Get** /applications/{application}/subscribed | List APIs subscribed by the application
[**ListApplicationMembers**](ApplicationApi.md#ListApplicationMembers) | **Get** /applications/{application}/members | List application members
[**ListApplicationSubscriptions**](ApplicationApi.md#ListApplicationSubscriptions) | **Get** /applications/{application}/subscriptions | List subscriptions for the application
[**ListApplications**](ApplicationApi.md#ListApplications) | **Get** /applications | List all the applications accessible to authenticated user
[**RenewApiKey1**](ApplicationApi.md#RenewApiKey1) | **Post** /applications/{application}/subscriptions/{subscription} | Renew an API key
[**RevokeApiKey2**](ApplicationApi.md#RevokeApiKey2) | **Delete** /applications/{application}/subscriptions/{subscription}/keys/{key} | Revoke an API key
[**TransferOwnership1**](ApplicationApi.md#TransferOwnership1) | **Post** /applications/{application}/members/transfer_ownership | Transfer the ownership of the APPLICATION
[**Update6**](ApplicationApi.md#Update6) | **Put** /applications/{application}/notificationsettings/{notificationId} | Update generic notification settings
[**Update7**](ApplicationApi.md#Update7) | **Put** /applications/{application}/notificationsettings | Update portal notification settings
[**Update8**](ApplicationApi.md#Update8) | **Put** /applications/{application}/alerts/{alert} | 
[**UpdateApplication**](ApplicationApi.md#UpdateApplication) | **Put** /applications/{application} | Update an application



## AddOrUpdateApplicationMember

> AddOrUpdateApplicationMember(ctx, application, optional)

Add or update an application member

User must have the MANAGE_MEMBERS permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string**|  | 
 **optional** | ***AddOrUpdateApplicationMemberOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddOrUpdateApplicationMemberOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ApplicationMembership**](ApplicationMembership.md)|  | 

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


## ApplicationLog

> ApplicationLog(ctx, application, log, optional)

Get a specific log

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string**|  | 
**log** | **string**|  | 
 **optional** | ***ApplicationLogOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApplicationLogOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **timestamp** | **optional.Int64**|  | 

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


## ApplicationLogs

> ApplicationLogs(ctx, application, optional)

Get application logs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string**|  | 
 **optional** | ***ApplicationLogsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApplicationLogsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **optional.Int64**|  | 
 **to** | **optional.Int64**|  | 
 **query** | **optional.String**|  | 
 **size** | **optional.Int32**|  | [default to 20]
 **page** | **optional.Int32**|  | [default to 1]

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


## Create5

> map[string]interface{} Create5(ctx, application, optional)

Create notification settings

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string**|  | 
 **optional** | ***Create5Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a Create5Opts struct


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


## Create6

> AlertEntity Create6(ctx, application, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string**|  | 
 **optional** | ***Create6Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a Create6Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of NewAlertEntity**](NewAlertEntity.md)|  | 

### Return type

[**AlertEntity**](AlertEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApplication

> ApplicationEntity CreateApplication(ctx, application)

Create an application

User must have API_CONSUMER or ADMIN role to create an application.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | [**NewApplicationEntity**](NewApplicationEntity.md)|  | 

### Return type

[**ApplicationEntity**](ApplicationEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: application/json
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


## Delete6

> Delete6(ctx, application, notificationId)

Delete notification settings

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string**|  | 
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


## Delete7

> Delete7(ctx, application, alert)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string**|  | 
**alert** | **string**|  | 

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


## DeleteApplication

> DeleteApplication(ctx, application)

Delete an application

User must have the DELETE permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string**|  | 

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


## DeleteApplicationMember

> DeleteApplicationMember(ctx, application, user)

Remove an application member

User must have the MANAGE_MEMBERS permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string**|  | 
**user** | **string**|  | 

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


## Get2

> []map[string]interface{} Get2(ctx, application)

Get notification settings

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string**|  | 

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


## GetApplication

> ApplicationEntity GetApplication(ctx, application)

Get an application

User must have the READ permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string**|  | 

### Return type

[**ApplicationEntity**](ApplicationEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHooks1

> []map[string]interface{} GetHooks1(ctx, )

Get the list of available hooks

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


## GetNotifiers1

> []NotifierEntity GetNotifiers1(ctx, application)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string**|  | 

### Return type

[**[]NotifierEntity**](NotifierEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPermissions1

> []MemberEntity GetPermissions1(ctx, application)

Get application members

User must have the APPLICATION_MEMBER permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string**|  | 

### Return type

[**[]MemberEntity**](MemberEntity.md)

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


## Hits1

> Hits1(ctx, application, optional)

Get Application analytics

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string**|  | 
 **optional** | ***Hits1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a Hits1Opts struct


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


## List3

> []AlertEntity List3(ctx, application)

List configured alerts of a given APPLICATION

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string**|  | 

### Return type

[**[]AlertEntity**](AlertEntity.md)

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


## ListApplicationMembers

> []MembershipListItem ListApplicationMembers(ctx, application)

List application members

User must have the READ permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string**|  | 

### Return type

[**[]MembershipListItem**](MembershipListItem.md)

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


## ListApplications

> []ApplicationEntity ListApplications(ctx, optional)

List all the applications accessible to authenticated user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListApplicationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListApplicationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group** | **optional.String**|  | 
 **query** | **optional.String**|  | 

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


## TransferOwnership1

> TransferOwnership1(ctx, application, optional)

Transfer the ownership of the APPLICATION

User must have the TRANSFER_OWNERSHIP permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string**|  | 
 **optional** | ***TransferOwnership1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TransferOwnership1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of TransferOwnership**](TransferOwnership.md)|  | 

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


## Update6

> GenericNotificationConfigEntity Update6(ctx, application, notificationId, optional)

Update generic notification settings

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string**|  | 
**notificationId** | **string**|  | 
 **optional** | ***Update6Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a Update6Opts struct


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


## Update7

> PortalNotificationConfigEntity Update7(ctx, application, optional)

Update portal notification settings

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string**|  | 
 **optional** | ***Update7Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a Update7Opts struct


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


## Update8

> AlertEntity Update8(ctx, application, alert, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string**|  | 
**alert** | **string**|  | 
 **optional** | ***Update8Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a Update8Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of UpdateAlertEntity**](UpdateAlertEntity.md)|  | 

### Return type

[**AlertEntity**](AlertEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplication

> ApplicationEntity UpdateApplication(ctx, application, optional)

Update an application

User must have the MANAGE_APPLICATION permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string**|  | 
 **optional** | ***UpdateApplicationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateApplicationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpdateApplicationEntity**](UpdateApplicationEntity.md)|  | 

### Return type

[**ApplicationEntity**](ApplicationEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

