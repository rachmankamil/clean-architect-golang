package users

import (
	"ca-amartha/businesses/users"
	"time"
)

type Users struct {
	Id        int
	Name      string
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (rec *Users) toDomain() users.Domain {
	return users.Domain{
		Id:        rec.Id,
		Name:      rec.Name,
		Username:  rec.Username,
		Password:  rec.Password,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func fromDomain(userDomain users.Domain) Users {
	return Users{
		Id:        userDomain.Id,
		Name:      userDomain.Name,
		Username:  userDomain.Username,
		Password:  userDomain.Password,
		CreatedAt: userDomain.CreatedAt,
		UpdatedAt: userDomain.UpdatedAt,
	}
}
