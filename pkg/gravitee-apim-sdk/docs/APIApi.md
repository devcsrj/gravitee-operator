# \APIApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOrUpdateApiMember**](APIApi.md#AddOrUpdateApiMember) | **Post** /apis/{api}/members | Add or update an API member
[**ApiLog**](APIApi.md#ApiLog) | **Get** /apis/{api}/logs/{log} | Get a specific log
[**ApiLogs**](APIApi.md#ApiLogs) | **Get** /apis/{api}/logs | Get API logs
[**ChangeSubscriptionStatus**](APIApi.md#ChangeSubscriptionStatus) | **Post** /apis/{api}/subscriptions/{subscription}/status | Change the status of a subscription
[**ClosePlan**](APIApi.md#ClosePlan) | **Post** /apis/{api}/plans/{plan}/_close | Close  a plan
[**Create**](APIApi.md#Create) | **Post** /apis/{api}/notificationsettings | Create notification settings
[**Create1**](APIApi.md#Create1) | **Post** /apis/{api}/metadata | Create an API metadata
[**Create2**](APIApi.md#Create2) | **Post** /apis/{api}/ratings | 
[**Create3**](APIApi.md#Create3) | **Post** /apis/{api}/alerts | 
[**Create4**](APIApi.md#Create4) | **Post** /apis/{api}/messages | 
[**CreateAnswer**](APIApi.md#CreateAnswer) | **Post** /apis/{api}/ratings/{rating}/answers | 
[**CreateApi**](APIApi.md#CreateApi) | **Post** /apis | Create an API
[**CreatePage**](APIApi.md#CreatePage) | **Post** /apis/{api}/pages | Create a page
[**CreatePlan**](APIApi.md#CreatePlan) | **Post** /apis/{api}/plans | Create a plan
[**CreateSubscription**](APIApi.md#CreateSubscription) | **Post** /apis/{api}/subscriptions | Subscribe to a plan
[**Delete**](APIApi.md#Delete) | **Delete** /apis/{api}/notificationsettings/{notificationId} | Delete notification settings
[**Delete1**](APIApi.md#Delete1) | **Delete** /apis/{api}/metadata/{metadata} | Delete a metadata
[**Delete2**](APIApi.md#Delete2) | **Delete** /apis/{api}/ratings/{rating} | 
[**Delete3**](APIApi.md#Delete3) | **Delete** /apis/{api}/ratings/{rating}/answers/{answer} | 
[**Delete4**](APIApi.md#Delete4) | **Delete** /apis/{api}/alerts/{alert} | 
[**Delete5**](APIApi.md#Delete5) | **Delete** /apis/{api} | Delete the API
[**DeleteApiMember**](APIApi.md#DeleteApiMember) | **Delete** /apis/{api}/members | Remove an API member
[**DeletePage**](APIApi.md#DeletePage) | **Delete** /apis/{api}/pages/{page} | Delete a page
[**DeletePlan**](APIApi.md#DeletePlan) | **Delete** /apis/{api}/plans/{plan} | Delete a plan
[**DeployAPI**](APIApi.md#DeployAPI) | **Post** /apis/{api}/deploy | Deploy API to gateway instances
[**DepreciatePlan**](APIApi.md#DepreciatePlan) | **Post** /apis/{api}/plans/{plan}/_depreciate | Depreciate a plan
[**DoLifecycleAction**](APIApi.md#DoLifecycleAction) | **Post** /apis/{api} | Manage the API&#39;s lifecycle
[**Events**](APIApi.md#Events) | **Get** /apis/{api}/events | Get API&#39;s events
[**ExportDefinition**](APIApi.md#ExportDefinition) | **Get** /apis/{api}/export | Export the API definition in JSON format
[**FetchAllPages**](APIApi.md#FetchAllPages) | **Post** /apis/{api}/pages/_fetch | Refresh all pages by calling their associated fetcher
[**FetchPage**](APIApi.md#FetchPage) | **Post** /apis/{api}/pages/{page}/_fetch | Refresh page by calling the associated fetcher
[**Get**](APIApi.md#Get) | **Get** /apis/{api}/notificationsettings | Get notification settings
[**Get1**](APIApi.md#Get1) | **Get** /apis/{api} | Get the API definition
[**GetApiMetadata**](APIApi.md#GetApiMetadata) | **Get** /apis/{api}/metadata/{metadata} | A metadata for the given API and metadata id
[**GetApiSubscription**](APIApi.md#GetApiSubscription) | **Get** /apis/{api}/subscriptions/{subscription} | Get a subscription
[**GetByApiAndUser**](APIApi.md#GetByApiAndUser) | **Get** /apis/{api}/ratings/current | 
[**GetEvents**](APIApi.md#GetEvents) | **Get** /apis/{api}/audit/events | 
[**GetHeaders**](APIApi.md#GetHeaders) | **Get** /apis/{api}/headers | Get the portal API headers values
[**GetHooks**](APIApi.md#GetHooks) | **Get** /apis/hooks | Get the list of available hooks
[**GetImage**](APIApi.md#GetImage) | **Get** /apis/{api}/media/{hash} | 
[**GetNotifiers**](APIApi.md#GetNotifiers) | **Get** /apis/{api}/notifiers | 
[**GetPage**](APIApi.md#GetPage) | **Get** /apis/{api}/pages/{page} | Get a page
[**GetPageContent**](APIApi.md#GetPageContent) | **Get** /apis/{api}/pages/{page}/content | Get the page&#39;s content
[**GetPermissions**](APIApi.md#GetPermissions) | **Get** /apis/{api}/members/permissions | Get API members
[**GetPlan**](APIApi.md#GetPlan) | **Get** /apis/{api}/plans/{plan} | Get a plan
[**GetQualityMetrics**](APIApi.md#GetQualityMetrics) | **Get** /apis/{api}/quality | Get the quality metrics of the API
[**GetSummaryByApi**](APIApi.md#GetSummaryByApi) | **Get** /apis/{api}/ratings/summary | 
[**Health**](APIApi.md#Health) | **Get** /apis/{api}/health | Health-check statistics for API
[**HealthAverage**](APIApi.md#HealthAverage) | **Get** /apis/{api}/health/average | Health-check average statistics for API
[**HealthcheckLog**](APIApi.md#HealthcheckLog) | **Get** /apis/{api}/health/logs/{log} | Health-check log
[**HealthcheckLogs**](APIApi.md#HealthcheckLogs) | **Get** /apis/{api}/health/logs | Health-check logs
[**Hits**](APIApi.md#Hits) | **Get** /apis/{api}/analytics | Get API analytics
[**ImportDefinition**](APIApi.md#ImportDefinition) | **Post** /apis/import | Create an API by importing an API definition
[**ImportPathMappingsFromPage**](APIApi.md#ImportPathMappingsFromPage) | **Post** /apis/{api}/import-path-mappings | Import path mappings from a page
[**ImportSwagger**](APIApi.md#ImportSwagger) | **Post** /apis/import/swagger | Create an API definition from a Swagger descriptor
[**IsAPISynchronized**](APIApi.md#IsAPISynchronized) | **Get** /apis/{api}/state | Get the state of the API
[**List**](APIApi.md#List) | **Get** /apis/{api}/ratings | 
[**List1**](APIApi.md#List1) | **Get** /apis/{api}/audit | 
[**List2**](APIApi.md#List2) | **Get** /apis/{api}/alerts | List configured alerts of a given API
[**ListApiKeysForSubscription**](APIApi.md#ListApiKeysForSubscription) | **Get** /apis/{api}/subscriptions/{subscription}/keys | List all API Keys for a subscription
[**ListApiMembers**](APIApi.md#ListApiMembers) | **Get** /apis/{api}/members | List API members
[**ListApiMetadatas**](APIApi.md#ListApiMetadatas) | **Get** /apis/{api}/metadata | List metadata for the given API
[**ListApiSubscribers**](APIApi.md#ListApiSubscribers) | **Get** /apis/{api}/subscribers | List subscribers for the API
[**ListApiSubscriptions**](APIApi.md#ListApiSubscriptions) | **Get** /apis/{api}/subscriptions | List subscriptions for the API
[**ListApis**](APIApi.md#ListApis) | **Get** /apis | List APIs
[**ListPages**](APIApi.md#ListPages) | **Get** /apis/{api}/pages | List pages
[**ListPlans**](APIApi.md#ListPlans) | **Get** /apis/{api}/plans | List plans for an API
[**PartialUpdatePage**](APIApi.md#PartialUpdatePage) | **Patch** /apis/{api}/pages/{page} | Update a page
[**Picture**](APIApi.md#Picture) | **Get** /apis/{api}/picture | Get the API&#39;s picture
[**ProcessApiSubscription**](APIApi.md#ProcessApiSubscription) | **Post** /apis/{api}/subscriptions/{subscription}/_process | Update a subscription
[**PublishPlan**](APIApi.md#PublishPlan) | **Post** /apis/{api}/plans/{plan}/_publish | Publicly publish plan
[**RenewApiKey**](APIApi.md#RenewApiKey) | **Post** /apis/{api}/subscriptions/{subscription} | Renew an API key
[**RevokeApiKey**](APIApi.md#RevokeApiKey) | **Delete** /apis/{api}/keys/{key} | 
[**RevokeApiKey1**](APIApi.md#RevokeApiKey1) | **Delete** /apis/{api}/subscriptions/{subscription}/keys/{key} | Revoke an API key
[**Rollback**](APIApi.md#Rollback) | **Post** /apis/{api}/rollback | Rollback API to a previous version
[**SearchApis**](APIApi.md#SearchApis) | **Post** /apis/_search | Search for API using the search engine
[**SearchPortalApis**](APIApi.md#SearchPortalApis) | **Post** /portal/apis/_search | Search for API using the search engine
[**TransferOwnership**](APIApi.md#TransferOwnership) | **Post** /apis/{api}/members/transfer_ownership | Transfer the ownership of the API
[**Update**](APIApi.md#Update) | **Put** /apis/{api}/notificationsettings/{notificationId} | Update generic notification settings
[**Update1**](APIApi.md#Update1) | **Put** /apis/{api}/notificationsettings | Update portal notification settings
[**Update2**](APIApi.md#Update2) | **Put** /apis/{api}/metadata/{metadata} | Update an API metadata
[**Update3**](APIApi.md#Update3) | **Put** /apis/{api}/ratings/{rating} | 
[**Update4**](APIApi.md#Update4) | **Put** /apis/{api}/alerts/{alert} | 
[**Update5**](APIApi.md#Update5) | **Put** /apis/{api} | Update the API
[**UpdateApiKey**](APIApi.md#UpdateApiKey) | **Put** /apis/{api}/keys/{key} | Update an API Key
[**UpdateApiSubscription**](APIApi.md#UpdateApiSubscription) | **Put** /apis/{api}/subscriptions/{subscription} | Update a subscription
[**UpdatePage**](APIApi.md#UpdatePage) | **Put** /apis/{api}/pages/{page} | Update a page
[**UpdatePlan**](APIApi.md#UpdatePlan) | **Put** /apis/{api}/plans/{plan} | Update a plan
[**UpdateWithDefinition**](APIApi.md#UpdateWithDefinition) | **Post** /apis/{api}/import | Update the API with an existing API definition
[**UploadImage**](APIApi.md#UploadImage) | **Post** /apis/{api}/media/upload | Create a picture for an API
[**Verify**](APIApi.md#Verify) | **Post** /apis/verify | Check if an API match the following criteria



## AddOrUpdateApiMember

> AddOrUpdateApiMember(ctx, api, optional)

Add or update an API member

User must have the MANAGE_MEMBERS permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
 **optional** | ***AddOrUpdateApiMemberOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddOrUpdateApiMemberOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ApiMembership**](ApiMembership.md)|  | 

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


## ApiLog

> ApiLog(ctx, api, log, optional)

Get a specific log

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
**log** | **string**|  | 
 **optional** | ***ApiLogOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApiLogOpts struct


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


## ApiLogs

> ApiLogs(ctx, api, optional)

Get API logs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
 **optional** | ***ApiLogsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ApiLogsOpts struct


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


## ClosePlan

> PlanEntity ClosePlan(ctx, api, plan)

Close  a plan

User must have the MANAGE_PLANS permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
**plan** | **string**|  | 

### Return type

[**PlanEntity**](PlanEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Create

> map[string]interface{} Create(ctx, api, optional)

Create notification settings

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
 **optional** | ***CreateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateOpts struct


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


## Create3

> AlertEntity Create3(ctx, api, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
 **optional** | ***Create3Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a Create3Opts struct


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


## Create4

> Create4(ctx, api, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
 **optional** | ***Create4Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a Create4Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of MessageEntity**](MessageEntity.md)|  | 

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


## CreateApi

> CreateApi(ctx, api)

Create an API

User must have API_PUBLISHER or ADMIN role to create an API.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | [**NewApiEntity**](NewApiEntity.md)|  | 

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


## CreatePage

> PageEntity CreatePage(ctx, api, page)

Create a page

User must have the MANAGE_PAGES permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
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


## CreatePlan

> PlanEntity CreatePlan(ctx, api, plan)

Create a plan

User must have the MANAGE_PLANS permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
**plan** | [**NewPlanEntity**](NewPlanEntity.md)|  | 

### Return type

[**PlanEntity**](PlanEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: application/json
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


## Delete

> Delete(ctx, api, notificationId)

Delete notification settings

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
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


## Delete4

> Delete4(ctx, api, alert)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
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


## Delete5

> Delete5(ctx, api)

Delete the API

User must have the DELETE permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 

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


## DeleteApiMember

> DeleteApiMember(ctx, api, user)

Remove an API member

User must have the MANAGE_MEMBERS permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
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


## DeletePage

> DeletePage(ctx, api, page)

Delete a page

User must have the MANAGE_PAGES permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
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


## DeletePlan

> DeletePlan(ctx, api, plan)

Delete a plan

User must have the MANAGE_PLANS permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
**plan** | **string**|  | 

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


## DeployAPI

> ApiEntity DeployAPI(ctx, api)

Deploy API to gateway instances

User must have the MANAGE_LIFECYCLE permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 

### Return type

[**ApiEntity**](ApiEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DepreciatePlan

> PlanEntity DepreciatePlan(ctx, api, plan)

Depreciate a plan

User must have the API_PLAN[UPDATE] permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
**plan** | **string**|  | 

### Return type

[**PlanEntity**](PlanEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DoLifecycleAction

> DoLifecycleAction(ctx, action, api)

Manage the API's lifecycle

User must have the MANAGE_LIFECYCLE permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**action** | **string**|  | 
**api** | **string**|  | 

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


## Events

> Events(ctx, api, optional)

Get API's events

User must have the MANAGE_LIFECYCLE permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
 **optional** | ***EventsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EventsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **optional.String**|  | [default to all]

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


## ExportDefinition

> ApiEntity ExportDefinition(ctx, api, optional)

Export the API definition in JSON format

User must have the MANAGE_APPLICATION permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
 **optional** | ***ExportDefinitionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ExportDefinitionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **optional.String**|  | [default to default]
 **exclude** | **optional.String**|  | 

### Return type

[**ApiEntity**](ApiEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAllPages

> PageEntity FetchAllPages(ctx, api)

Refresh all pages by calling their associated fetcher

User must have the MANAGE_PAGES permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 

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


## FetchPage

> PageEntity FetchPage(ctx, api, page)

Refresh page by calling the associated fetcher

User must have the MANAGE_PAGES permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
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


## Get

> []map[string]interface{} Get(ctx, api)

Get notification settings

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 

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


## Get1

> ApiEntity Get1(ctx, api)

Get the API definition

User must have the READ permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 

### Return type

[**ApiEntity**](ApiEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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


## GetEvents

> GetEvents(ctx, )



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


## GetHeaders

> []ApiHeaderEntity GetHeaders(ctx, api)

Get the portal API headers values

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 

### Return type

[**[]ApiHeaderEntity**](ApiHeaderEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHooks

> []map[string]interface{} GetHooks(ctx, )

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


## GetImage

> GetImage(ctx, api, hash)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
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


## GetNotifiers

> []NotifierEntity GetNotifiers(ctx, api)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 

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


## GetPage

> GetPage(ctx, api, page, optional)

Get a page

User must have the READ permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
**page** | **string**|  | 
 **optional** | ***GetPageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **portal** | **optional.Bool**|  | 

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


## GetPageContent

> GetPageContent(ctx, api, page)

Get the page's content

User must have the READ permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
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


## GetPermissions

> []MemberEntity GetPermissions(ctx, api)

Get API members

User must have the MANAGE_MEMBERS permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 

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


## GetPlan

> PlanEntity GetPlan(ctx, api, plan)

Get a plan

User must have the READ permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
**plan** | **string**|  | 

### Return type

[**PlanEntity**](PlanEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQualityMetrics

> ApiQualityMetricsEntity GetQualityMetrics(ctx, api)

Get the quality metrics of the API

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 

### Return type

[**ApiQualityMetricsEntity**](ApiQualityMetricsEntity.md)

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


## Health

> Health(ctx, api, optional)

Health-check statistics for API

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
 **optional** | ***HealthOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a HealthOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **optional.String**|  | [default to availability]
 **field** | **optional.String**|  | [default to endpoint]

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


## HealthAverage

> HealthAverage(ctx, api, type_, optional)

Health-check average statistics for API

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
**type_** | **string**|  | 
 **optional** | ***HealthAverageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a HealthAverageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **from** | **optional.Int64**|  | 
 **to** | **optional.Int64**|  | 
 **interval** | **optional.Int64**|  | 

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


## HealthcheckLog

> HealthcheckLog(ctx, api, log)

Health-check log

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
**log** | **string**|  | 

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


## HealthcheckLogs

> HealthcheckLogs(ctx, api, optional)

Health-check logs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
 **optional** | ***HealthcheckLogsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a HealthcheckLogsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **query** | **optional.String**|  | 
 **size** | **optional.Int32**|  | [default to 10]
 **page** | **optional.Int32**|  | [default to 1]
 **transition** | **optional.Bool**|  | 

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


## Hits

> Hits(ctx, api, optional)

Get API analytics

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
 **optional** | ***HitsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a HitsOpts struct


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


## ImportDefinition

> ImportDefinition(ctx, definition)

Create an API by importing an API definition

Create an API by importing an existing API definition in JSON format

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definition** | **string**|  | 

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


## ImportPathMappingsFromPage

> ApiEntity ImportPathMappingsFromPage(ctx, api, page)

Import path mappings from a page

User must have the MANAGE_APPLICATION permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
**page** | **string**|  | 

### Return type

[**ApiEntity**](ApiEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportSwagger

> NewApiEntity ImportSwagger(ctx, swagger)

Create an API definition from a Swagger descriptor

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**swagger** | [**ImportSwaggerDescriptorEntity**](ImportSwaggerDescriptorEntity.md)|  | 

### Return type

[**NewApiEntity**](NewApiEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IsAPISynchronized

> ApiEntity IsAPISynchronized(ctx, api)

Get the state of the API

User must have the MANAGE_LIFECYCLE permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 

### Return type

[**ApiEntity**](ApiEntity.md)

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


## List1

> MetadataPageAuditEntity List1(ctx, api, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
 **optional** | ***List1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a List1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mgmt** | **optional.Bool**|  | 
 **api2** | **optional.String**|  | 
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


## List2

> []AlertEntity List2(ctx, api)

List configured alerts of a given API

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 

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


## ListApiMembers

> []MembershipListItem ListApiMembers(ctx, api)

List API members

User must have the MANAGE_MEMBERS permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 

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


## ListApis

> []ApiListItem ListApis(ctx, optional)

List APIs

List all the APIs accessible to the current user or only public APIs for non authenticated users.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListApisOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListApisOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **view** | **optional.String**|  | 
 **group** | **optional.String**|  | 
 **top** | **optional.Bool**|  | 
 **contextPath** | **optional.String**|  | 
 **owner** | **optional.String**|  | 
 **label** | **optional.String**|  | 
 **state** | **optional.String**|  | 
 **visibility** | **optional.String**|  | 
 **version** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **tag** | **optional.String**|  | 

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


## ListPages

> []PageEntity ListPages(ctx, api, optional)

List pages

User must have the READ permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
 **optional** | ***ListPagesOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListPagesOpts struct


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


## ListPlans

> []PlanEntity ListPlans(ctx, api, optional)

List plans for an API

List all the plans accessible to the current user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
 **optional** | ***ListPlansOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListPlansOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **optional.String**|  | [default to published]

### Return type

[**[]PlanEntity**](PlanEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdatePage

> PageEntity PartialUpdatePage(ctx, api, page, optional)

Update a page

User must have the MANAGE_PAGES permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
**page** | **string**|  | 
 **optional** | ***PartialUpdatePageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PartialUpdatePageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | [**optional.Interface of UpdatePageEntity**](UpdatePageEntity.md)|  | 

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


## Picture

> Picture(ctx, api)

Get the API's picture

User must have the READ permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 

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


## PublishPlan

> PlanEntity PublishPlan(ctx, api, plan)

Publicly publish plan

User must have the MANAGE_PLANS permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
**plan** | **string**|  | 

### Return type

[**PlanEntity**](PlanEntity.md)

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


## RevokeApiKey

> RevokeApiKey(ctx, api, key)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
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


## Rollback

> ApiEntity Rollback(ctx, api, api)

Rollback API to a previous version

User must have the MANAGE_LIFECYCLE permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
**api** | [**UpdateApiEntity**](UpdateApiEntity.md)|  | 

### Return type

[**ApiEntity**](ApiEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchApis

> []ApiListItem SearchApis(ctx, q)

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


## TransferOwnership

> TransferOwnership(ctx, api, optional)

Transfer the ownership of the API

User must have the TRANSFER_OWNERSHIP permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
 **optional** | ***TransferOwnershipOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TransferOwnershipOpts struct


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


## Update

> GenericNotificationConfigEntity Update(ctx, api, notificationId, optional)

Update generic notification settings

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
**notificationId** | **string**|  | 
 **optional** | ***UpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateOpts struct


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


## Update1

> PortalNotificationConfigEntity Update1(ctx, api, optional)

Update portal notification settings

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
 **optional** | ***Update1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a Update1Opts struct


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


## Update4

> AlertEntity Update4(ctx, api, alert, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
**alert** | **string**|  | 
 **optional** | ***Update4Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a Update4Opts struct


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


## Update5

> ApiEntity Update5(ctx, api, api)

Update the API

User must have the MANAGE_APPLICATION permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
**api** | [**UpdateApiEntity**](UpdateApiEntity.md)|  | 

### Return type

[**ApiEntity**](ApiEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApiKey

> ApiKeyEntity UpdateApiKey(ctx, api, key, optional)

Update an API Key

User must have the MANAGE_API_KEYS permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
**key** | **string**|  | 
 **optional** | ***UpdateApiKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateApiKeyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ApiKeyEntity**](ApiKeyEntity.md)|  | 

### Return type

[**ApiKeyEntity**](ApiKeyEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

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


## UpdatePage

> PageEntity UpdatePage(ctx, api, page, page)

Update a page

User must have the MANAGE_PAGES permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
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


## UpdatePlan

> PlanEntity UpdatePlan(ctx, api, plan, plan)

Update a plan

User must have the MANAGE_PLANS permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
**plan** | **string**|  | 
**plan** | [**UpdatePlanEntity**](UpdatePlanEntity.md)|  | 

### Return type

[**PlanEntity**](PlanEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWithDefinition

> ApiEntity UpdateWithDefinition(ctx, api, definition)

Update the API with an existing API definition

User must have the MANAGE_APPLICATION permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
**definition** | **string**|  | 

### Return type

[**ApiEntity**](ApiEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadImage

> PageEntity UploadImage(ctx, api, optional)

Create a picture for an API

User must have the API_DOCUMENTATION permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**api** | **string**|  | 
 **optional** | ***UploadImageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UploadImageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | **optional.Interface of *os.File****optional.*os.File**|  | 

### Return type

[**PageEntity**](PageEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Verify

> Verify(ctx, optional)

Check if an API match the following criteria

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VerifyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VerifyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of VerifyApiParam**](VerifyApiParam.md)|  | 

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

