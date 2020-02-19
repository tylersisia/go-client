# \BalanceApi

All URIs are relative to *http://127.0.0.1:8070*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BalanceAddressGet**](BalanceApi.md#BalanceAddressGet) | **Get** /balance/{address} | Get the balance for a specific address
[**BalanceGet**](BalanceApi.md#BalanceGet) | **Get** /balance | Get the balance for the entire wallet container
[**BalancesGet**](BalanceApi.md#BalancesGet) | **Get** /balances | Get the balance for every address


# **BalanceAddressGet**
> InlineResponse20012 BalanceAddressGet(ctx, address)
Get the balance for a specific address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| The address to use for this operation. Should be a valid, 99 character TRTL address. | 

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BalanceGet**
> InlineResponse20012 BalanceGet(ctx, )
Get the balance for the entire wallet container

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BalancesGet**
> []InlineResponse20013 BalancesGet(ctx, )
Get the balance for every address

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]InlineResponse20013**](inline_response_200_13.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

