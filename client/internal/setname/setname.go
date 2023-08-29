package setname

import "fmt"

// The type describes the client.
// This type have only one field is the name.
type client struct {
	name string
}

// This method assigns a nickname
// to the client.
func SetName() string {
	var newName string
	fmt.Print("Please enter your nickname: ")
	fmt.Scanln(&newName)
	newClient := client{
		name: newName,
	}
	return newClient.name
}
