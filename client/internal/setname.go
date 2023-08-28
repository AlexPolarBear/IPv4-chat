package internal

import "fmt"

type client struct {
	name string
}

func setName() string {
	var newName string
	fmt.Print("Please enter your nickname:")
	fmt.Scan(&newName)
	newClient := client{
		name: newName,
	}
	return newClient.name
}
