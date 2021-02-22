package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Model struct {
	// id
	Id     uint32	`gorm:"primary_key" json:"id"`
	// 创建时间
	CreatedOn      uint32   `json:"created_on"`
	// 创建人
	CreatedBy      string  `json:"created_by"`
	// 修改时间
	ModifiedOn     uint32   `json:"modified_on"`
	// 修改人
	ModifiedBy     string  `json:"modified_by"`
	// 删除时间
	DeletedOn      uint32   `json:"deleted_on"`
	// 是否删除，0为未删除，1为删除
	IsDel  uint8    `json:"is_del"`
}

//func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
func NewDBEngine() (*gorm.DB, error) {
	//s := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local")
	//db, err := gorm.Open(databaseSetting.DBType,s,
	//	databaseSetting.Username,
	//	databaseSetting.Password,
	//	databaseSetting.Host,
	//	databaseSetting.DBName,
	//	databaseSetting.Charset,
	//	databaseSetting.ParseTime,
	//	)
	//DBType: mysql
	//Username: root
	//Password: 123456
	//Host: 127.0.0.1:3006
	//DBName: blog-service
	//TablePrefix: blog_
	//Charset: utf8
	//ParseTime: True
	//MaxIdleConns: 10
	//MaxOpenConns: 30
	db, err := gorm.Open("mysql",fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		"root",
		"123456",
		"127.0.0.1:3306",
		"blog_service",
		"utf8",
		true,
	))
	if err != nil {
		return nil,err
	}
	//if global.ServerSetting.RunMode == "debug" {
	//	db.LogMode(true)
	//}
	db.DB().SetMaxIdleConns(30)
	db.DB().SetMaxOpenConns(30)
	return db, err
}