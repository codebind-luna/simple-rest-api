// user/model.go
package user

type User struct {
    Id   int64
    FullName string  `xorm:"varchar(25)"`
}

type Users []User