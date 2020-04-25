# \IdentityProvidersApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIdentityProvider**](IdentityProvidersApi.md#CreateIdentityProvider) | **Post** /configuration/identities | Create an identity provider
[**DeleteIdentityProvider**](IdentityProvidersApi.md#DeleteIdentityProvider) | **Delete** /configuration/identities/{identityProvider} | Delete an identity provider
[**GetIdentityProvider**](IdentityProvidersApi.md#GetIdentityProvider) | **Get** /configuration/identities/{identityProvider} | Get an identity provider
[**ListIdentityProviders**](IdentityProvidersApi.md#ListIdentityProviders) | **Get** /configuration/identities | Get the list of identity providers
[**ListSocialIdentityProvider**](IdentityProvidersApi.md#ListSocialIdentityProvider) | **Get** /portal/identities | Get the list of social identity providers
[**UpdateIdentityProvider**](IdentityProvidersApi.md#UpdateIdentityProvider) | **Put** /configuration/identities/{identityProvider} | Update an identity provider



## CreateIdentityProvider

> IdentityProviderEntity CreateIdentityProvider(ctx, identityProvider)

Create an identity provider

User must have the PORTAL_IDENTITY_PROVIDER[CREATE] permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityProvider** | [**NewIdentityProviderEntity**](NewIdentityProviderEntity.md)|  | 

### Return type

[**IdentityProviderEntity**](IdentityProviderEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdentityProvider

> DeleteIdentityProvider(ctx, identityProvider)

Delete an identity provider

User must have the IDENTITY_PROVIDER[DELETE] permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityProvider** | **string**|  | 

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


## GetIdentityProvider

> GetIdentityProvider(ctx, identityProvider)

Get an identity provider

User must have the IDENTITY_PROVIDER[READ] permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityProvider** | **string**|  | 

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


## ListIdentityProviders

> []IdentityProviderListItem ListIdentityProviders(ctx, )

Get the list of identity providers

User must have the PORTAL_IDENTITY_PROVIDER[READ] permission to use this service

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]IdentityProviderListItem**](IdentityProviderListItem.md)

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


## UpdateIdentityProvider

> IdentityProviderEntity UpdateIdentityProvider(ctx, identityProvider, dictionary)

Update an identity provider

User must have the IDENTITY_PROVIDER[UPDATE] permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**identityProvider** | **string**|  | 
**dictionary** | [**UpdateIdentityProviderEntity**](UpdateIdentityProviderEntity.md)|  | 

### Return type

[**IdentityProviderEntity**](IdentityProviderEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

