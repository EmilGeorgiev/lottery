# Custom AnteHanlder to deduct fees from the sender account

- Status: accepted
- Deciders: [Emil Georgiev (emogeorgiev88@gmail.com)]
- Date: [YYYY-MM-DD when the decision was last updated] <!-- optional. To customize the ordering without relying on Git creation dates and filenames -->
- Tags: [space and/or comma separated list of tags] <!-- optional -->

Technical Story: [description | ticket/issue URL] <!-- optional -->

## Context and Problem Statement

When the user sends a valid transaction to enter the lottery, the application must charge it with the fees. The transaction must be rejected if the used 
doesn't provide the fees.

## Decision Drivers <!-- optional -->

- This is a requiretment in the task. This is the way the lottery works.
- Prevent spam and abuse from end-user.

## Considered Options

- Custom AnteHandler
- Charge the custumer directgly in the EnterLottery handler


## Decision Outcome

Chosen option: "Custom AnteHandler", because this is a good practice. Most applications implement fee mechanisms to prevent spam by using the AnteHandler

### Positive Consequences <!-- optional -->

- Separation of the concerns. The application will have a separate implementation for deducting the fees of the user account.
- The application will follow the best practices

### Negative Consequences <!-- optional -->

- It will take more time to be implemented


## Links <!-- optional -->

- [Fees in Cosmos SDK](https://docs.cosmos.network/main/basics/gas-fees)
- [Ante](https://github.com/cosmos/cosmos-sdk/tree/main/x/auth/ante)
