# lottery
**lottery** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

In the [.docs/adr](https://github.com/EmilGeorgiev/lottery/tree/master/docs/adr) you can find all descriptive information 
for architecture decisions.

This readme file contains two DEMO examples one for [happy path](https://github.com/EmilGeorgiev/lottery#demo) and one 
for [edge cases](https://github.com/EmilGeorgiev/lottery#demo-edge-cases) that the lottery
handle. For the DEMO is used a small cmd program that is placed in [.cmd/demo/main.go](https://github.com/EmilGeorgiev/lottery/blob/master/cmd/demo/main.go). 
Also, the folder [.cmd/demo/](https://github.com/EmilGeorgiev/lottery/tree/master/cmd/demo) contains files with the 
results of the local execution of the DEMO.

At the end of this file, you can see the answers of the [Bonus Strategies](https://github.com/EmilGeorgiev/lottery#bonus-strategy) questions.

What was not implemented? Block on every 5 minutes is not implemented. You can find why in [ADR-11](https://github.com/EmilGeorgiev/lottery/blob/master/docs/adr/20230119-adr-011-block-on-every-5-minutes.md).

### Rules of the lottery

Anyone can enter the lottery as long as they have enough funds.
A winner is chosen at the end of a block if the lottery has 10 or more valid lottery
transactions. If there weren’t enough transactions, the lottery continues. Once a winner is chosen, a payout is sent and the next lottery cycle begins

## Get started

First, you should have Ignite CLI installed on your computer. This application is built with version v0.25.2. 
To install it at the command line type this command

```
curl https://get.ignite.com/cli@v0.25.2! | bash
```

You can verify the version of Ignite CLI you have once it is installed:

```
ignite version
```

This prints:

```
Ignite CLI version:	v0.25.2
...
```

When you are in the lottery folder run this command to start the chain 

```
ignite chain serve
```

The command compiles the source code into a binary called **lotteryd**. It installs dependencies, builds, initializes 
the node with a single validator, adds accounts, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

## DEMO

In the previous steps, we started the application now we can interact with it by sending queries and transactions.
The simple client application will send 20 entered lottery transactions from 20 different clients, 
with 20 different bets. For example *client1: 1token*, *client2: 2token*, *client3: 3token*, ... and so on. These 20 
transactions are repeated 100 times or until the clients run out of funds.  Run the command below to execute the client:

```
go run ./cmd/demo/main.go
```

This program can take several minutes. After it finished we can see the results. 
(**N0TE we will execute a large number with commands, your result from the commands will be different from the main.**). 

#### List all finished lotteries

We can list all finished lotteries by command:

```
lotteryd query lottery list-finished-lottery --limit 200
```

The result is:

![.cmd/demo/images/img_5.png](.cmd/demo/images/img_5.png)

The output is too big, you can an example of the result in the file [cmd/demo/list-finished-lotteries](https://github.com/EmilGeorgiev/lottery/blob/master/cmd/demo/list-finished-lotteries).
**NOTE** the file contains just an example. The data in the file is not from your execution of the program.

#### Get the system info

We can get the current system info, by  using the command:

```
lotteryd query lottery show-system-info
```

The result is:

![.cmd/demo/images/img_6.png](.cmd/demo/images/img_6.png)

As we can see in the lottery pool we have **240token** that will be paid as a reward to the next winner who placed the heights bet.
The system-info show what will be the **nextID** of the next finished lottery.

#### Get all balances of the clients

First, we should export the addresses of the clients in variables:

```
export client1=$(lotteryd keys show client1 -a) 
export client2=$(lotteryd keys show client2 -a)
export client3=$(lotteryd keys show client3 -a)
export client4=$(lotteryd keys show client4 -a)
export client5=$(lotteryd keys show client5 -a)
export client6=$(lotteryd keys show client6 -a)
export client7=$(lotteryd keys show client7 -a)
export client8=$(lotteryd keys show client8 -a)
export client9=$(lotteryd keys show client9 -a)
export client10=$(lotteryd keys show client10 -a)
export client11=$(lotteryd keys show client11 -a)
export client12=$(lotteryd keys show client12 -a)
export client13=$(lotteryd keys show client13 -a)
export client14=$(lotteryd keys show client14 -a)
export client15=$(lotteryd keys show client15 -a)
export client16=$(lotteryd keys show client16 -a)
export client17=$(lotteryd keys show client17 -a)
export client18=$(lotteryd keys show client18 -a)
export client19=$(lotteryd keys show client19 -a)
export client20=$(lotteryd keys show client20 -a)
```

Then we can use these exported variables to get the balances:

```
lotteryd query bank balances $client1
lotteryd query bank balances $client2
lotteryd query bank balances $client3
lotteryd query bank balances $client4
lotteryd query bank balances $client5
lotteryd query bank balances $client6
lotteryd query bank balances $client7
lotteryd query bank balances $client8
lotteryd query bank balances $client9
lotteryd query bank balances $client10
lotteryd query bank balances $client11
lotteryd query bank balances $client12
lotteryd query bank balances $client13
lotteryd query bank balances $client14
lotteryd query bank balances $client15
lotteryd query bank balances $client16
lotteryd query bank balances $client17
lotteryd query bank balances $client18
lotteryd query bank balances $client19
lotteryd query bank balances $client20
```

An example of the result for the client2 is:

![.cmd/demo/images/img_7.png](.cmd/demo/images/img_7.png)

In the file [./cmd/demo/balances](https://github.com/EmilGeorgiev/lottery/blob/master/cmd/demo/balances) you can see all balances. 
**NOTE** the file contains an example. This is not your data:

![.cmd/demo/images/img.png](.cmd/demo/images/img.png)


#### Get the current lottery:

We can see the current lottery:

```
lotteryd query lottery show-lottery
```

The result is:

![.cmd/demo/images/img_8.png](.cmd/demo/images/img_8.png)

As we can see the current lottery has zero transactions.


#### How can we be sure that the lottery works correctly?

Let's check whether the lottery working correctly. For example, we will use the client2 and will follow all his bets 
and rewards, we will calculate his balance and we will see whether our result matches the actual balance.  

 - First, we will get his address

```
lotteryd keys show client2 -a
```

The result is:

```
cosmos1kngwxau7zp6aydwqj9n2rwxes47xxrkcfxgs2w
```

 - calculate all the payments that **client2** made. We know that the client2 always places the bet with **2token**, 
also we know that he must pay a fee of **5token** per transaction. If we know the number of all lotteries in which the 
client2 placed a valid bet, we can calculate the total payments. Let's get that number by using this command:

```
lotteryd query lottery list-finished-lottery --limit 200 | grep -c "user_address: cosmos1kngwxau7zp6aydwqj9n2rwxes47xxrkcfxgs2w"
``` 

The result is:

![.cmd/demo/images/img_9.png](.cmd/demo/images/img_9.png)

The user placed  **81** valid bets. This means that he paid 5token fee + **2token**, total a **7token** per bet. The total tokens 
that he paid are **81x7 = 567**

 - check how many times the client2 won the lottery. We can see the result by using the command:

```
lotteryd query lottery list-finished-lottery --limit 200 | grep -B 2 -A 1 "winner: cosmos1kngwxau7zp6aydwqj9n2rwxes47xxrkcfxgs2w"
```

The result is:

![.cmd/demo/images/img_10.png](.cmd/demo/images/img_10.png)

client2 won 9 times the lottery. Also, from the picture, we can see how many rewards he got for every lottery. 
(100+55+55+55+55+55+115+75+133= 698). If we take out the payments from the rewards 698-567 = 131 is the profit. 
If we add the profit to the initial balance (500token) we will get that
the balance of the client2 should be **631token** we can confirm that:

 - Confirm that the balance of the client is **631token**.  

```
lotteryd query bank balances cosmos1kngwxau7zp6aydwqj9n2rwxes47xxrkcfxgs2w
```

The result is:

![.cmd/demo/images/img_12.png](.cmd/demo/images/img_12.png)

As we can see the balance is 631token, as we expected. 

## DEMO edge cases

In this demo, we will test some edge cases

### Users can't place a bet without exactly 5 fees.

 - Users can't place a bet without any fees:

```
lotteryd tx lottery enter-lottery 1 token --from cosmos1kngwxau7zp6aydwqj9n2rwxes47xxrkcfxgs2w -y
```

The result is:

```
raw_log: 'Tx must contains exactly 5 fee: insufficient fee'
```

 - The user can't place a bet with a lower fee than 5

```
lotteryd tx lottery enter-lottery 1 token --from cosmos1kngwxau7zp6aydwqj9n2rwxes47xxrkcfxgs2w --fees 4token -y
```

The result is:

```
raw_log: 'Tx must contains exactly 5 fee: insufficient fee'
``` 

 - Users can't place a bet with a fee higher than 5

```
lotteryd tx lottery enter-lottery 1 token --from cosmos1kngwxau7zp6aydwqj9n2rwxes47xxrkcfxgs2w --fees 6token -y
```

The result is:

```
raw_log: 'Tx must contains exactly 5 fee: insufficient fee'
```

### The bet of the user should be between 1-100

 - Place bet higher then 100

```
lotteryd tx lottery enter-lottery 101 token --from cosmos1wkx4pmsmy0hkc9xurhz4mxfwulwfh7rkl2ces9 --fees 5token -y
```

Then we can see that the current lottery is empty. The bet is not placed:

![.cmd/demo/images/img_13.png](.cmd/demo/images/img_13.png)


### The lottery won't fire if there are fewer than 10 users

 - place 9  bets from 9 different clients. **NOTE** you should have the addresses of the clients exported in variables. 
We did that already

```
lotteryd tx lottery enter-lottery 1 token --from client1 --fees 5token -y
lotteryd tx lottery enter-lottery 2 token --from client2 --fees 5token -y
lotteryd tx lottery enter-lottery 3 token --from client3 --fees 5token -y
lotteryd tx lottery enter-lottery 4 token --from client4 --fees 5token -y
lotteryd tx lottery enter-lottery 5 token --from client5 --fees 5token -y
lotteryd tx lottery enter-lottery 6 token --from client6 --fees 5token -y
lotteryd tx lottery enter-lottery 7 token --from client7 --fees 5token -y
lotteryd tx lottery enter-lottery 8 token --from client8 --fees 5token -y
lotteryd tx lottery enter-lottery 9 token --from client9 --fees 5token -y
```

Then we can see the lottery:

```
lotteryd query lottery show-lottery
```

The result is

![.cmd/demo/images/img_14.png](.cmd/demo/images/img_14.png)


There are 9 transactions in the lottery and the lottery will not fire. We can list the finished lotteries:

```
lotteryd query lottery list-finished-lottery
```

the result is

![.cmd/demo/images/img_15.png](.cmd/demo/images/img_15.png)    

as we can see the list is empty. But if we place one more bet:

```
lotteryd tx lottery enter-lottery 10 token --from client10 --fees 5token -y
```

Then the current lottery is:

![.cmd/demo/images/img_17.png](.cmd/demo/images/img_17.png)


And the finished lotteries are

![.cmd/demo/images/img_16.png](.cmd/demo/images/img_16.png)

The system-info is:

![.cmd/demo/images/img_18.png](.cmd/demo/images/img_18.png)


## Bonus strategy

### 1. Assuming uniform random bet from all other clients, what is the best strategy for client1?
    
The best strategy for client1 is to place a bigger bet than all other clients. If he doesn't know what is the 
uniform bet from all other clients, he should place the maximum allowed bet (100). By doing that he will be the only 
one that can win money from the lottery because all other clients placed the lowest bet (except if all other clients 
also placed a bet of 100) and the lottery will never pay a reward to them. When client1 wins the lottery he will get 
the whole lottery pool from the current and previous lotteries. Also, it is important the client has a big balance to 
avoid cases in which he is out of money. 

Let's test the strategy. Run the application with 10 clients with a balance of 50000 tokens. Everyone participates 
in the 100 lotteries. The result is: client1 has **90109token** in the balance.  

![.cmd/demo/images/img_1.png](.cmd/demo/images/img_1.png)

The all other clients have **44701token**.

![.cmd/demo/images/img_2.png](.cmd/demo/images/img_2.png)

We confirm that this is the best strategy for client1 


### 2. Assuming uniform random bet from all other clients, and client1 behaves in the strategy mentioned (1.), what is the best strategy for client2?

Let's test the strategy and see the result. Let's say that we have 10 clients with 50000token balance, client1
and client2 place the highest bet (100), and all other clients place a random bet. After the test finished we see that
the first two players have bigger balances:
 
![.cmd/demo/images/img_3.png](.cmd/demo/images/img_3.png)

All other players suffer a loss. They have **44701token** in their balances: 

![.cmd/demo/images/img_4.png](.cmd/demo/images/img_4.png)

### 3. What is the Nash equilibrium?

We have Nash equilibrium when all players placed the highest bet (100). This is also **strict** Nash equilibrium because 
every player will suffer a loss by changing his strategy.  


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
