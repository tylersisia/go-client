# InlineResponse2005

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletBlockCount** | **int32** | The amount of blocks the wallet has scanned | [optional] [default to null]
**LocalDaemonBlockCount** | **int32** | The amount of blocks the daemon the wallet is connected to has synced | [optional] [default to null]
**NetworkBlockCount** | **int32** | The amount of blocks the network has | [optional] [default to null]
**PeerCount** | **int32** | The amount of peers (incoming + outgoing) peers the daemon has | [optional] [default to null]
**Hashrate** | **int32** | The hashrate the last local block the daemon knows about has | [optional] [default to null]
**IsViewWallet** | **bool** | Whether this wallet is a view only wallet. Certain operations are illegal on a view only wallet, such as transferring. | [optional] [default to null]
**SubWalletCount** | **int32** | The amount of subwallets in the container | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


