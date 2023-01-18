# New field DateTime to the enter lottery tx

- Status: accepted
- Deciders: Emil Georgiev
- Date: [YYYY-MM-DD when the decision was last updated] <!-- optional. To customize the ordering without relying on Git creation dates and filenames -->
- Tags: [space and/or comma separated list of tags] <!-- optional -->

Technical Story: [description | ticket/issue URL] <!-- optional -->

## Context and Problem Statement

The winner is choosen by hashing the transactions data in the lottery. If the same users send transactions in the same order with the same bet, the winner will be the same. For example if we have the following bets: 
    - client1: 1token,
    - client2: 2token,
    - client3: 3token,
    ....
    - client20: 20token 

Every lottery that have exactly the same order of transactions will produce the same winner.

## Decision Drivers <!-- optional -->

- The lottery should avoid choosing the same winner multiple times. 

## Decision Outcome

To produce different output, the data should be unique before hashing (for more information about the hashing see [ADR-Hashing](https://github.com/EmilGeorgiev/lottery/blob/master/docs/adr/20230117-choose-md5-as-a-hashing-function.md)). The datetime + account address + bet guarantee that every transaction will be unique and the hash result will be different even if the same users sent the same bet in the same order.

### Positive Consequences <!-- optional -->

- different winner every time

### Negative Consequences <!-- optional -->

- additional data to keep in the storage

