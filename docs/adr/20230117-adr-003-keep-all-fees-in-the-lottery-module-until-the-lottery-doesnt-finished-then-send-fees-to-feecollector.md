# ADR-003: Keep all fees in the lottery module at least the lottery doesn't finished. Then send fees to FeeCollector

- Status: accepted
- Deciders: Emil Georgiev (emogeorgiev88@gmail.com)
- Date: [YYYY-MM-DD when the decision was last updated] <!-- optional. To customize the ordering without relying on Git creation dates and filenames -->
- Tags: [space and/or comma separated list of tags] <!-- optional -->

Technical Story: [description | ticket/issue URL] <!-- optional -->

## Context and Problem Statement

If the winner of the lottery placed the highest bet, the application must pay the whole lottery pool (not only from the current lottery) + fee from the current lottery. By defaut all fees are transfered to the FeeCollector from the AnteHandler, but our custum ante handler prevent that and send the fees to the lottery module.

## Decision Outcome

The lottery module will contains the all the fees until lottery doesn't finished, because in this way the application can easy pay reward to the winner if he placed the highest bet. The application will send all bets + all fees from the lottery module to the winner account.
If the winner doesn't placed the highest bet, the application can sends the fees to the FeeCollector.  

### Positive Consequences <!-- optional -->

- the application will handle easy the case in which the winner placed the highest bet. 

### Negative Consequences <!-- optional -->

- The lottery module should keep fees and care about them at least until the lottery doesn't finish.

## Links
[FeeCollector](https://github.com/allinbits/cosmos-sdk/blob/69ab58ed2e7f50f4a5d2726454a1c467f17f4e09/x/auth/types/keys.go#L15)
