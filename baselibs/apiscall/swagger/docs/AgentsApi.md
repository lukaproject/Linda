# {{classname}}

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgentsFreeNodeIdPost**](AgentsApi.md#AgentsFreeNodeIdPost) | **Post** /agents/free/{nodeId} | free node
[**AgentsInfoNodeIdGet**](AgentsApi.md#AgentsInfoNodeIdGet) | **Get** /agents/info/{nodeId} | get node info
[**AgentsJoinNodeIdPost**](AgentsApi.md#AgentsJoinNodeIdPost) | **Post** /agents/join/{nodeId} | join free node to a bag
[**AgentsListGet**](AgentsApi.md#AgentsListGet) | **Get** /agents/list | list nodes, return node infos by query
[**AgentsListidsGet**](AgentsApi.md#AgentsListidsGet) | **Get** /agents/listids | list nodes, return node ids by query
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
> []ApisNodeInfo AgentsListGet(ctx, optional)
list nodes, return node infos by query

list nodes, return node infos by query, query format support prefix=, createAfter=, idAfter=, limit=.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AgentsApiAgentsListGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgentsApiAgentsListGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **prefix** | **optional.String**| find all infos with this prefix | 
 **createAfter** | **optional.Int32**| find all infos created after this time (ms) | 
 **limit** | **optional.Int32**| max count of node infos in result | 
 **idAfter** | **optional.String**| find all node infos which id greater or equal to this id | 

### Return type

[**[]ApisNodeInfo**](apis.NodeInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AgentsListidsGet**
> []string AgentsListidsGet(ctx, optional)
list nodes, return node ids by query

list nodes, return node ids by query, query format support prefix=, createAfter=, idAfter=, limit=.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AgentsApiAgentsListidsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgentsApiAgentsListidsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **prefix** | **optional.String**| find all ids with this prefix | 
 **createAfter** | **optional.Int32**| find all ids created after this time (ms) | 
 **limit** | **optional.Int32**| max count of node ids in result | 
 **idAfter** | **optional.String**| find all node ids which id greater or equal to this id | 

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

