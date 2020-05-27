package dataaccess

import (
	"log"

	"Snai.Gin/35.gorm/entities"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open("mysql", "root:810618@/snai?charset=utf8&parseTime=True")

	if err == nil {
		hasTable := Db.HasTable(entities.User{})
		if !hasTable {
			Db.CreateTable(entities.User{})
		}
	} else {
		log.Panicln("err:", err.Error())
	}
}

func AddUser(user *entities.User) error {
	return Db.Create(&user).Error
}

func GetUser(id int) (*entities.User, error) {
	var user entities.User
	err := Db.First(&user, id).Error
	return &user, err
}

func GetUsers() (*[]entities.User, error) {
	var users []entities.User
	err := Db.Find(&users).Error
	return &users, err
}

func ModifyUser(user *entities.User) error {
	return Db.Save(user).Error
}

func DeleteUser(user *entities.User) error {
	return Db.Delete(user).Error
}
