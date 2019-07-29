# GoChain Filesystem CLI

```sh
> gofs help
NAME:
   gofs - GoChain filesystem cli tool

USAGE:
   gofs [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
     pin       Pin a CID
     rate      Get the current storage rate in atto GO per ByteHour.
     cost      Get the current storage cost in atto GO for the given size and duration.
     add       Add and pin a file.
     status    Get the current storage status for a CID.
     receipts  Query for receipts.
     help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --url value       GOFS API URL. (default: "https://gofs.io/api/v0/") [$GOFS_API]
   --contract value  Contract address. (default: "0x0000000000000000000000000000000000001234") [$GOFS_CONTRACT]
   --rpc value       RPC URL. (default: "https://rpc.gochain.io") [$GOFS_RPC]
   --help, -h        show help
   --version, -v     print the version
```
