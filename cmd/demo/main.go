package main

import (
	"fmt"
	"os/exec"
)

func main() {
	users := getUsers()
	for i, u := range users {
		fmt.Println(u.address)
		bet := fmt.Sprintf("%d", i+1)
		cmd := exec.Command("lotteryd", "tx", "lottery", "enter-lottery", bet, "token", "--from", u.address, "-y")
		b, err := cmd.Output()
		fmt.Println(err)
		fmt.Println(string(b))
	}

	cmd := exec.Command("lotteryd", "query", "lottery", "show-lottery")
	out, err := cmd.Output()
	if err != nil {
		// if there was any error, print it here
		fmt.Println("could not run command: ", err)
	}
	fmt.Println("Oytput: ", string(out))

	cmd = exec.Command("lotteryd", "query", "lottery", "list-finished-lottery")
	out, err = cmd.Output()
	if err != nil {
		// if there was any error, print it here
		fmt.Println("could not run command: ", err)
	}
	fmt.Println("Oytput: ", string(out))
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
