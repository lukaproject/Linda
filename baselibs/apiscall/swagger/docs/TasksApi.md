# {{classname}}

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BagsBagNameTasksGet**](TasksApi.md#BagsBagNameTasksGet) | **Get** /bags/{bagName}/tasks | list tasks
[**BagsBagNameTasksPost**](TasksApi.md#BagsBagNameTasksPost) | **Post** /bags/{bagName}/tasks | add task
[**BagsBagNameTasksTaskNameGet**](TasksApi.md#BagsBagNameTasksTaskNameGet) | **Get** /bags/{bagName}/tasks/{taskName} | get task

# **BagsBagNameTasksGet**
> []ApisTask BagsBagNameTasksGet(ctx, bagName, optional)
list tasks

list tasks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bagName** | **string**| bag&#x27;s name | 
 **optional** | ***TasksApiBagsBagNameTasksGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TasksApiBagsBagNameTasksGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **prefix** | **optional.String**| find all tasks which taskName with this prefix | 
 **createAfter** | **optional.Int32**| find all tasks created after this time (ms) | 
 **limit** | **optional.Int32**| max count of tasks in result | 
 **idAfter** | **optional.String**| find all tasks which taskName greater or equal to this id | 

### Return type

[**[]ApisTask**](apis.Task.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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

