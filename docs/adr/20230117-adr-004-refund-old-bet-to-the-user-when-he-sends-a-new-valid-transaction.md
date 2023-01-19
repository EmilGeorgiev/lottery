# ADR-004: Refund old bet to the user when he sends a new valid transaction

- Status: accepted
- Deciders: Emil Georgiev (emogeorgiev88@gmail.com)
- Date: [YYYY-MM-DD when the decision was last updated] <!-- optional. To customize the ordering without relying on Git creation dates and filenames -->
- Tags: [space and/or comma separated list of tags] <!-- optional -->

Technical Story: [description | ticket/issue URL] <!-- optional -->

## Context and Problem Statement

The user can send more than one transaction per lottery and the application must keep only the last bet. The problem is that the user is charged with the value of the bet on every transaction, For example, if a user sends two transactions the first one with bet 5token and the second one with bet 1token, the application will set the bet 1, but will charge the user with total 6token.

## Decision Outcome

The application will refund the old bet to the user and keep only the bet from the last valid transaction. Note that the fees will not be refunded because computing power is already used to handle the previous transaction. Also, if the fees are refunded the users can send spam and abuse transactions because they will pay only once.

