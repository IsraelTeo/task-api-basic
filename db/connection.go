package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DSN = "root:@tcp(localhost:3306)/tasksapi"
var GDB *gorm.DB

func DBConnection() error {
	var err error
	GDB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}
