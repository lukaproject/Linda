# {{classname}}

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BagnodesBagNameGet**](BagsApi.md#BagnodesBagNameGet) | **Get** /bagnodes/{bagName} | list bag nodes
[**BagsBagNameDelete**](BagsApi.md#BagsBagNameDelete) | **Delete** /bags/{bagName} | delete bag
[**BagsBagNameGet**](BagsApi.md#BagsBagNameGet) | **Get** /bags/{bagName} | get bag
[**BagsGet**](BagsApi.md#BagsGet) | **Get** /bags | list bags [no implementation]
[**BagsPost**](BagsApi.md#BagsPost) | **Post** /bags | add bag

# **BagnodesBagNameGet**
> ApisListBagNodesResp BagnodesBagNameGet(ctx, )
list bag nodes

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

# **BagsGet**
> ApisListBagsResp BagsGet(ctx, )
list bags [no implementation]

list bags

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

