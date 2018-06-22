package db

import (
	"github.com/go-xorm/xorm"
    "github.com/lunaimaginea/rest_api_test/user"
)

func InitializeDb(engine *xorm.Engine) {
	engine.CreateTables(&user.User{})
}