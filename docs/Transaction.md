# Transaction

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockHeight** | **float32** | The block this transaction is contained in. | [optional] [default to null]
**Fee** | **float32** | The amount the sender paid in miner fees for this transaction. In atomic units. | [optional] [default to null]
**Hash** | **string** | The hash of this transaction | [optional] [default to null]
**IsCoinbaseTransaction** | **bool** | Whether this transaction is a miner reward or a normal transaction. | [optional] [default to null]
**PaymentID** | **string** | An identifier supplied by the sender. May be \&quot;\&quot; (empty string) | [optional] [default to null]
**Timestamp** | **float32** | The unix timestamp of the block this transaction is contained in. | [optional] [default to null]
**UnlockTime** | **float32** | When this transaction unlocks for spending. If &gt;&#x3D; 50000000, treated as a timestamp. Normally zero. | [optional] [default to null]
**Transfers** | [***TransactionTransfers**](Transaction_transfers.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


