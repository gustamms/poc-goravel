package models

import (
	"github.com/goravel/framework/database/orm"
)

type User struct {
	orm.Model
	Name     string
	Username string
	Password string
}

func (r *User) TableName() string {
	return "users"
}
