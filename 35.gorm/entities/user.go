package entities

type User struct {
	ID      int    `gorm:"primary_key"`
	Name    string `sql:"type:text;"`
	Address string `sql:"type:text;"`
}
