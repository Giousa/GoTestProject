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

func InitMysqlEngin(ip string,root string,password string,databaseName string) *xorm.Engine  {

	//数据库引擎

	engine, err := xorm.NewEngine("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8",root,password,ip,databaseName))

	if err != nil {
		fmt.Println("mysql数据库连接失败,err = ",err)
		panic(err.Error())
	}

	//设置显示sql语句
	engine.ShowSQL(true)
	engine.SetMaxOpenConns(10)

	return engine

}

