# \PlanApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClosePlan**](PlanApi.md#ClosePlan) | **Post** /apis/{api}/plans/{plan}/_close | Close  a plan
[**CreatePlan**](PlanApi.md#CreatePlan) | **Post** /apis/{api}/plans | Create a plan
[**DeletePlan**](PlanApi.md#DeletePlan) | **Delete** /apis/{api}/plans/{plan} | Delete a plan
[**DepreciatePlan**](PlanApi.md#DepreciatePlan) | **Post** /apis/{api}/plans/{plan}/_depreciate | Depreciate a plan
[**GetPlan**](PlanApi.md#GetPlan) | **Get** /apis/{api}/plans/{plan} | Get a plan
[**ListPlans**](PlanApi.md#ListPlans) | **Get** /apis/{api}/plans | List plans for an API
[**PublishPlan**](PlanApi.md#PublishPlan) | **Post** /apis/{api}/plans/{plan}/_publish | Publicly publish plan
[**UpdatePlan**](PlanApi.md#UpdatePlan) | **Put** /apis/{api}/plans/{plan} | Update a plan



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

