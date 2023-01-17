# lottery
**lottery** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

### Rules of the lottery

Anyone can enter the lottery as long as they have enough funds.
A winner is chosen at the end of a block if the lottery has 10 or more valid lottery
transactions. If there weren’t enough transactions, the lottery continues. Once a winner is chosen, a payout is sent and the next lottery cycle begins

## Get started

First ypu should have Ignite CLI installed on you computer. This application is build with version v0.25.2. To install it at the command line

```
curl https://get.ignite.com/cli@v0.25.2! | bash
```
You can verify the version of Ignite CLI you have once it is installed:

```
ignite version
```

This prints its version:

```
Ignite CLI version:	v0.25.2
...
```

When you are at lottery folder run this command to start the chain 

```
ignite chain serve
```

The command compiles the source code into a binary called **lotteryd**. It installs dependencies, builds, initializes the node with a single validator, add accounts, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

## DEMO

The application is already started so we can interact with it through sending queries and transactions.
The project contains a simple client application that sends 20 entered lottery transaction from 20 different clients, with 20 different bets. For example *client1: 1token*, *client2: 2token*, *client3: 3token*, ... and so on. Run the command below to execute the client

```
go run ./cmd/demo/main.go
```

Wait for the program to finish and then see the results: 

```

```

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.ignite.com/EmilGeorgiev/lottery@latest! | sudo bash
```
`EmilGeorgiev/lottery` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)
