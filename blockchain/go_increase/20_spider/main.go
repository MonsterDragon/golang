package main

import (
	"20_spider/linkmysql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"sort"
	"strings"
	"unicode/utf8"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Response struct {
	Code      int    `json:"code"`
	Total     int    `json:"total"`
	PageIndex int    `json:"PageIndex"`
	PageSize  int    `json:"pageSize"`
	Message   string `json:"message"`
	Data      Data   `json:"data"`
}

type Data struct {
	MaxBlocks           int      `json:"maxBlocks"`
	MaxBlockReward      float64  `json:"maxBlockReward"`
	TotalBlockReward    float64  `json:"totalBlockReward"`
	TotalBlockRewardStr string   `json:"totalBlockRewardStr"`
	Miners              []Miners `json:"miners"`
}

type Miners struct {
	Miner              string  `json:"miner"`
	IsVerified         int     `json:"isVerified"`
	SectorSize         int     `json:"sectorSize"`
	SectorSizeStr      string  `json:"sectorSizeStr"`
	MinedBlocks        int     `json:"minedBlocks"`
	BlockReward        float64 `json:"blockReward"`
	BlockRewardStr     string  `json:"blockRewardStr"`
	BlockRewardPercent string  `json:"blockRewardPercent"`
	QualityPower       int     `json:"qualityPower"`
	QualityPowerStr    string  `json:"qualityPowerStr"`
	LuckyValue         float64 `json:"luckyValue"`
	LuckyValue2        float64 `json:"luckyValue2"`
	WinCount           int     `json:"winCount"`
	Region             string  `json:"region"`
	FavorTag           string  `json:"favorTag"`
	Tag                string  `json:"tag"`
}

var qyys_miner []int = []int{3831, 3989, 4131, 4228, 4267, 4362, 4411, 4550, 4555, 4563, 4569, 4570, 4572, 4613, 4676, 4689, 4695, 4719, 4733, 4749, 4775, 4802, 4847, 4928, 4974, 4994, 5041,
	5199, 5334, 5418, 5452, 5643, 5660, 5771, 5889, 5935, 6009, 6098, 13558, 13726, 15117}

func PosterData(pageIndex int, statsType string) (body []byte, err error) {

	url := "https://api.filscout.com/api/v1/minerrank/blocks"
	contentType := "application/json"
	data := fmt.Sprintf("{\"continent\":0,\"pageIndex\":%d,\"pageSize\":50,\"sectorSize\":\"\",\"statsType\":\"%s\"}", pageIndex, statsType)
	resp, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read from resp.Body failed, err:%v\n", err)
		return
	}
	return body, err
}

func PrintLines(filePath string, values interface{}) error {

	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	rv := reflect.ValueOf(values)
	if rv.Kind() != reflect.Slice {
		return errors.New("Not a slice")
	}
	for i := 0; i < rv.Len(); i++ {
		fmt.Fprintln(f, rv.Index(i).Interface())
	}

	return nil
}

func Distinct(db *gorm.DB, minerid string) (id uint) {

	miner := new(linkmysql.Miners_1)
	db.Where("miner_id = ?", minerid).First(miner)
	if (*miner).ID != 0 {
		return (*miner).ID
	} else {
		return 0
	}
}

