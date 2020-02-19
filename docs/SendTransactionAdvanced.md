# SendTransactionAdvanced

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destinations** | [**[]SendTransactionAdvancedDestinations**](SendTransactionAdvanced_destinations.md) |  | [default to null]
**Mixin** | **float32** | The mixin level to use | [optional] [default to null]
**Fee** | **float32** | The fee to use with this transaction (in atomic units) | [optional] [default to null]
**FeePerByte** | **float64** | The amount in atomic units to pay for each byte of the resulting transaction size | [optional] [default to 1.953125]
**SourceAddresses** | **[]string** |  | [optional] [default to null]
**PaymentID** | **string** | The payment ID to use | [optional] [default to null]
**ChangeAddress** | **string** | The address in this wallet to return any &#39;change&#39; to if we have to spend more than the requested amount. Defaults to primary address. | [optional] [default to null]
**UnlockTime** | **float32** | When to unlock the transaction. A user cannot spend locked funds until the unlock time has been reached. Defaults to zero if not given. Can use either a block height, or a unix timestamp. | [optional] [default to null]
**Extra** | **string** | Hex representation of any extra data to be included in the &#x60;tx_extra&#x60; field of the transaction. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


