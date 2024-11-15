package store

import "github.com/google/uuid"

var (
	users = []*User{
		{
			ID:   "1",
			Name: "Alice",
		},
		{
			ID:   "2",
			Name: "Bob",
		},
	}
)

type (
	User struct {
		ID   string
		Name string
	}
)

func Users() []*User {
	return users
}

func GetUser(id string) *User {
	for _, u := range users {
		if u.ID == id {
			return u
		}
	}
	return nil
}

func CreateUser(name string) *User {
	user := &User{
		ID:   uuid.New().String(),
		Name: name,
	}
	users = append(users, user)
	return user
}
