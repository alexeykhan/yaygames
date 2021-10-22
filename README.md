# YAY GAMES

## Problem #1. Get Merkle Proof.

### Step 1. Generate Merkle Proof.

```Bash
$ go run ./cmd/merkleproof/main.go

> [0xddeba7e8a21e9093f1f9126f657f0d5436d49d5d94fd97f06a785ee82746cb83,
  0x34f820d6e5262d6b19a5c5ef998f8ef59b75c5f9a5bbc16faad5171c0758e0a4,
  0x51e4f91f8d74ce4c620f9d9b6e02e028e57b90e72e8806cf23b647beed468e4e,
  0x4d3f3ee35a5803e6defb94a3e7c581c36124baf82f3e146e5e1e73142f1f5936,
  0x15b68bbc0ab9dba042a1dce13ff9e8c6de24b5f3f1cbda633bde9ac6148cb4ec,
  0x9ebbe33359ae58cffb7e098bfa14dde4c627de813d3dc60cc637b26d23931e2e]
```

### Step 2. Check Claim.

1. Open [Read Contract](https://testnet.bscscan.com/address/0xCd71C520CB6358D3B455f1d018813FF667001618#readContract)


2. Enter data for first array element:
   - `_target (address):` 0x42e66ECe2b29B3c08B5D6013634484a69303caC0
   - `_category (uint256):` 1
   - `_amount (uint256):` 65870135
   - `_merkleProof (bytes32[]):` [0xddeba7e8a21e9093f1f9126f657f0d5436d49d5d94fd97f06a785ee82746cb83,0x34f820d6e5262d6b19a5c5ef998f8ef59b75c5f9a5bbc16faad5171c0758e0a4,0x51e4f91f8d74ce4c620f9d9b6e02e028e57b90e72e8806cf23b647beed468e4e,0x4d3f3ee35a5803e6defb94a3e7c581c36124baf82f3e146e5e1e73142f1f5936,0x15b68bbc0ab9dba042a1dce13ff9e8c6de24b5f3f1cbda633bde9ac6148cb4ec,0x9ebbe33359ae58cffb7e098bfa14dde4c627de813d3dc60cc637b26d23931e2e]


3. Check claim: `checkClaim method Response | bool: true`

