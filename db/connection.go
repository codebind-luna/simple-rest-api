package db

import (
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

func Connection() ( *xorm.Engine, error) {
	var engine *xorm.Engine


    var err error
    engine, err = xorm.NewEngine("sqlite3", "./test.db")
    engine.ShowSQL(true)

	
	return engine, err
}