# \AddressesApi

All URIs are relative to *http://127.0.0.1:8070*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddressesAddressDelete**](AddressesApi.md#AddressesAddressDelete) | **Delete** /addresses/{address} | Deletes the given subwallet from the container
[**AddressesAddressPaymentIDGet**](AddressesApi.md#AddressesAddressPaymentIDGet) | **Get** /addresses/{address}/{paymentID} | Creates an integrated address from an address and payment ID
[**AddressesCreatePost**](AddressesApi.md#AddressesCreatePost) | **Post** /addresses/create | Creates a new, random address in the wallet container
[**AddressesGet**](AddressesApi.md#AddressesGet) | **Get** /addresses | Gets a list of all addresses in the wallet container
[**AddressesImportPost**](AddressesApi.md#AddressesImportPost) | **Post** /addresses/import | Imports a subwallet with the given private spend key
[**AddressesImportViewPost**](AddressesApi.md#AddressesImportViewPost) | **Post** /addresses/import/view | Imports a view only subwallet with the given publicSpendKey
[**AddressesPrimaryGet**](AddressesApi.md#AddressesPrimaryGet) | **Get** /addresses/primary | Gets the &#39;primary&#39; address


# **AddressesAddressDelete**
> AddressesAddressDelete(ctx, address)
Deletes the given subwallet from the container

Note that you cannot delete the 'primary' address, the first address created in the wallet.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| The address to use for this operation. Should be a valid, 99 character TRTL address. | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddressesAddressPaymentIDGet**
> InlineResponse2008 AddressesAddressPaymentIDGet(ctx, address, paymentID)
Creates an integrated address from an address and payment ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| The address to use for this operation. Should be a valid, 99 character TRTL address. | 
  **paymentID** | **string**| The payment ID to use for this operation. Should be a 64 char hex string. | 

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddressesCreatePost**
> InlineResponse201 AddressesCreatePost(ctx, )
Creates a new, random address in the wallet container

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse201**](inline_response_201.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddressesGet**
> InlineResponse2006 AddressesGet(ctx, )
Gets a list of all addresses in the wallet container

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddressesImportPost**
> InlineResponse2011 AddressesImportPost(ctx, body)
Imports a subwallet with the given private spend key

It is HIGHLY recommended you provide a scan height with this operation - wallet syncing will have to begin again from the scan height given (defaults to zero) if the scan height is less than the height of the current wallet sync status.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body3**](Body3.md)|  | 

### Return type

[**InlineResponse2011**](inline_response_201_1.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddressesImportViewPost**
> InlineResponse2011 AddressesImportViewPost(ctx, body)
Imports a view only subwallet with the given publicSpendKey

It is HIGHLY recommended you provide a scan height with this operation - wallet syncing will have to begin again from the scan height given (defaults to zero) if the scan height is less than the height of the current wallet sync status.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body4**](Body4.md)|  | 

### Return type

[**InlineResponse2011**](inline_response_201_1.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddressesPrimaryGet**
> InlineResponse2007 AddressesPrimaryGet(ctx, )
Gets the 'primary' address

The primary address is the first wallet created, and the one used as the change address if not specified.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

