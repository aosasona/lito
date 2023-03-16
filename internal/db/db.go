package db

import (
	"fmt"

	"github.com/aosasona/lito/internal/errtype"
	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
)

func Init() *xorm.Engine {
	conn, err := xorm.NewEngine("sqlite", "./data.db")
	if err != nil {
		panic(fmt.Sprintf("%s: %s", errtype.UNABLE_TO_CONNECT_TO_DB, err.Error()))
	}
	return conn
}
