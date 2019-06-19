package twofer

import "fmt"

// ShareWith will take name arg, will retrun string
func ShareWith(name string) string {
	if name == "Bob" {
		s := "One for Bob, one for me."
		fmt.Println(s)
		return s
	} else if name == "Alice" {
		s := "One for Alice, one for me."
		fmt.Println(s)
		return s
	} else {
		s := "One for you, one for me."
		fmt.Println(s)
		return s
	}
}
