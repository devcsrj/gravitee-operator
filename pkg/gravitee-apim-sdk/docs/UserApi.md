# \UserApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUser**](UserApi.md#CreateUser) | **Post** /users | 
[**Delete19**](UserApi.md#Delete19) | **Delete** /user/notifications/{notification} | 
[**DeleteAll**](UserApi.md#DeleteAll) | **Delete** /user/notifications | 
[**DeleteUser**](UserApi.md#DeleteUser) | **Delete** /users/{id} | 
[**FindAll1**](UserApi.md#FindAll1) | **Get** /users | Search for API using the search engine
[**GetCurrentUser**](UserApi.md#GetCurrentUser) | **Get** /user | Get the authenticated user
[**GetCurrentUserPicture**](UserApi.md#GetCurrentUserPicture) | **Get** /user/avatar | Get user&#39;s avatar
[**GetGroups**](UserApi.md#GetGroups) | **Get** /users/{id}/groups | 
[**GetMemberships1**](UserApi.md#GetMemberships1) | **Get** /users/{id}/memberships | 
[**GetUser**](UserApi.md#GetUser) | **Get** /users/{id} | 
[**GetUserAvatar**](UserApi.md#GetUserAvatar) | **Get** /users/{id}/avatar | 
[**GetUserTasks**](UserApi.md#GetUserTasks) | **Get** /user/tasks | 
[**List13**](UserApi.md#List13) | **Get** /user/notifications | 
[**Login**](UserApi.md#Login) | **Post** /user/login | Login
[**Logout**](UserApi.md#Logout) | **Post** /user/logout | Logout
[**RegisterUser**](UserApi.md#RegisterUser) | **Post** /users/register | 
[**ResetPassword**](UserApi.md#ResetPassword) | **Post** /users/{id}/resetPassword | 
[**SearchUsers**](UserApi.md#SearchUsers) | **Get** /search/users | Search for users
[**UpdateCurrentUser**](UserApi.md#UpdateCurrentUser) | **Put** /user | Update user



## CreateUser

> CreateUser(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateUserOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateUserOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of RegisterUserEntity**](RegisterUserEntity.md)|  | 

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


## Delete19

> Delete19(ctx, notification)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notification** | **string**|  | 

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


## DeleteAll

> DeleteAll(ctx, )



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


## DeleteUser

> DeleteUser(ctx, id)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

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


## FindAll1

> UserEntity FindAll1(ctx, optional)

Search for API using the search engine

User must have the MANAGEMENT_USERS[READ] permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FindAll1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FindAll1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **optional.String**|  | 
 **size** | **optional.Int32**|  | [default to 20]
 **page** | **optional.Int32**|  | [default to 1]

### Return type

[**UserEntity**](UserEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentUser

> GetCurrentUser(ctx, )

Get the authenticated user

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


## GetCurrentUserPicture

> GetCurrentUserPicture(ctx, )

Get user's avatar

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


## GetGroups

> []UserGroupEntity GetGroups(ctx, id)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**[]UserGroupEntity**](UserGroupEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMemberships1

> UserMembershipList GetMemberships1(ctx, id, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
 **optional** | ***GetMemberships1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetMemberships1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **optional.String**|  | 

### Return type

[**UserMembershipList**](UserMembershipList.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> UserEntity GetUser(ctx, id)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**UserEntity**](UserEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserAvatar

> GetUserAvatar(ctx, id)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

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


## GetUserTasks

> PagedResult GetUserTasks(ctx, )



### Required Parameters

This endpoint does not need any parameter.

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


## List13

> PagedResultPortalNotificationEntity List13(ctx, )



### Required Parameters

This endpoint does not need any parameter.

### Return type

[**PagedResultPortalNotificationEntity**](PagedResultPortalNotificationEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Login

> Login(ctx, )

Login

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


## Logout

> Logout(ctx, )

Logout

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


## RegisterUser

> RegisterUser(ctx, optional)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RegisterUserOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RegisterUserOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of NewExternalUserEntity**](NewExternalUserEntity.md)|  | 

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


## ResetPassword

> ResetPassword(ctx, id)



### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

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


## SearchUsers

> []SearchableUser SearchUsers(ctx, q)

Search for users

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**q** | **string**|  | 

### Return type

[**[]SearchableUser**](SearchableUser.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCurrentUser

> UpdateCurrentUser(ctx, optional)

Update user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateCurrentUserOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateCurrentUserOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of UpdateUserEntity**](UpdateUserEntity.md)|  | 

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

