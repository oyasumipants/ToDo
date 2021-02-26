package models

import "time"

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	PassWord  string
	CreatesAr time.Time
}

func (u *User) CreateUser() (err error) {
	cmd := `insert into users (
		uuid,
		name,
		email,
		password,
		created_at) values (?, ?, ?, ?, ?)`

	// Dbは modelsパッケージ内(base.go)に存在するので，使用できる
	_, err := Db.Exec(cmd)
}
