# YAY GAMES

## Problem #1. Get Merkle Proof.

```Bash
$ go run ./cmd/merkleproof/main.go

> 2021/10/23 03:05:45 Reading data from source file...
> 2021/10/23 03:05:45 Hashing initial leaves...
> 2021/10/23 03:05:45 Building Merkle Tree from leaves...
> 2021/10/23 03:05:45 Generated Proof: 
 [0xddeba7e8a21e9093f1f9126f657f0d5436d49d5d94fd97f06a785ee82746cb83,
  0x34f820d6e5262d6b19a5c5ef998f8ef59b75c5f9a5bbc16faad5171c0758e0a4,
  0x51e4f91f8d74ce4c620f9d9b6e02e028e57b90e72e8806cf23b647beed468e4e,
  0x4d3f3ee35a5803e6defb94a3e7c581c36124baf82f3e146e5e1e73142f1f5936,
  0x15b68bbc0ab9dba042a1dce13ff9e8c6de24b5f3f1cbda633bde9ac6148cb4ec,
  0x9ebbe33359ae58cffb7e098bfa14dde4c627de813d3dc60cc637b26d23931e2e]
> 2021/10/23 03:05:45 Checking claim via smart contract...
> 2021/10/23 03:05:46 Huray! Valid proof.
```

## Problem #2. ExecuteRound Daemon.

```Bash
$ go run ./cmd/daemon/main.go

> 2021/10/23 02:59:08 ExecuteRound Daemon Started:
> 2021/10/23 02:59:11 ——————————————————————
> 2021/10/23 02:59:11 Start Block: 5951302
> 2021/10/23 02:59:11 Current Block: 5979635
> 2021/10/23 02:59:11 Next Round Border: 5980102
> 2021/10/23 02:59:11 Blocks Left: 467
> 2021/10/23 02:59:14 ——————————————————————
> 2021/10/23 02:59:14 Start Block: 5951302
> 2021/10/23 02:59:14 Current Block: 5979636
> 2021/10/23 02:59:14 Next Round Border: 5980102
> 2021/10/23 02:59:14 Blocks Left: 466
> 2021/10/23 02:59:14 .......

$ ^C
> C2021/10/23 02:54:01 OS exit signal received: interrupt
> 2021/10/23 02:54:01 starting graceful shutdown: <nil>
> 2021/10/23 02:54:01 context canceled
> exit status 1
```

<!-- 
// abigen --abi=./src/BtcVsEth.abi --pkg=contract --out=./pkg/contract/contract.go
// ./build/avalanchego --network-id=fuji
// $GOPATH/pkg/mod/github.com/ava-labs/avalanchego@v1.6.3/build/avalanchego --network-id=fuji
-->