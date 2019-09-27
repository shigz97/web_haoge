package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Eloquent *gorm.DB

func init() {
	var err error
	Eloquent, err = gorm.Open("mysql", "root:******@/shigz?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("failed to connnect mysql,err:", err)
	}
	if Eloquent.Error != nil {
		fmt.Println("database error,err:", Eloquent.Error)
	}
}
