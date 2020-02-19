# Body8

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mixin** | **float32** | The mixin level to use | [optional] [default to null]
**SourceAddresses** | **[]string** | The addresses to draw funds for the transaction from (must be an address in this container) | [optional] [default to null]
**Destination** | **string** | The destination address to send funds to. Must exist in this wallet. | [optional] [default to null]
**OptimizeTarget** | **float32** | If given, we will not fuse inputs larger than this value. Value given must be a valid input amount, i.e. only a single significant leading digit. For example, 20000 is fine, 23456 is not. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


