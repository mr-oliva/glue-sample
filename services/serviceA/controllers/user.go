package controllers

type User struct {
	repository Repository
}

type Repository interface {
	GetAllUsers() ([]string, error)
	GetUserNameById(id int) (string, error)
}

func NewUser(r Repository) *User {
	return &User{repository: r}
}

func (u *User) GetList() ([]string, error) {
	return u.repository.GetAllUsers()
}

func (u *User) GetNameById(id int) (string, error) {
	return u.repository.GetUserNameById(id)
}
