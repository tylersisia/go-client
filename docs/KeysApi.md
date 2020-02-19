# \KeysApi

All URIs are relative to *http://127.0.0.1:8070*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KeysAddressGet**](KeysApi.md#KeysAddressGet) | **Get** /keys/{address} | Gets the public and private spend key for the given address
[**KeysGet**](KeysApi.md#KeysGet) | **Get** /keys | Gets the wallet containers shared private view key
[**KeysMnemonicAddressGet**](KeysApi.md#KeysMnemonicAddressGet) | **Get** /keys/mnemonic/{address} | Gets the mnemonic seed for the given address, if possible


# **KeysAddressGet**
> InlineResponse2003 KeysAddressGet(ctx, address)
Gets the public and private spend key for the given address

Note that this method cannot be used with a view only wallet

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| The address to use for this operation. Should be a valid, 99 character TRTL address. | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **KeysGet**
> InlineResponse2002 KeysGet(ctx, )
Gets the wallet containers shared private view key

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **KeysMnemonicAddressGet**
> InlineResponse2004 KeysMnemonicAddressGet(ctx, address)
Gets the mnemonic seed for the given address, if possible

Note that this method cannot be used with a view only wallet

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| The address to use for this operation. Should be a valid, 99 character TRTL address. | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

