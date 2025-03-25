package types

import "fmt"

type Permission uint8

func (p Permission) String() string {
	return fmt.Sprintf("%o", uint8(p))
}

func (p Permission) HasPermission(perm Permission) bool {
	return p&perm == perm
}

func (p Permission) Add(perms ...Permission) Permission {
	result := p
	for _, perm := range perms {
		result |= perm
	}
	return result
}
