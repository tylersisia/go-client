# \TransactionsApi

All URIs are relative to *http://127.0.0.1:8070*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TransactionsAddressAddressStartHeightEndHeightGet**](TransactionsApi.md#TransactionsAddressAddressStartHeightEndHeightGet) | **Get** /transactions/address/{address}/{startHeight}/{endHeight} | Returns transactions for the wallet starting at start height until end height, that belong to the given address
[**TransactionsAddressAddressStartHeightGet**](TransactionsApi.md#TransactionsAddressAddressStartHeightGet) | **Get** /transactions/address/{address}/{startHeight} | Returns transactions for the wallet starting at start height for 1,000 blocks, that belong to the given address
[**TransactionsGet**](TransactionsApi.md#TransactionsGet) | **Get** /transactions | Gets a list of all transactions in the wallet container
[**TransactionsHashHashGet**](TransactionsApi.md#TransactionsHashHashGet) | **Get** /transactions/hash/{hash} | Gets details on the given transaction, if found
[**TransactionsPrepareAdvancedPost**](TransactionsApi.md#TransactionsPrepareAdvancedPost) | **Post** /transactions/prepare/advanced | Creates a transaction but does not relay it to the network
[**TransactionsPrepareBasicPost**](TransactionsApi.md#TransactionsPrepareBasicPost) | **Post** /transactions/prepare/basic | Creates a transaction but does not relay it to the network
[**TransactionsPreparedHashDelete**](TransactionsApi.md#TransactionsPreparedHashDelete) | **Delete** /transactions/prepared/{hash} | Cancels a previously prepared transaction
[**TransactionsPrivatekeyHashGet**](TransactionsApi.md#TransactionsPrivatekeyHashGet) | **Get** /transactions/privatekey/{hash} | Gets the transaction private key of the given transaction. This can be used to audit a transaction.
[**TransactionsSendAdvancedPost**](TransactionsApi.md#TransactionsSendAdvancedPost) | **Post** /transactions/send/advanced | Sends a transaction
[**TransactionsSendBasicPost**](TransactionsApi.md#TransactionsSendBasicPost) | **Post** /transactions/send/basic | Sends a transaction
[**TransactionsSendFusionAdvancedPost**](TransactionsApi.md#TransactionsSendFusionAdvancedPost) | **Post** /transactions/send/fusion/advanced | Sends a fusion transaction
[**TransactionsSendFusionBasicPost**](TransactionsApi.md#TransactionsSendFusionBasicPost) | **Post** /transactions/send/fusion/basic | Sends a fusion transaction
[**TransactionsSendPreparedPost**](TransactionsApi.md#TransactionsSendPreparedPost) | **Post** /transactions/send/prepared | Sends a previously prepared transaction
[**TransactionsStartHeightEndHeightGet**](TransactionsApi.md#TransactionsStartHeightEndHeightGet) | **Get** /transactions/{startHeight}/{endHeight} | Returns transactions for the wallet starting at start height until end height
[**TransactionsStartHeightGet**](TransactionsApi.md#TransactionsStartHeightGet) | **Get** /transactions/{startHeight} | Returns transactions for the wallet starting at start height for 1,000 blocks
[**TransactionsUnconfirmedAddressGet**](TransactionsApi.md#TransactionsUnconfirmedAddressGet) | **Get** /transactions/unconfirmed/{address} | Gets a list of unconfirmed, outgoing transactions, for the given address
[**TransactionsUnconfirmedGet**](TransactionsApi.md#TransactionsUnconfirmedGet) | **Get** /transactions/unconfirmed | Gets a list of all unconfirmed, outgoing transactions in the wallet container


# **TransactionsAddressAddressStartHeightEndHeightGet**
> InlineResponse2009 TransactionsAddressAddressStartHeightEndHeightGet(ctx, address, startHeight, endHeight)
Returns transactions for the wallet starting at start height until end height, that belong to the given address

Note that start height must be < end height. Also note that the transfers array will still contain transfers to other addresses, if present.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| The address to use for this operation. Should be a valid, 99 character TRTL address. | 
  **startHeight** | **float32**| The starting block height to use for this operation. | 
  **endHeight** | **float32**| The ending block height to use for this operation. | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TransactionsAddressAddressStartHeightGet**
> InlineResponse2009 TransactionsAddressAddressStartHeightGet(ctx, address, startHeight)
Returns transactions for the wallet starting at start height for 1,000 blocks, that belong to the given address

Note that the transfers array will still contain transfers to other addresses, if present

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| The address to use for this operation. Should be a valid, 99 character TRTL address. | 
  **startHeight** | **float32**| The starting block height to use for this operation. | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TransactionsGet**
> InlineResponse2009 TransactionsGet(ctx, )
Gets a list of all transactions in the wallet container

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TransactionsHashHashGet**
> InlineResponse20010 TransactionsHashHashGet(ctx, hash)
Gets details on the given transaction, if found

Note that the transaction must be contained in this wallet, and must not be unconfirmed. E.g, you must get this hash back when calling /transactions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hash** | **string**| The transaction hash to use for this operation. Should be a 64 char hex string. | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TransactionsPrepareAdvancedPost**
> InlineResponse2013 TransactionsPrepareAdvancedPost(ctx, body)
Creates a transaction but does not relay it to the network

Allows you to review the created transactions fee before deciding whether to commit to paying that fee. Prepared transactions can be sent using `/transactions/send/prepared`, or cancelled with `/transactions/prepared`. Note that every parameters sans destinations is optional.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SendTransactionAdvanced**](SendTransactionAdvanced.md)|  | 

### Return type

[**InlineResponse2013**](inline_response_201_3.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TransactionsPrepareBasicPost**
> InlineResponse2013 TransactionsPrepareBasicPost(ctx, body)
Creates a transaction but does not relay it to the network

This method will take funds from all subwallets as needed, and will use the primary address as the change address. It also uses a default fee, and default mixin. If this is not acceptable, please use the /advanced call instead. Allows you to review the created transactions fee before deciding whether to commit to paying that fee. Prepared transactions can be sent using `/transactions/send/prepared`, or cancelled with `/transactions/prepared`. Note that every parameters sans destinations is optional.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body6**](Body6.md)|  | 

### Return type

[**InlineResponse2013**](inline_response_201_3.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TransactionsPreparedHashDelete**
> TransactionsPreparedHashDelete(ctx, hash)
Cancels a previously prepared transaction

While it is not mandatory to call this method for a prepared transaction you do not wish to send, it is highly advised, as it will free up RAM.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hash** | **string**| The prepared transaction hash to cancel. This hash is returned from the &#x60;/transactions/prepared/basic&#x60; or &#x60;/transactions/prepared/advanced&#x60; methods. | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TransactionsPrivatekeyHashGet**
> InlineResponse20011 TransactionsPrivatekeyHashGet(ctx, hash)
Gets the transaction private key of the given transaction. This can be used to audit a transaction.

The transaction must have been sent by this wallet container. If the wallet container has been reimported, it will not be present.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hash** | **string**| The transaction hash to use for this operation. Should be a 64 char hex string. | 

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TransactionsSendAdvancedPost**
> InlineResponse2012 TransactionsSendAdvancedPost(ctx, body)
Sends a transaction

Note that every parameters sans destinations is optional.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SendTransactionAdvanced**](SendTransactionAdvanced.md)|  | 

### Return type

[**InlineResponse2012**](inline_response_201_2.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TransactionsSendBasicPost**
> InlineResponse2012 TransactionsSendBasicPost(ctx, body)
Sends a transaction

This method will take funds from all subwallets as needed, and will use the primary address as the change address. It also uses a default fee, and default mixin. If this is not acceptable, please use the /advanced call instead.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body5**](Body5.md)|  | 

### Return type

[**InlineResponse2012**](inline_response_201_2.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TransactionsSendFusionAdvancedPost**
> InlineResponse2015 TransactionsSendFusionAdvancedPost(ctx, body)
Sends a fusion transaction

Fusion transactions are zero fee, and seek to combine small inputs into larger ones, to allow for larger transactions. Many fusions may be required to fully optimize a wallet. Every parameter is optional.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body8**](Body8.md)|  | 

### Return type

[**InlineResponse2015**](inline_response_201_5.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TransactionsSendFusionBasicPost**
> InlineResponse2015 TransactionsSendFusionBasicPost(ctx, )
Sends a fusion transaction

Fusion transactions are zero fee, and seek to combine small inputs into larger ones, to allow for larger transactions. Many fusions may be required to fully optimize a wallet.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2015**](inline_response_201_5.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TransactionsSendPreparedPost**
> InlineResponse2014 TransactionsSendPreparedPost(ctx, body)
Sends a previously prepared transaction

Submits a transaction that was previously prepared with `/transactions/prepare/basic` or `/transactions/prepare/advanced` to the network.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body7**](Body7.md)|  | 

### Return type

[**InlineResponse2014**](inline_response_201_4.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TransactionsStartHeightEndHeightGet**
> InlineResponse2009 TransactionsStartHeightEndHeightGet(ctx, startHeight, endHeight)
Returns transactions for the wallet starting at start height until end height

Note that start height must be < end height

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **startHeight** | **float32**| The starting block height to use for this operation. | 
  **endHeight** | **float32**| The ending block height to use for this operation. | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TransactionsStartHeightGet**
> InlineResponse2009 TransactionsStartHeightGet(ctx, startHeight)
Returns transactions for the wallet starting at start height for 1,000 blocks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **startHeight** | **float32**| The starting block height to use for this operation. | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TransactionsUnconfirmedAddressGet**
> InlineResponse2009 TransactionsUnconfirmedAddressGet(ctx, address)
Gets a list of unconfirmed, outgoing transactions, for the given address

Note that this DOES NOT include incoming transactions in the pool. This only applies to transactions that have been sent by this wallet file, and have not been added to a block yet.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **address** | **string**| The address to use for this operation. Should be a valid, 99 character TRTL address. | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TransactionsUnconfirmedGet**
> InlineResponse2009 TransactionsUnconfirmedGet(ctx, )
Gets a list of all unconfirmed, outgoing transactions in the wallet container

Note that this DOES NOT include incoming transactions in the pool. This only applies to transactions that have been sent by this wallet file, and have not been added to a block yet.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

