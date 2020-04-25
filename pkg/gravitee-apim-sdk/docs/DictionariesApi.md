# \DictionariesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDictionary**](DictionariesApi.md#CreateDictionary) | **Post** /configuration/dictionaries | Create a dictionary
[**DeleteDictionary**](DictionariesApi.md#DeleteDictionary) | **Delete** /configuration/dictionaries/{dictionary} | Delete a dictionary
[**DeployDictionary**](DictionariesApi.md#DeployDictionary) | **Post** /configuration/dictionaries/{dictionary}/_deploy | Deploy dictionary to API gateway
[**DoLifecycleAction1**](DictionariesApi.md#DoLifecycleAction1) | **Post** /configuration/dictionaries/{dictionary} | Manage the dictionary&#39;s lifecycle
[**GetDictionary**](DictionariesApi.md#GetDictionary) | **Get** /configuration/dictionaries/{dictionary} | Get a dictionary
[**ListDictionaries**](DictionariesApi.md#ListDictionaries) | **Get** /configuration/dictionaries | Get the list of global dictionaries
[**UndeployDictionary**](DictionariesApi.md#UndeployDictionary) | **Post** /configuration/dictionaries/{dictionary}/_undeploy | Undeploy dictionary to API gateway
[**UpdateDictionary**](DictionariesApi.md#UpdateDictionary) | **Put** /configuration/dictionaries/{dictionary} | Update a dictionary



## CreateDictionary

> DictionaryEntity CreateDictionary(ctx, dictionary)

Create a dictionary

User must have the DICTIONARY[CREATE] permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dictionary** | [**NewDictionaryEntity**](NewDictionaryEntity.md)|  | 

### Return type

[**DictionaryEntity**](DictionaryEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDictionary

> DeleteDictionary(ctx, dictionary)

Delete a dictionary

User must have the DICTIONARY[DELETE] permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dictionary** | **string**|  | 

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


## DeployDictionary

> DictionaryEntity DeployDictionary(ctx, dictionary)

Deploy dictionary to API gateway

User must have the DICTIONARY[UPDATE] permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dictionary** | **string**|  | 

### Return type

[**DictionaryEntity**](DictionaryEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DoLifecycleAction1

> DoLifecycleAction1(ctx, action, dictionary)

Manage the dictionary's lifecycle

User must have the DICTIONARY[LIFECYCLE] permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**action** | **string**|  | 
**dictionary** | **string**|  | 

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


## GetDictionary

> GetDictionary(ctx, dictionary)

Get a dictionary

User must have the DICTIONARY[READ] permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dictionary** | **string**|  | 

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


## ListDictionaries

> []DictionaryListItem ListDictionaries(ctx, )

Get the list of global dictionaries

User must have the DICTIONARY[READ] permission to use this service

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]DictionaryListItem**](DictionaryListItem.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UndeployDictionary

> DictionaryEntity UndeployDictionary(ctx, dictionary)

Undeploy dictionary to API gateway

User must have the DICTIONARY[UPDATE] permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dictionary** | **string**|  | 

### Return type

[**DictionaryEntity**](DictionaryEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDictionary

> DictionaryEntity UpdateDictionary(ctx, dictionary, dictionary)

Update a dictionary

User must have the DICTIONARY[UPDATE] permission to use this service

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dictionary** | **string**|  | 
**dictionary** | [**UpdateDictionaryEntity**](UpdateDictionaryEntity.md)|  | 

### Return type

[**DictionaryEntity**](DictionaryEntity.md)

### Authorization

[gravitee-auth](../README.md#gravitee-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

