# Use the bank module to transfer coins

- Status: accepted
- Deciders: Emil Georgiev (emogeorgiev88@gmail.com)
- Date: [YYYY-MM-DD when the decision was last updated] <!-- optional. To customize the ordering without relying on Git creation dates and filenames -->
- Tags: [space and/or comma separated list of tags] <!-- optional -->

Technical Story: [description | ticket/issue URL] <!-- optional -->

## Context and Problem Statement

When the user sends a valid transaction and enters the lottery, the application needs to charge the user with fees + bet. How we will transfer the coins from the user's account to the module? 


## Decision Outcome

The bank module, from the Cosmos SDK will be used. It is responsible for handling multi-asset coin transfers between accounts. The module has interface with methods like SendCoinsFromModuleToAccount and SendCoinsFromAccountToModule that can be used.


## Links <!-- optional -->

- [Bank module](https://github.com/cosmos/cosmos-sdk/tree/main/x/bank)
