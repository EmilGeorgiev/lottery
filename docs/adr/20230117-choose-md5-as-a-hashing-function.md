# Use MD5 as a hashing function when calculating the winner

- Status: accepted
- Deciders Emil Georgiev (emogeorgioev88@gmail.com)
- Date: [YYYY-MM-DD when the decision was last updated] <!-- optional. To customize the ordering without relying on Git creation dates and filenames -->
- Tags: [space and/or comma separated list of tags] <!-- optional -->

Technical Story: [description | ticket/issue URL] <!-- optional -->

## Context and Problem Statement

At the end of the lottery, the application should append the data of the transactions
(retaining their order), then hash the data to get the result. Then, take the lowest 16 bits of the resulting hash and do a modulo on the number of
lottery transactions in the block to determine the winner! The application should use an efficient hash function for deciding the winner

## Decision Drivers <!-- optional -->

- The hash function should be fast when hashing the data because the lottery can contain many transactions.
- The lowest 16 bits are good to be different even if minimum changes are made to the data. Otherwise, the winner will be the same most of the time. 

## Considered Options

- MD4 
- MD5 
- SHA1 
- SHA224 
- SHA256 
- SHA384 
- SHA512 
- MD5SHA1 
- RIPEMD160 
- SHA3_224 
- SHA3_256 
- SHA3_384 
- SHA3_512 
- SHA512_224 
- SHA512_256 
- BLAKE2s_256 
- BLAKE2b_256 
- BLAKE2b_384 
- BLAKE2b_512 

## Decision Outcome

The MD5 and SHA1 fit our decision drivers. Both were designed to check the consistency of large data and they are very fast. The lottery can contain a large number of transactions and a slower hash function can rise a problem. Also, the lowest 16 bits are different even if we change one byte of the whole lottery, while at the others the lowest 16 bits frequently remain the same when the changes are made in the data. 
Between MD5 and SH1, I have chosen MD5 because is faster than SHA1. NOTE. These algorithms were not designed to protect passwords or other sensitive information. They are far too weak for that. But our case is different.

### Positive Consequences <!-- optional -->

- fast sped when hashing large lottery with many transactions
- different lowest 16 bits every time

## Links <!-- optional -->

- [MD5](https://en.wikipedia.org/wiki/MD5)
- [SHA-1](https://en.wikipedia.org/wiki/SHA-1)
