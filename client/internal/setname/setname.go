package setname

import "fmt"

type client struct {
	name string
}

func SetName() string {
	var newName string
	fmt.Print("Please enter your nickname: ")
	fmt.Scanln(&newName)
	newClient := client{
		name: newName,
	}
	return newClient.name
}
