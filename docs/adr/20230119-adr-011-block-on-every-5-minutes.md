# ADR-011: Block on every 5 minutes

- Status: draft
- Deciders: Emil Georgiev (emogeorgiev88@gmail.com)
- Date: [YYYY-MM-DD when the decision was last updated] <!-- optional. To customize the ordering without relying on Git creation dates and filenames -->
- Tags: [space and/or comma separated list of tags] <!-- optional -->

Technical Story: [description | ticket/issue URL] <!-- optional -->

## Context and Problem Statement

We should configure the chain to produce a block on every 5 minutes.

## Decision Drivers <!-- optional -->

- this is requirement in the task.


## Decision Outcome

After a long research on how this can be done, I tried different approaches:

 1. Update configuration in $HOME/.lottery/config/config.toml in the section "Consensus Configuration Options". This section contains several configuration values:

```
#How long we wait for a proposal block before prevoting nil 
timeout_propose = "1s" - 
# How much timeout_propose increases with each round
timeout_propose_delta = "500ms"
# How long we wait after receiving +2/3 prevotes for “anything” (ie. not a single block or nil)
timeout_prevote = "1s"
# How much the timeout_prevote increases with each round
timeout_prevote_delta = "500ms"
# How long we wait after receiving +2/3 precommits for “anything” (ie. not a single block or nil)
timeout_precommit = "1s"
# How much the timeout_precommit increases with each round
timeout_precommit_delta = "500ms"
# How long we wait after committing a block, before starting on the new
# height (this gives us a chance to receive some more precommits, even
# though we already have +2/3).
timeout_commit = "1s"
```


I updated all of them, but the results ware not what I expected.

  2. Update configuration in $HOME/.lottery/config/genesis.toml. In this file I update the field "time_iota_ms", but again the result wasn't what I expected.

  3. Update configuration in $HOME/.tendermint/config/config.toml in the section "Consensus Configuration Options". This section contains several configuration values:    

```
# How long we wait for a proposal block before prevoting nil
timeout_propose = "300s"
# How much timeout_propose increases with each round
timeout_propose_delta = "300s"
# How long we wait after receiving +2/3 prevotes for “anything” (ie. not a single block or nil)
timeout_prevote = "1s"
# How much the timeout_prevote increases with each round
timeout_prevote_delta = "300s"
# How long we wait after receiving +2/3 precommits for “anything” (ie. not a single block or nil)
timeout_precommit = "300s"
# How much the timeout_precommit increases with each round
timeout_precommit_delta = "300s"
# How long we wait after committing a block, before starting on the new
# height (this gives us a chance to receive some more precommits, even
# though we already have +2/3).
timeout_commit = "300s"

# How many blocks to look back to check existence of the node's consensus votes before joining consensus
# When non-zero, the node will panic upon restart
# if the same consensus key was used to sign {double_sign_check_height} last blocks.
# So, validators should stop the state machine, wait for some blocks, and then restart the state machine to avoid panic.
double_sign_check_height = 0

# Make progress as soon as we have all the precommits (as if TimeoutCommit = 0)
skip_timeout_commit = false

# EmptyBlocks mode and possible interval between empty blocks
create_empty_blocks = true
create_empty_blocks_interval = "300s"
```

I updated all of them but the results weren't what I expected.

  4. In the tendermint project on GitHub I found this [issue](https://github.com/tendermint/tendermint/issues/5911). From the issue, I understand
that "It's unclear how-to change block times. Even if it was explained well it would be confusing because it is local to each node and not global.". 
So it seems that there is a way to change the time locally,

  5. Looking in the Cosmos Hub Forum and Discord channel, I found that the block time can be configurable from the section "Consensus Configuration Options" of configuration files, but I tried that already.     


My decision is to leave it for now and try again later if I have time.

### Positive Consequences <!-- optional -->

- [e.g., improvement of quality attribute satisfaction, follow-up decisions required, …]
- …

### Negative Consequences <!-- optional -->

- [e.g., compromising quality attribute, follow-up decisions required, …]
- …

## Pros and Cons of the Options <!-- optional -->

### [option 1]

[example | description | pointer to more information | …] <!-- optional -->

- Good, because [argument a]
- Good, because [argument b]
- Bad, because [argument c]
- … <!-- numbers of pros and cons can vary -->

### [option 2]

[example | description | pointer to more information | …] <!-- optional -->

- Good, because [argument a]
- Good, because [argument b]
- Bad, because [argument c]
- … <!-- numbers of pros and cons can vary -->

### [option 3]

[example | description | pointer to more information | …] <!-- optional -->

- Good, because [argument a]
- Good, because [argument b]
- Bad, because [argument c]
- … <!-- numbers of pros and cons can vary -->

## Links <!-- optional -->

- [Link type](link to adr) <!-- example: Refined by [xxx](yyyymmdd-xxx.md) -->
- … <!-- numbers of links can vary -->
