/**
 *@Desc:
 *@Author:Giousa
 *@Date:2020/6/30
 */
package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" //不能忘记导入
	"github.com/go-xorm/xorm"
)

func InitMysqlEngin() *xorm.Engine  {

	//数据库引擎
	engine, err := xorm.NewEngine("mysql", "root:h5s/X_7FLkzj@tcp(47.103.115.252:3306)/das?charset=utf8")

	if err != nil {
		fmt.Println("mysql数据库连接失败,err = ",err)
		panic(err.Error())
	}

	//设置显示sql语句
	engine.ShowSQL(true)
	engine.SetMaxOpenConns(10)

	return engine

}

