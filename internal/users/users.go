package users

import (
	"fmt"
	"sync"
)

type UserManager interface {
	GetUserDetails(Userid int) *User
	GetAllUsers() []*User
	GetByUserName(username string) []*User
	CreateUser(name, description string, id int)
}

type UserRepo struct {
	Users []*User
}

var userrepo *UserRepo
var once sync.Once

type IUser interface {
	GetUserName() string
	GetDescription() string
}
type User struct {
	userid      int
	name        string
	description string
}

func (u *User) GetUserName() string {
	return u.name
}

func (u *User) GetDescription() string {
	return u.description
}

func (u *User) Update(message string) {
	fmt.Println("User:- "+u.name+" User id:-", u.userid, " notified : ", message)
}

// Singelton to create a User repo
func CreateUserRepo() *UserRepo {
	once.Do(func() {
		userrepo = &UserRepo{}
	})
	return userrepo
}

func (u *UserRepo) CreateUser(name, description string, id int) {

	t := &User{
		userid:      id,
		name:        name,
		description: description,
	}
	u.Users = append(u.Users, t)
	fmt.Println("User ", name, " added")
}

func (u *UserRepo) GetAllUsers() []*User {
	return u.Users
}

func (u *UserRepo) GetUserDetails(id int) *User {
	for _, us := range u.Users {
		if us.userid == id {
			return us
		}
	}
	return nil
}

func (u *UserRepo) GetByUserName(username string) []*User {
	var res []*User
	for _, us := range u.Users {
		if us.name == username {
			res = append(res, us)
		}
	}
	return res
}
