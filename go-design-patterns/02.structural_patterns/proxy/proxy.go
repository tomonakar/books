package proxy

import "fmt"

type UserFinder interface {
	FindUser(id int32) (User, error)
}

type User struct {
	ID int32
}

type UserList []User

type UserListProxy struct {
	MockedDatabase      *UserList
	StackCache          UserList
	StackSize           int
	LastSearchUsedCache bool
}

func (u *UserListProxy) FindUser(id int32) (User, error) {
	return User{}, fmt.Errorf("Not implemented")
}
