# GoFS - GoChain FileSystem - https://gofs.io

GoFS is a pay-to-pin IPFS service built on GoChain.
Payments processed on the blockchain fund the storage of files on GoFS to ensure the 
files are always available on IPFS. There are three ways to interface with GoFS: the web interface, the CLI and the API. 
You can find information about all 3 below.

## How to Use

There are two primary functions that you'll use when interacting with GoFS: adding a file and paying for storage.

### 1. Add: Upload a file
 
Add a new file to IPFS by uploading it to GoFS. This can be done through the
[web interface](https://gofs.io) or the [JSON API](#Add). New files are 
initially pinned for a grace period of one hour.

### 2. Pin: Pay to extend the life of a file

Pay to pin a file on IPFS. Payments are made on the blockchain to the GoFS
smart contract. This can be done through the [web interface](https://gofs.io/)
with MetaMask, on the command line with the [gofs cli](./cmd/gofs), or 
programmatically against the [contract itself](#Contract-ABI). Each payment
purchases storage for a particular CID (measured in _byte-hours_). The amount of 
storage credited is calculated based on the contract rate (measured in _attoGo/byte-hour_) 
and this value is included on the emitted `Pinned` event. When GoFS processes these 
events, the storage amount and the file size determine how much to extend the expiration.
The `Pinned` events emitted by the GoFS contract serve as a public, auditable trail of receipts.

#### File Wallets

File wallets are addresses which can receive standard txs to extend the life of a particular file.
They are mini-contracts which only contain a fallback function to forward payment to the
GoFS contract. This removes barriers for users created by complex smart contract interaction,
while still utilizing the same underlying mechanisms, so the same kind of `Pinned` events are 
emitted when payments come through a wallet.

## Limitations

- Size Limit: 1GB - Larger files will be rejected.
- Minimum Pin Duration: 1 month - Unpinned files will not be fetched, unless 1 month has been funded.
- Recursive Directories: Not yet supported - Coming soon.
- CID: v0 not allowed, v1 required, base32 encoding preferred - Contract will reject v0.

## Web Interface

The [web interface](https://gofs.io) is the most user friendly way to use GoFS, and supports MetaMask integration.

## CLI

The `gofs` [command line interface](./cmd/gofs) provides access to both the contract and the web api.

```shell script
> gofs status bafkreicysg23kiwv34eg2d7qweipxwosdo2py4ldv42nbauguluen5v6am
File size: 6B
Expires in 58m1s at 2019-09-04 15:37:42 -0500 CDT.
```

## JSON API

The JSON API used by the web interface is available at: `https://api.gofs.io/v0/`

### Info

Get general information about GoFS.

#### Request:

[`GET /info`](https://api.gofs.io/v0/info)

#### Response:

An object with the following fields:

- rate: current storage rate in attoGo per byte-hour
- contract_address: address of the GoFS contract
- explorer_url: block explorer URL for the GoChain network
- rpc_url: RPC URL for the GoChain network
- max_file_size: maximum supported file size, in bytes

Example:

```json
{
  "rate": 2837942,
  "contract_address": "0xded28050fdbf604e12056e516c05e154cb5dd1bc",
  "explorer_url": "https:\/\/explorer.gochain.io\/",
  "rpc_url": "https:\/\/rpc.gochain.io",
  "max_file_size": 1000000000
}
```

### Status

Get the status for a file by IPFS CID hash.

#### Request:

[`GET /status/{cid}`](https://api.gofs.io/v0/status/bafkreicysg23kiwv34eg2d7qweipxwosdo2py4ldv42nbauguluen5v6am)

#### Response:

`HTTP 404`, if never added to GoFS. Otherwise, an object with the following fields:

- expiration: unix timestamp of file expiration
- size: file size, in bytes

Example:

```json
{
  "expiration": 1567619841,
  "size": 6737
}
```

#### Try it

```sh
curl https://api.gofs.io/v0/status/bafybeida2gv6hnykvwadda6zjcxuo5yrxzb6z7j7fhi2mz7o55carn6jla
```

### Add

Add a file to IPFS by uploading to GoFS. 
Maximum file size is limited (see `max_file_size` from [/info](#info)).
Files are initially pinned for a grace period of one hour.
This can be extended by calling `pin()` on the GoFS [smart contract](#Contract-ABI).
See [How to Use](#How-to-Use) for other methods.

#### Request:

`PUT /add`

Body: File data, or a multi-part form file. 

#### Response:

An object with the following fields:
- cid: IPFS CID hash for file
- expiration: unix timestamp of file expiration
- size: file size, in bytes

Example:

```json
{
  "cid": "bafkreic62jyg5yvckkumrnsqo43wfltlao4khbbf4mtj3if7hrbxbmikya",
  "expiration": 1567619841,
  "size": 6737
}
```

#### Try it

```sh
curl -X PUT -T myfile.png https://api.gofs.io/v0/add
```

### Convert CID

Convert an IPFS CID hash to various standard forms: binary, default base, and event topic hash. 

#### Request:

[`GET /convert-cid/{cid}`](https://api.gofs.io/v0/convert-cid/bafkreicysg23kiwv34eg2d7qweipxwosdo2py4ldv42nbauguluen5v6am)

#### Response:

An object with the following fields:
- binary: hex encoded binary CID hash (call `Pinned` with this)
- base: default encoded (base32) CID hash
- hash: hex encoded keccak256 of the binary CID (filter `Pinned` event topics with this)
- version: CID version

Example:

```json
{
  "binary": "0x015512205ed2706ee2a252a8c8b650773762ae6b03b8a38425e3269da0bf3c4370b10ac0",
  "base": "bafkreic62jyg5yvckkumrnsqo43wfltlao4khbbf4mtj3if7hrbxbmikya",
  "hash": "0x60632b18db19d0f6d10a0f7dcf0eea38e8114eb867f34252c1f2c6ff148dc557",
  "version": 1
}
```

## Contract ABI

If you want to interact with the contract directly, you can do that too. The GOFS contract is at 0xTODO, and implements the [contracts/IGOFS.abi](contracts/IGOFS.abi) interface:

```solidity
// The IGOFS interface defines the public functions for GOFS.
interface IGOFS {
    // Returns the current rate in attoGO per byte-hour.
    function rate() external view returns (uint);

    // Pin a CID. Value (the GO value for your transaction) must be greater than 0. The GO value you pass in
    // here is how long the file will be pinned for. CID must not be version 0.
    // Emits Pinned events.
    function pin(bytes calldata cid) external payable;

    // Get the address of the deposit wallet for a cid. Returns 0x0 if none exists.
    function wallet(bytes calldata cid) external view returns (address);

    // Create a deposit wallet for a cid. Returns false if one already exists. CID must not be version 0.
    // Emits CreatedWallet events.
    // Uses <=300000 gas.
    function newWallet(bytes calldata cid) external;

    // Emitted when a cid is pinned.
    event Pinned(address indexed user, bytes indexed cid, uint bh);
    // Emitted when a new wallet is created.
    event CreatedWallet(address indexed user, bytes indexed cid, address wallet);
    ...
}
```
