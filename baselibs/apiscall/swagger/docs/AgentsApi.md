# {{classname}}

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgentsFreeNodeIdPost**](AgentsApi.md#AgentsFreeNodeIdPost) | **Post** /agents/free/{nodeId} | free node
[**AgentsInfoNodeIdGet**](AgentsApi.md#AgentsInfoNodeIdGet) | **Get** /agents/info/{nodeId} | get node info
[**AgentsJoinNodeIdPost**](AgentsApi.md#AgentsJoinNodeIdPost) | **Post** /agents/join/{nodeId} | join free node to a bag
[**AgentsListGet**](AgentsApi.md#AgentsListGet) | **Get** /agents/list | list nodes, return all node ids
[**AgentsUploadfilesPost**](AgentsApi.md#AgentsUploadfilesPost) | **Post** /agents/uploadfiles | upload files to nodes

# **AgentsFreeNodeIdPost**
> string AgentsFreeNodeIdPost(ctx, body, nodeId)
free node

free node

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApisNodeFreeReq**](ApisNodeFreeReq.md)| Node free request | 
  **nodeId** | **string**| node id | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentsInfoNodeIdGet**
> ApisNodeInfo AgentsInfoNodeIdGet(ctx, nodeId)
get node info

get node info by node id

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
> string AgentsJoinNodeIdPost(ctx, body, nodeId)
join free node to a bag

join free node to a bag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApisNodeJoinReq**](ApisNodeJoinReq.md)| Node join request | 
  **nodeId** | **string**| node id | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentsListGet**
> []string AgentsListGet(ctx, )
list nodes, return all node ids

### Required Parameters
This endpoint does not need any parameter.

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentsUploadfilesPost**
> AgentsUploadfilesPost(ctx, body)
upload files to nodes

upload files to nodes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApisUploadFilesReq**](ApisUploadFilesReq.md)| upload files request | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

