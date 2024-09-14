# {{classname}}

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgentsFreeNodeIdPost**](DefaultApi.md#AgentsFreeNodeIdPost) | **Post** /agents/free/{nodeId} | free node
[**AgentsInfoNodeIdGet**](DefaultApi.md#AgentsInfoNodeIdGet) | **Get** /agents/info/{nodeId} | get node info
[**AgentsJoinNodeIdPost**](DefaultApi.md#AgentsJoinNodeIdPost) | **Post** /agents/join/{nodeId} | join free node to a bag
[**BagnodesBagNameGet**](DefaultApi.md#BagnodesBagNameGet) | **Get** /bagnodes/{bagName} | list bag nodes [no implementation]
[**BagsBagNameDelete**](DefaultApi.md#BagsBagNameDelete) | **Delete** /bags/{bagName} | delete bag
[**BagsBagNameGet**](DefaultApi.md#BagsBagNameGet) | **Get** /bags/{bagName} | get bag
[**BagsBagNameTasksPost**](DefaultApi.md#BagsBagNameTasksPost) | **Post** /bags/{bagName}/tasks | add task
[**BagsBagNameTasksTaskNameGet**](DefaultApi.md#BagsBagNameTasksTaskNameGet) | **Get** /bags/{bagName}/tasks/{taskName} | get task
[**BagsGet**](DefaultApi.md#BagsGet) | **Get** /bags | list bag [no implementation]
[**BagsPost**](DefaultApi.md#BagsPost) | **Post** /bags | add bag
[**HealthcheckPost**](DefaultApi.md#HealthcheckPost) | **Post** /healthcheck | health check

# **AgentsFreeNodeIdPost**
> ApisNodeFreeResp AgentsFreeNodeIdPost(ctx, body, nodeId)
free node

free node

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApisNodeFreeReq**](ApisNodeFreeReq.md)| Node free request | 
  **nodeId** | **string**| node id | 

### Return type

[**ApisNodeFreeResp**](apis.NodeFreeResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentsInfoNodeIdGet**
> ApisNodeInfo AgentsInfoNodeIdGet(ctx, nodeId)
get node info

get node info

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**| node id | 

### Return type

[**ApisNodeInfo**](apis.NodeInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentsJoinNodeIdPost**
> ApisNodeJoinResp AgentsJoinNodeIdPost(ctx, body, nodeId)
join free node to a bag

join free node to a bag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApisNodeJoinReq**](ApisNodeJoinReq.md)| Node join request | 
  **nodeId** | **string**| node id | 

### Return type

[**ApisNodeJoinResp**](apis.NodeJoinResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BagnodesBagNameGet**
> ApisListBagNodesResp BagnodesBagNameGet(ctx, )
list bag nodes [no implementation]

list all node ids which belong to this node

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApisListBagNodesResp**](apis.ListBagNodesResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BagsBagNameDelete**
> ApisDeleteBagResp BagsBagNameDelete(ctx, bagName)
delete bag

delete bag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bagName** | **string**| bag&#x27;s name | 

### Return type

[**ApisDeleteBagResp**](apis.DeleteBagResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BagsBagNameGet**
> ApisGetBagResp BagsBagNameGet(ctx, bagName)
get bag

get bag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bagName** | **string**| bag&#x27;s name | 

### Return type

[**ApisGetBagResp**](apis.GetBagResp.md)

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

# **BagsGet**
> ApisListBagsResp BagsGet(ctx, )
list bag [no implementation]

list bag

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApisListBagsResp**](apis.ListBagsResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BagsPost**
> ApisAddBagResp BagsPost(ctx, body)
add bag

create a new bag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApisAddBagReq**](ApisAddBagReq.md)| bag&#x27;s request | 

### Return type

[**ApisAddBagResp**](apis.AddBagResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HealthcheckPost**
> HealthcheckPost(ctx, )
health check

health check

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

