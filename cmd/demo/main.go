package main

import (
	"fmt"
	"os/exec"
)

func main() {
	users := getUsers()
	for i, u := range users {
		bet := fmt.Sprintf("%d", i+1)
		cmd := exec.Command("lotteryd", "tx", "lottery", "enter-lottery", bet, "token", "--from", u.address, "--fees", "5token", "-y")
		if _, err := cmd.Output(); err != nil {
			fmt.Println(err)
		}
	}
}

type user struct {
	address string
}

func getUsers() []user {
	users := make([]user, 20)

	for i := 0; i < 20; i++ {
		client := fmt.Sprintf("client%d", i+1)
		cmd := exec.Command("lotteryd", "keys", "show", client, "-a")
		out, _ := cmd.Output()
		users[i] = user{address: string(out[:len(out)-1])} // cut last byte because it is a new line "\n"
	}

	return users
}
