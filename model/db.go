package model

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"os"
)

const (
	SQL_NUM     = 150 //SQL批处理条数
	MAX_CLIENT  = 400 //最大链接个数
	INIT_CLIENT = 10  //初始化链接个数
)

var (
	MDB        *xorm.Engine //数据库
	analysisDB *xorm.Engine //数据库
	ShopDB     *xorm.Engine //数据库

)

func init() {
	//====================================================================

	url := "root:root@admin_251@tcp(10.0.0.251:3306)/"
	MDB, _ = xorm.NewEngine("mysql", url+"ejavashop_m?charset=utf8")
	ShopDB, _ = xorm.NewEngine("mysql", url+"ejavashop?charset=utf8")
	analysisDB, _ = xorm.NewEngine("mysql", url+"ejavashop_analysis?charset=utf8")
	if os.Getenv("GO_DEV") == "1" {
		MDB.ShowSQL = true
		ShopDB.ShowSQL = true
		analysisDB.ShowSQL = true

	}
	analysisDB.SetMaxIdleConns(INIT_CLIENT)
	analysisDB.SetMaxOpenConns(MAX_CLIENT)
	ShopDB.SetMaxIdleConns(INIT_CLIENT)
	ShopDB.SetMaxOpenConns(MAX_CLIENT)
	MDB.SetMaxIdleConns(INIT_CLIENT)
	MDB.SetMaxOpenConns(MAX_CLIENT)
	//====================================================================

}
func NoData(b bool) error {

	if b {
		return nil
	}
	return errors.New("empty")
}

//错误消息定义
func NoDataMsg(b bool, msg string) error {
	if b {
		return nil
	}
	return errors.New(msg)
}
