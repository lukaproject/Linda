# {{classname}}

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BagsBagNameTasksPost**](TasksApi.md#BagsBagNameTasksPost) | **Post** /bags/{bagName}/tasks | add task
[**BagsBagNameTasksTaskNameGet**](TasksApi.md#BagsBagNameTasksTaskNameGet) | **Get** /bags/{bagName}/tasks/{taskName} | get task

# **BagsBagNameTasksPost**
> ApisAddTaskResp BagsBagNameTasksPost(ctx, body, bagName)
add task

add task

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApisAddTaskReq**](ApisAddTaskReq.md)| add tasks&#x27;s request | 
  **bagName** | **string**| bag&#x27;s name | 

### Return type

[**ApisAddTaskResp**](apis.AddTaskResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BagsBagNameTasksTaskNameGet**
> ApisGetTaskResp BagsBagNameTasksTaskNameGet(ctx, bagName, taskName)
get task

get task

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bagName** | **string**| bag&#x27;s name | 
  **taskName** | **string**| task&#x27;s name | 

### Return type

[**ApisGetTaskResp**](apis.GetTaskResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

