// user/create.go
package user

import (
  "github.com/go-xorm/xorm"
)

func Create(engine *xorm.Engine, users *Users) {
  engine.Insert(users)
}