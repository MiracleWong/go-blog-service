package model

import (
	"fmt"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"

	"github.com/MiracleWong/go-blog-service/pkg/setting"
)

type Model struct {
	// 创建人
	CreatedBy string `json:"created_by"`
	// 创建时间
	CreatedOn uint32 `json:"created_on"`
	// 删除时间
	DeletedOn uint32 `json:"deleted_on"`
	// id
	ID uint32 `json:"id"`
	// 是否删除，0为未删除，1为删除
	IsDel uint8 `json:"is_del"`
	// 修改人
	ModifiedBy string `json:"modified_by"`
	// 修改时间
	ModifiedOn uint32 `json:"modified_on"`
}

func NewDBEngine(databaseSettingS *setting.DatabaseSettingS) (*gorm.DB, error) {
	//s := fmt.Printf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local")
	dsn := "root:jxkj@123@tcp(127.0.0.1:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	fmt.Println("数据库连接成功: ", db)
	return db, err
}
