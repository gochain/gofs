# GOFS CLI

The `gofs` command line interface provides access to both the contract and the web api.

## Installation

Quickest:

```shell script
> go install github.com/gochain/gofs/cmd/gofs
```

To build with proper version info:

```shell script
> git clone https://github.com/gochain/gofs.git
> cd gofs && make install
```

## How to Use

Run `gofs [command] --help` for more details about each command.

```shell script
> gofs --help
NAME:
   gofs - GoChain filesystem cli tool

USAGE:
   gofs [global options] command [command options] [arguments...]

VERSION:
   v0.7.0

COMMANDS:
   pin       Pin a new CID or extend the expiry of an existing file.
   wallet    Get the deposit wallet for the CID.
   rate      Get the current storage rate in attoGO per byte-hour.
   cost      Get the current storage cost in attoGO for the given size and duration.
   add       Add and pin a file.
   status    Get the current storage status for a CID.
   receipts  Query for receipts.
   help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --url value       GOFS API URL. (default: "https://api.gofs.io") [$GOFS_API]
   --contract value  Hex contract address. (default: "0x0000000000000000000000000000000000001234") [$GOFS_CONTRACT]
   --rpc value       RPC URL. (default: "https://rpc.gochain.io") [$GOFS_RPC]
   --help, -h        show help
   --version, -v     print the version
```

## Example Session

Here is an example session using the `gofs` command, after setting `WEB3_PRIVATE_KEY`.

```shell script
# Get the current rate from the contract.
> gofs rate
Current storage rate: 2837942 attoGO per byte-hour.

Cost:
	 1.00KB for 30d: 0.000002043318240000 GO
	 1.00MB for 30d: 0.002043318240000000 GO
	 1.00GB for 30d: 2.043318240000000000 GO
	 1.00KB for 1y: 0.000024860371920000 GO
	 1.00MB for 1y: 0.024860371920000000 GO
	 1.00GB for 1y: 24.860371920000000000 GO

# Calculate the cost of storing a 1.3KB file for 1 year. 
> gofs cost -s 1.3KB -e 1y
  11419200 bytes-hours: 0.000032407027286400 GO

# Upload and pin a file to GOFS.
> gofs add file.txt
File uploaded and pinned.
CID: bafkreicysg23kiwv34eg2d7qweipxwosdo2py4ldv42nbauguluen5v6am
Pinned until: 2019-09-04 15:37:42 -0500 CDT
File size: 6B

# Get the current status of the file from GOFS.
> gofs status bafkreicysg23kiwv34eg2d7qweipxwosdo2py4ldv42nbauguluen5v6am
File size: 6B
Expires in 58m1s at 2019-09-04 15:37:42 -0500 CDT.

# Pay to pin a file for 1 year, by calling the contract (with WEB3_PRIVATE_KEY).
> gofs pin --extend 1y bafkreicysg23kiwv34eg2d7qweipxwosdo2py4ldv42nbauguluen5v6am
Purchased 1237865 byte-hours of storage for bafkreicysg23kiwv34eg2d7qweipxwosdo2py4ldv42nbauguluen5v6am.
Tx: 0x2b923b6271a24dcf50eb9efc12627c60876e0b69ea068dd0bda629fcac8e78bc

# Get recent receipts for the file from the contract.
> gofs receipts --from 7400000 --to 7461960 --cids bafkreicysg23kiwv34eg2d7qweipxwosdo2py4ldv42nbauguluen5v6am
Block   Tx  Log Removed CID                                               BH   User                                       
7461944 186 0   false   zb2rhcc1wJn2GHDLT2YkmPq5b69cXc2xfRZZmyufbjFUfBkxr 1000 0x7A2772Edb801670450021e0d6Cd35606c9875fA5

# Check for a deposit wallet, and create if none exists (with WEB3_PRIVATE_KEY).
> gofs wallet bafkreicysg23kiwv34eg2d7qweipxwosdo2py4ldv42nbauguluen5v6am
No deposit wallet exists for this CID. Create one? Y/N
> y
Created new wallet - Tx: 0xeb33cb8f51278e11d837b411da12a64d0cab8fd6aa2688dd0fdf3ae204f3147d
Deposit wallet: 0x2510e1DD247C37F4cf852bEeCc94A6B9Df05eE50

# Use web3 to send a standard tx to the deposit wallet (with WEB3_PRIVATE_KEY).
> web3 send 0.01 to 0x2510e1DD247C37F4cf852bEeCc94A6B9Df05eE50 --gas
Transaction address: 0x9d889e70f7fed252e6123a8d56fe99a7723c8f18122a43b77efce720aee71b1a

# Get recent receipts for the file from the contract.
> gofs receipts --from 7461000 --to 7462090 --cids bafkreicysg23kiwv34eg2d7qweipxwosdo2py4ldv42nbauguluen5v6am
Block   Tx  Log Removed CID                                               BH         User                                       
7461944 186 0   false   zb2rhcc1wJn2GHDLT2YkmPq5b69cXc2xfRZZmyufbjFUfBkxr 1000       0x7A2772Edb801670450021e0d6Cd35606c9875fA5 
7462085 26  0   false   zb2rhcc1wJn2GHDLT2YkmPq5b69cXc2xfRZZmyufbjFUfBkxr 3523680187 0x7A2772Edb801670450021e0d6Cd35606c9875fA5
```
