# \WalletApi

All URIs are relative to *http://127.0.0.1:8070*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WalletCreatePost**](WalletApi.md#WalletCreatePost) | **Post** /wallet/create | Creates a new wallet
[**WalletDelete**](WalletApi.md#WalletDelete) | **Delete** /wallet | Closes and saves the opened wallet
[**WalletImportKeyPost**](WalletApi.md#WalletImportKeyPost) | **Post** /wallet/import/key | Imports a wallet with a private spend and view key
[**WalletImportSeedPost**](WalletApi.md#WalletImportSeedPost) | **Post** /wallet/import/seed | Imports a wallet using a mnemonic seed
[**WalletImportViewPost**](WalletApi.md#WalletImportViewPost) | **Post** /wallet/import/view | Imports a view only wallet with a private view key and public address
[**WalletOpenPost**](WalletApi.md#WalletOpenPost) | **Post** /wallet/open | Opens an already existing wallet


# **WalletCreatePost**
> WalletCreatePost(ctx, body)
Creates a new wallet

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Wallet**](Wallet.md)|  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WalletDelete**
> WalletDelete(ctx, )
Closes and saves the opened wallet

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WalletImportKeyPost**
> WalletImportKeyPost(ctx, body)
Imports a wallet with a private spend and view key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**WalletKeyImport**](WalletKeyImport.md)|  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WalletImportSeedPost**
> WalletImportSeedPost(ctx, body)
Imports a wallet using a mnemonic seed

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**WalletSeedImport**](WalletSeedImport.md)|  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WalletImportViewPost**
> WalletImportViewPost(ctx, body)
Imports a view only wallet with a private view key and public address

Note that view only wallets can only see incoming transactions, so balance may be inflated, and they cannot send transactions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**WalletViewImport**](WalletViewImport.md)|  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WalletOpenPost**
> WalletOpenPost(ctx, body)
Opens an already existing wallet

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Wallet**](Wallet.md)|  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

