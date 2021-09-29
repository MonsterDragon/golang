package linkmysql

import (
	"fmt"
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

type Miners_1 struct {
	ID      uint   `gorm:"primary_key;AUTO_INCREMENT"`
	MinerId string `gorm:"column:miner_id;type:varchar(64);unique"`
}

type Lucky24h_1 struct {
	ID                 uint `gorm:"primary_key;AUTO_INCREMENT"`
	IsVerified         int
	CreateAt           time.Time
	SectorSizeStr      string `gorm:"varchar(127)"`
	MinedBlocks        int
	BlockRewardStr     string
	BlockRewardPercent string
	QualityPowerStr    string
	LuckyValue         float64 `gorm:"PTRCISION:3"`
	WinCount           int
	Tag                string

	Miners_1 Miners_1 `gorm:"ForeignKey:Refer;AssociationForeignKey:ID"`
	Refer    uint
}

type Lucky7d_1 struct {
	ID                 uint `gorm:"primary_key;AUTO_INCREMENT"`
	IsVerified         int
	CreateAt           time.Time
	SectorSizeStr      string `gorm:"varchar(127)"`
	MinedBlocks        int
	BlockRewardStr     string
	BlockRewardPercent string
	QualityPowerStr    string
	LuckyValue         float64 `gorm:"PTRCISION:3"`
	WinCount           int
	Tag                string

	Miners_1 Miners_1 `gorm:"ForeignKey:Refer;AssociationForeignKey:ID"`
	Refer    uint
}

type Lucky30d_1 struct {
	ID                 uint `gorm:"primary_key;AUTO_INCREMENT"`
	IsVerified         int
	CreateAt           time.Time
	SectorSizeStr      string `gorm:"varchar(127)"`
	MinedBlocks        int
	BlockRewardStr     string
	BlockRewardPercent string
	QualityPowerStr    string
	LuckyValue         float64 `gorm:"PTRCISION:3"`
	WinCount           int
	Tag                string

	Miners_1 Miners_1 `gorm:"ForeignKey:Refer;AssociationForeignKey:ID"`
	Refer    uint
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
	db.AutoMigrate(&Lucky24h_1{}, &Lucky7d_1{}, &Lucky30d_1{})
}

func InsertMiner(db *gorm.DB, minerid string) {
	miner := Miners_1{MinerId: minerid}
	fmt.Printf("%#v\n", miner)
	result_1 := db.Create(&miner)
	if result_1.Error != nil {
		panic(result_1.Error)
	}
}

func Insert24h(db *gorm.DB, id uint, ver int, size string, mb int, brs string, brp string, qps string, lv float64, wc int, tag string) {
	current := time.Now()
	h, _ := time.ParseDuration("8h")
	h1 := current.Add(h)
	lucky24h := Lucky24h_1{IsVerified: ver, CreateAt: h1, SectorSizeStr: size, MinedBlocks: mb, BlockRewardStr: brs, BlockRewardPercent: brp, QualityPowerStr: qps, LuckyValue: lv, WinCount: wc, Tag: tag, Refer: id}
	result_2 := db.Create(&lucky24h)
	if result_2.Error != nil {
		panic(result_2.Error)
	}
}

func Insert7d(db *gorm.DB, id uint, ver int, size string, mb int, brs string, brp string, qps string, lv float64, wc int, tag string) {
	current := time.Now()
	h, _ := time.ParseDuration("8h")
	h1 := current.Add(h)
	lucky7d := Lucky7d_1{IsVerified: ver, CreateAt: h1, SectorSizeStr: size, MinedBlocks: mb, BlockRewardStr: brs, BlockRewardPercent: brp, QualityPowerStr: qps, LuckyValue: lv, WinCount: wc, Tag: tag, Refer: id}
	result_3 := db.Create(&lucky7d)
	if result_3.Error != nil {
		panic(result_3.Error)
	}
}

func Insert30d(db *gorm.DB, id uint, ver int, size string, mb int, brs string, brp string, qps string, lv float64, wc int, tag string) {
	current := time.Now()
	h, _ := time.ParseDuration("8h")
	h1 := current.Add(h)
	lucky30d := Lucky30d_1{IsVerified: ver, CreateAt: h1, SectorSizeStr: size, MinedBlocks: mb, BlockRewardStr: brs, BlockRewardPercent: brp, QualityPowerStr: qps, LuckyValue: lv, WinCount: wc, Tag: tag, Refer: id}
	result_4 := db.Create(&lucky30d)
	if result_4.Error != nil {
		panic(result_4.Error)
	}
}

// func InsertData(table_name string, db *sql.DB, args ...interface{}) {

// 	column_name := "()"
// 	insert_sql := "insert into " + table_name +
// 	ret, err := db.Exec(insert_sql)
// }
