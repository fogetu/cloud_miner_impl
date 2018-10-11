package poolmodel

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// City 城市
type City struct {
	ID      uint   // 列名为 `id`
	Country string // 列名为 `country`
	Map     string // 列名为 `map`
	Name    string // 列名为 `name`
	State   string // 列名为 `state`
}

// TableName t_city
func (City) TableName() string {
	return "t_city"
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
