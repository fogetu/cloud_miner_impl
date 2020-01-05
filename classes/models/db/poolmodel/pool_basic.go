package poolmodel

import (
	"fmt"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/fogetu/miner_service_intf/mine_intf"
)

// City 城市
type City struct {
	ID      uint   // 列名为 `id`
	Country string // 列名为 `country`
	Map     string // 列名为 `map`
	Name    string // 列名为 `name`
	State   string // 列名为 `state`
}

type Pool struct {
	AutoID         uint32    `orm:"column:auto_id;"`
	ID             string    `orm:"column:id;size:36"`
	Status         uint32    `orm:"column:status"`
	Type           uint8     `orm:"column:type"`
	UserID         uint      `orm:"column:user_id"`
	PoolTypeID     string    `orm:"column:pool_type_id;size:36"`
	FriendlyID     string    `orm:"column:friendly_id;size:20;null"`
	Name           string    `orm:"column:name;size:50"`
	Class          uint8     `orm:"column:class"`
	Abbreviation   string    `orm:"column:abbreviation;size:10"`
	MinerCount     uint      `orm:"column:miner_count"`
	MinerAmount    float64   `orm:"column:miner_amount;"`
	AmountCoin     uint8     `orm:"column:amount_coin;"`
	MinerCoin      float64   `orm:"column:miner_coin;"`
	BasicYield     float64   `orm:"column:basic_yield;"`
	PresentedYield float64   `orm:"column:presented_yield;"`
	AnnualYield    float64   `orm:"column:annual_yield;"`
	Duration       uint      `orm:"column:duration;"`
	Sold           uint      `orm:"column:sold;"`
	Expires        uint64    `orm:"column:expires;"`
	ExitAt         uint64    `orm:"column:exit_at;null"`
	MiningDate     time.Time `orm:"column:mining_date;type:date;"`
	MiningLeft     uint8     `orm:"column:mining_left;"`
	MiningCount    uint8     `orm:"column:mining_count;"`
	CreatedAt      time.Time `orm:"column:created_at;type:datetime;"`
	UpdatedAt      time.Time `orm:"column:updated_at;type:datetime;"`
}

// TableName t_city
func (City) TableName() string {
	return "t_city"
}

// TableName t_city
func (Pool) TableName() string {
	return "pool"
}

// FindOne 查找一条记录
func FindPool() mine_intf.PoolItem {
	db, err := gorm.Open("mysql", "root:kekexi00@/cloud_miner?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("连接数据库失败")
	}
	defer db.Close()
	var pool Pool
	db.First(&pool)
	fmt.Println(pool)
	var poolItem mine_intf.PoolItem
	poolItem.AutoId = pool.AutoID
	poolItem.Id = pool.ID
	poolItem.Status = pool.Status
	poolItem.Type = uint32(pool.Type)
	poolItem.UserId = uint32(pool.UserID)
	poolItem.PoolTypeId = pool.PoolTypeID
	poolItem.FriendlyId = pool.FriendlyID
	poolItem.Name = pool.Name
	poolItem.Class = uint32(pool.Class)
	poolItem.Abbreviation = pool.Abbreviation
	poolItem.MinerCount = uint32(pool.MinerCount)
	poolItem.MinerAmount = pool.MinerAmount
	poolItem.AmountCoin = uint32(pool.AmountCoin)
	poolItem.MinerCoin = pool.MinerCoin
	poolItem.BasicYield = pool.BasicYield
	poolItem.PresentedYield = pool.PresentedYield
	poolItem.AnnualYield = pool.AnnualYield
	poolItem.Sold = uint32(pool.Sold)
	poolItem.Expires = pool.Expires
	poolItem.ExitAt = pool.ExitAt
	poolItem.Duration = uint32(pool.Duration)
	var (
		year  int
		month time.Month
		day   int
	)
	year, month, day = pool.MiningDate.Date()
	poolItem.MiningDate = strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(day)
	poolItem.MiningLeft = uint32(pool.MiningLeft)
	poolItem.MiningCount = uint32(pool.MiningCount)
	year, month, day = pool.CreatedAt.Date()
	poolItem.CreatedAt = strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(day)
	year, month, day = pool.UpdatedAt.Date()
	poolItem.UpdatedAt = strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(day)
	return poolItem
}

// FindOne 查找一条记录
func FindOne() City {
	db, err := gorm.Open("mysql", "root:kekexi00@/demo-schema?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("连接数据库失败")
	}
	defer db.Close()

	// // 自动迁移模式
	// db.AutoMigrate(&Product{})

	// // 创建
	// db.Create(&Product{Code: "L1212", Price: 1000})

	// 读取
	var city City
	db.First(&city, 1) // 查询id为1的product
	fmt.Println(city)
	return city
	// row2 := db.First(&city, "name = ?", "涂涂") // 查询code为l1212的product

	// 更新 - 更新product的price为2000
	// db.Model(&product).Update("Price", 2000)

	// 删除 - 删除product
	// db.Delete(&product)
}
