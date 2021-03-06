# SMRepl - A Spacemesh Repl

## Overview

A basic reference Spacemesh Wallet and a full node manager command line REPL. 
Designed to work together with a locally running [go-spacemesh full node](https://github.com/spacemeshos/go-spacemesh) or with the public Spacemesh API. 
Target users are developers who want to hack on the Spacemesh platform and Spacemesh testnet users who are comfortable with terminal apps and repls. 
For non-terminal savvy users we recommend using [Smapp - the Spacemesh App](https://github.com/spacemeshos/smapp). Smapp is available for all major desktop platforms and
includes a wallet and a node manager.

## Features

SMRepl is a Spacemesh API client and it implements basic wallet features via a REPL interface. Using it you can create a new coin account, execute transactions, check account balance and view transactions. You can also use it to se tup smeshing (Spacemesh mining) for your local node.
You can also use it when working on Spacemesh code by connecting it to a locally built node or to a [local Spacemesh devnet](https://github.com/spacemeshos/local-testnet)
You can also connect to a public Spacemesh network API. In this configuraiton you can execute transactions, view account balances, past transactions and smeshing rewards.

## Building

### Build for your current platform with go:

```bash
go get && go build
```

### Build for all platforms:

```bash
make
```

### Build for a specific platforms:

```bash
make build-win
```

```bash
make build-mac
```

```bash
make build-linux
```

### With `docker`:

```
make dockerbuild-go
```

---

## CLI Flags

Use `-wallet_directory` to override the default of current working directory when opening and creating wallets.

Use `-wallet` to specify a wallet to pre-open when starting cli-wallet. cli-wallet will look in current directory
unless `-wallet_directory` has been specified.

## Using with a public Spacemesh API server

You can use your wallet without running a full node by connecting it to a public Spacemesh api service for a Spacemesh
network. Use the `-grpc-server` and `-secure` flags connect to a remote Spacemesh api server. For example:

```bash
./cli_wallet_darwin_amd64 -server api-123.spacemesh.io:443 -secure
```

> Note that communications with the server will be secure using https/tls but the wallet doesn't currently verify the server identity.

Your can find the grpc server url of a public Spacemesh testnet by copying the value of the `grpcAPI` field from the [Spacemesh discovery srvice data](https://discover.spacemesh.io/networks.json).


## Using with a local Spacemesh full node

1. Join a Spacemesh network by running [go-spacemesh](https://github.com/spacemeshos/go-spacemesh/releases)
   or [Smapp](https://github.com/spacemeshos/smapp/releases) on your computer.
1. Build the wallet from this repository and run it. For example on OS X:

```bash
make build-mac
./cli_wallet_darwin_amd64
```

By default, the wallet attempts to connect to the api server provided by your locally running Spacemesh full node using
the default node's grpc api port (localhost:9092). When you run your full node directly in terminal, you can configure
which api services will be available to your wallet by your node by changing entries int he api section of your node's
config file:

```json
{
   "api" : {
      "grpc": "node, mesh, globalstate, transaction, smesher"
   }
}
```
