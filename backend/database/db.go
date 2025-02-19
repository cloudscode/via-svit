package database

import (
	"fmt"
	"github.com/cloudscode/via-svit/backend/models"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var DB *xorm.Engine

func InitDB() {
	var err error
	// 根据实际情况修改 dsn 参数
	dsn := "root:password@tcp(127.0.0.1:3306)/via_svit_db?charset=utf8mb4&parseTime=True"
	DB, err = xorm.NewEngine("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("数据库连接失败: %v", err))
	}
	// 自动同步订单表结构
	if err = DB.Sync2(new(models.Order)); err != nil {
		panic(fmt.Sprintf("表结构同步失败: %v", err))
	}
}
