package request

import "ca-amartha/businesses/users"

type Users struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (req *Users) ToDomain() *users.Domain {
	return &users.Domain{
		Name:     req.Name,
		Username: req.Username,
		Password: req.Password,
	}
}