func GetMinerData(page_num int, param string, db *gorm.DB) {

	var r1 Response
	body, _ := PosterData(page_num, param)
	err := json.Unmarshal(body, &r1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("page_num: %#v\n", page_num)
	for _, i := range r1.Data.Miners {

		id := Distinct(db, i.Miner)
		if id == 0 {
			linkmysql.InsertMiner(db, i.Miner)
		}
	}
}

// wg sync.WaitGroup

func ExecMiner(db *gorm.DB) {

	var param_list []string = []string{"24h"}

	for _, i := range param_list {
		for j := 1; j < 200; j++ {
			GetMinerData(j, i, db)
		}
	}
}

func FilterEmoji(content string) string {
	new_content := ""
	for _, value := range content {
		_, size := utf8.DecodeRuneInString(string(value))
		if size <= 3 {
			new_content += string(value)
		}
	}
	return new_content
}

func Get24hData(page_num int, param string, db *gorm.DB) {

	var r1 Response
	body, _ := PosterData(page_num, param)
	err := json.Unmarshal(body, &r1)
	if err != nil {
		panic(err)
	}
	for _, i := range r1.Data.Miners {
		id := Distinct(db, i.Miner) // 判断当前查询的miner_id是否在miners_1的表中
		if id == 0 {
			linkmysql.InsertMiner(db, i.Miner)
			id = Distinct(db, i.Miner)
		}
		tag := FilterEmoji(i.Tag)
		index := sort.SearchInts(qyys_miner, int(id))
		if index < len(qyys_miner) && qyys_miner[index] == int(id) {
			linkmysql.Insert24h(db, id, i.IsVerified, i.SectorSizeStr, i.MinedBlocks, i.BlockRewardStr, i.BlockRewardPercent, i.QualityPowerStr, i.LuckyValue, i.WinCount, "qyys")
		} else {
			linkmysql.Insert24h(db, id, i.IsVerified, i.SectorSizeStr, i.MinedBlocks, i.BlockRewardStr, i.BlockRewardPercent, i.QualityPowerStr, i.LuckyValue, i.WinCount, tag)
		}
	}
}

func Exec24h(db *gorm.DB) {
	var param_list []string = []string{"24h"}

	for _, i := range param_list {
		for j := 1; j < 200; j++ {
			Get24hData(j, i, db)
		}
	}
}

func Get7dData(page_num int, param string, db *gorm.DB) {

	var r1 Response
	body, _ := PosterData(page_num, param)
	err := json.Unmarshal(body, &r1)
	if err != nil {
		panic(err)
	}
	for _, i := range r1.Data.Miners {
		id := Distinct(db, i.Miner)
		if id == 0 {
			linkmysql.InsertMiner(db, i.Miner)
			id = Distinct(db, i.Miner)
		}
		tag := FilterEmoji(i.Tag)
		index := sort.SearchInts(qyys_miner, int(id))
		if index < len(qyys_miner) && qyys_miner[index] == int(id) {
			linkmysql.Insert7d(db, id, i.IsVerified, i.SectorSizeStr, i.MinedBlocks, i.BlockRewardStr, i.BlockRewardPercent, i.QualityPowerStr, i.LuckyValue, i.WinCount, "qyys")
		} else {
			linkmysql.Insert7d(db, id, i.IsVerified, i.SectorSizeStr, i.MinedBlocks, i.BlockRewardStr, i.BlockRewardPercent, i.QualityPowerStr, i.LuckyValue, i.WinCount, tag)
		}
	}
}

func Exec7d(db *gorm.DB) {
	var param_list []string = []string{"7d"}

	for _, i := range param_list {
		for j := 1; j < 200; j++ {
			Get7dData(j, i, db)
		}
	}
}

func Get30dData(page_num int, param string, db *gorm.DB) {

	var r1 Response
	body, _ := PosterData(page_num, param)
	err := json.Unmarshal(body, &r1)
	if err != nil {
		panic(err)
	}
	for _, i := range r1.Data.Miners {
		id := Distinct(db, i.Miner)
		if id == 0 {
			linkmysql.InsertMiner(db, i.Miner)
			id = Distinct(db, i.Miner)
		}
		tag := FilterEmoji(i.Tag)
		index := sort.SearchInts(qyys_miner, int(id))
		if index < len(qyys_miner) && qyys_miner[index] == int(id) {
			linkmysql.Insert30d(db, id, i.IsVerified, i.SectorSizeStr, i.MinedBlocks, i.BlockRewardStr, i.BlockRewardPercent, i.QualityPowerStr, i.LuckyValue, i.WinCount, "qyys")
		} else {
			linkmysql.Insert30d(db, id, i.IsVerified, i.SectorSizeStr, i.MinedBlocks, i.BlockRewardStr, i.BlockRewardPercent, i.QualityPowerStr, i.LuckyValue, i.WinCount, tag)
		}
	}
}

func Exec30d(db *gorm.DB) {
	var param_list []string = []string{"30d"}

	for _, i := range param_list {
		for j := 1; j < 200; j++ {
			Get30dData(j, i, db)
		}
	}
}

func main() {

	// // var param_list []string = []string{"24h", "7d", "30d"}

	db := linkmysql.ConnMysql("root", "123", "miner_info")
	// ExecMiner(db)
	Exec24h(db)
	Exec7d(db)
	Exec30d(db)

	// linkmysql.Migrate(db)
	// linkmysql.InsertData(db)
	defer db.Close()
	// test_1()
}
