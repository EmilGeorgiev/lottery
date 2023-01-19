package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

type Balance struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}

type UserBalances struct {
	Balances []Balance `json:"balances"`
}

func main() {
	users := getUsers()

	var usersThatHasEnoughFunds []user
	for n := 0; n < 100; n++ {

		for i, u := range users {
			bet := fmt.Sprintf("%d", u.bet)
			cmd := exec.Command("lotteryd", "tx", "lottery", "enter-lottery", bet, "token", "--from", u.address, "--fees", "5token", "-y")

			if err := cmd.Run(); err != nil {
				fmt.Println(err)
			}
			fmt.Printf("Client%d, Bet: %d, for the %d times \n", u.bet, u.bet, n+1)

			// check whether the used has enough funds to enter the next lottery
			resp, err := http.Get(fmt.Sprintf("http://localhost:1317/cosmos/bank/v1beta1/balances/%s", u.address))
			if err != nil {
				fmt.Printf("Can't get balance of the client: %s, error: %s\n", u.address, err.Error())
				continue
			}

			var ub UserBalances
			if err = json.NewDecoder(resp.Body).Decode(&ub); err != nil {
				fmt.Println("can't decode response: ", err.Error())
			}

			if len(ub.Balances) == 0 {
				continue
			}

			balance, err := strconv.ParseInt(ub.Balances[0].Amount, 10, 64)
			if err != nil {
				fmt.Println("Can't parse the balance: ", err)
				continue
			}

			if balance < int64(5+i+1) {
				fmt.Printf("Client%d with address %s is out of funds.\n", u.bet, u.address)
				continue
			}
			usersThatHasEnoughFunds = append(usersThatHasEnoughFunds, u)
		}

		if len(usersThatHasEnoughFunds) < 10 {
			fmt.Println("There are less then 10 clients with enough tokens to enter the next lottery.")
			os.Exit(0)
			return
		}

		users = usersThatHasEnoughFunds
		usersThatHasEnoughFunds = []user{}
	}
}

type user struct {
	address string
	bet     int64
}

func getUsers() []user {
	users := make([]user, 20)

	for i := 0; i < 20; i++ {
		client := fmt.Sprintf("client%d", i+1)
		cmd := exec.Command("lotteryd", "keys", "show", client, "-a")
		out, _ := cmd.Output()
		users[i] = user{
			address: string(out[:len(out)-1]), // cut last byte because it is a new line "\n"
			bet:     int64(i + 1),
		}
	}

	return users
}
