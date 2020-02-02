package pool

import "github.com/jinzhu/gorm"

var (
	poolDB *gorm.DB
)

func InitDB(db *gorm.DB) {
	poolDB = db
}