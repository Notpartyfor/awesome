package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Orders   []Order
}

type Order struct {
	gorm.Model
	UserID uint
	Price  float64
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	dbGorm, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 自动迁移
	//dbGorm.AutoMigrate(&User{},&Order{})

	// 简单创建
	//user := &User{
	//	Username:"二号",
	//}
	//dbGorm.Create(user)

	//order := &Order{
	//	UserID:2,
	//	Price:4000,
	//}
	//dbGorm.Create(order)

	// 一般的preload可以联系起来
	users := make([]*User, 0)
	//dbGorm.Find(&users)
	//dbGorm.Preload("Orders").Find(&users)

	// 自定义条件的Preload
	//dbGorm.Preload("Orders", orderConditions).Find(&users)
	newsIDs := make([]interface{}, 0)
	if err := dbGorm.Raw("SELECT `id` FROM `users` WHERE `username` = ?", "二号").Scan(&newsIDs).Error; err != nil {
		fmt.Println(err)
	}

	for _, v := range users {
		fmt.Println(v.Username)
		fmt.Println(v.Orders)
	}

}

func orderConditions(db *gorm.DB) *gorm.DB {
	return db.
		Order("orders.price DESC").Where("orders.price > ?", "3000")
}
