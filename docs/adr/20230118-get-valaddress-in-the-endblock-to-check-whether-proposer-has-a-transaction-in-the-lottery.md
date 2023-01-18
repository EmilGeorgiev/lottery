# Get ConsAddress of the proposer in the EndBlock and check whether he has a transaction in the lottery

- Status: [draft | proposed | rejected | accepted | deprecated | … | superseded by [xxx](yyyymmdd-xxx.md)] <!-- optional -->
- Deciders: [list everyone involved in the decision] <!-- optional -->
- Date: [YYYY-MM-DD when the decision was last updated] <!-- optional. To customize the ordering without relying on Git creation dates and filenames -->
- Tags: [space and/or comma separated list of tags] <!-- optional -->

Technical Story: [description | ticket/issue URL] <!-- optional -->

## Context and Problem Statement

The chosen block proposer can't have any lottery transactions with itself as a sender, if this is the case, then the lottery won’t fire this block, and continue on the next one

## Decision Drivers <!-- optional -->

- Avoid the proposer of the block manipulating the data in order to win the lottery


## Decision Outcome

An address is public information normally used to reference an account. Addresses are derived from public keys using [ADR-28](https://github.com/cosmos/cosmos-sdk/blob/main/docs/architecture/adr-028-public-key-addresses.md). Three types of addresses specify a context when an account is used: 
    - [AccAddress](https://github.com/cosmos/cosmos-sdk/blob/1dba6735739e9b4556267339f0b67eaec9c609ef/types/address.go#L129) - identifies users, which are the sender of a message
    - [ValAddress](https://github.com/cosmos/cosmos-sdk/blob/23e864bc987e61af84763d9a3e531707f9dfbc84/types/address.go#L298) - identifies validator operators.
    - [ConsAddress](https://github.com/cosmos/cosmos-sdk/blob/74d7a0dfcd9f47d8a507205f82c264a269ef0612/types/address.go#L466). - identifies validator nodes that are participating in consensus. Validator nodes are derived using the ed25519 curve.

In the EndBlock we can get the ConsAddress from the context by using [context.BlockHeader().ProposerAddress](https://github.com/tendermint/tendermint/blob/64747b2b184184ecba4f4bffc54ffbcb47cfbcb0/proto/tendermint/types/types.pb.go#L284) and convert the bytes to ConsAddress.

When we have ConsAddress of the proposer we should check whether some of the users in the lottery have the same ConsAddress. How this can be done? Addresses are used to reference accounts. We can generate all of the three types of addresses by using [PubKey](https://github.com/cosmos/cosmos-sdk/blob/9fd866e3820b3510010ae172b682d71594cd8c14/crypto/types/types.go#L9). For example:

```go
acc := accauntKeeper.GetAccount(ctx, accAddr)
consAddr := types.GetConsAddress(acc.GetPubKey())

```
    
The method GetAccount comes from the Auth module's keeper interface AccountKeeperI, so we need to use the bank module to get the account. Finally, we can compare both ConsAddress. If the addresses are equal the application should not fire the lottery in this block. 

### Positive Consequences <!-- optional -->

- avoid validators to participating in the lottery

### Negative Consequences <!-- optional -->

- The program should get all accounts of the users of the lottery in the EndBlock

