# {{classname}}

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FilesDownloadBlockFileNameGet**](FilesApi.md#FilesDownloadBlockFileNameGet) | **Get** /files/download/{block}/{fileName} | download file
[**FilesUploadPost**](FilesApi.md#FilesUploadPost) | **Post** /files/upload | Upload file

# **FilesDownloadBlockFileNameGet**
> FilesDownloadBlockFileNameGet(ctx, fileName, block)
download file

download file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileName** | **string**| file name | 
  **block** | **string**| block name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FilesUploadPost**
> string FilesUploadPost(ctx, fileName, block, file)
Upload file

Upload file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileName** | **string**|  | 
  **block** | **string**|  | 
  **file** | ***os.File*****os.File**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

