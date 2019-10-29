package models

import (
	"github.com/go-xorm/xorm"
	DB "gpi/libriries/database"
)

type Common struct {
}

var dbConn *xorm.Engine

func init () {
	dbConn = DB.Engine
}