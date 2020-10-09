package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ClassifyCategory struct {
	gorm.Model
	Name                 string             `gorm:"type:varchar(60);not null;"`
	Description          string             `gorm:"type:varchar(600);not null;"`
	ParentID             uint               `gorm:"default:0;"`
	Disable              bool               `gorm:"default:false"`
	SubEquipmentCategory []ClassifyCategory `gorm:"foreignkey:ParentID"`
}

type Test struct {
	gorm.Model
	Name string `gorm:"column:name"`
	CCID uint   `gorm:"column:ccid"`
	CID  uint   `gorm:"column:cid"`
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	dbGorm, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	dsn = "root:123456@tcp(127.0.0.1:3306)/asm_collaborative?charset=utf8mb4&parseTime=True&loc=Local"
	dbColla, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 简单创建
	//test := &Test{
	//	Name:  "就这",
	//	CCID:  4,
	//	CID:   5,
	//}
	//
	//dbGorm.Create(test)

	KeyWord := fmt.Sprintf("%s%s%s", "%", "类", "%")
	test := &Test{}
	if err := dbGorm.First(&test).Error; err != nil {
		fmt.Println(err)
	}
	fmt.Println(test)

	classIds := make([]*ClassifyCategory, 0)
	if err := dbColla.Select("id", "parent_id").Where("name LIKE ?", KeyWord).Model(ClassifyCategory{}).Find(&classIds).Error; err != nil {
		fmt.Println(err)
	}

	fmt.Println(ProjectCategoryFilter(classIds))

	any := make([]*interface{}, 0)
	joinParams := "tests INNER JOIN classify_categories ON tests.cid = classify_categories.id"
	if err := dbColla.Where("classify_categories.name LIKE ?", KeyWord).Joins(joinParams).Find(&any).Error; err != nil {
		fmt.Println(err)
	}

}

func ProjectCategoryFilter(categories []*ClassifyCategory) []uint {
	classIds := make([]*ClassifyCategory, 0)
	dsn := "root:123456@tcp(127.0.0.1:3306)/asm_collaborative?charset=utf8mb4&parseTime=True&loc=Local"
	dbColla, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if err := dbColla.Model(ClassifyCategory{}).Find(&classIds).Error; err != nil {
		fmt.Println(err)
	}

	m := make(map[uint]uint, 0)
	for _, c := range classIds {
		m[c.ID] = c.ParentID
	}

	Ids := make([]uint, 0)

	for _, c := range categories {
		if c.ID == 4 || c.ID == 5 {
			Ids = append(Ids, c.ID)
			continue
		}
		if c.ParentID == 4 || c.ParentID == 5 {
			Ids = append(Ids, c.ID)
			continue
		}
		if m[c.ParentID] == 4 || m[c.ParentID] == 5 {
			Ids = append(Ids, c.ID)
			continue
		}

	}
	return Ids
}
