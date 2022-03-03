package main

import (
	"time"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// type Miners struct {
// 	ID          uint         `gorm:"primary_key;AUTO_INCREMENT"`
// 	MinerId     string       `gorm:"column:miner_id;type:varchar(64);unique"`

// 	Lucky24h_1s []Lucky24h_1 `gorm:"foreignkey:Refer;association_foreignkey:MinerId"`
// }

// type Lucky24h_1 struct {
// 	ID      uint `gorm:"primary_key;AUTO_INCREMENT"`
// 	Refer   string
// 	AddTime time.Time
// }

// func (Lucky24h_1) TableName() string {
// 	return "hhhh"
// }

// func (Miners) TableName() string {
// 	return "mmmm"
// }

type Lucky24h struct {
	ID       int `gorm:"primary_key;AUTO_INCREMENT"`
	CreateAt time.Time
}

func ConnMysql(user string, pass string, dbname string) (db *gorm.DB) {

	param_string := user + ":" + pass + "@tcp(" + "10.1.2.65:3306" + ")/" + dbname + "?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open("mysql", param_string)
	if err != nil {
		panic(err)
	}
	return db
}

func Migrate(db *gorm.DB) {

	db.SingularTable(true)
	// db.AutoMigrate(&Miners_1{}, &Lucky24h_1{}, &Lucky7d_1{}, &Lucky30d_1{})
	// db.AutoMigrate(&Miners_1{})
	db.AutoMigrate(&Lucky24h{})
}

func Insert24h(db *gorm.DB) {
	current := time.Now()
	h, _ := time.ParseDuration("8h")
	h1 := current.Add(h)
	lucky24h := Lucky24h{CreateAt: h1}
	result := db.Create(&lucky24h)
	if result.Error != nil {
		panic(result.Error)
	}
}

func main() {
	db := ConnMysql("root", "123", "spider")
	// Migrate(db)
	Insert24h(db)
}
