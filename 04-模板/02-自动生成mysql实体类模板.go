package main

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"io/ioutil"
	"os"
	"sort"
	"testProject/04-模板/db"
)

var engin *xorm.Engine

type MysqlModel struct {
	IP string
	Root string
	Password string
	DabaseName string
}

type TableModel struct {
	ColumnArray []ColumnModel
}

type ColumnModel struct {
	TableName string
	ColumnMap map[int]map[string]string
}

func (mysql *MysqlModel) initMysql(){
	engin = db.InitMysqlEngin(mysql.IP,mysql.Root,mysql.Password,mysql.DabaseName)
}

func findMysqlTableInfo()  {
	table,_ := engin.DBMetas()

	var tableModel TableModel

	for _,v := range table{
		tableName := v.Name
		fmt.Println("------------------------------")
		fmt.Println("表名称：",tableName)

		var columnModel ColumnModel
		columnModel.ColumnMap = make(map[int]map[string]string)


		column := v.Columns
		for index,i := range column(){
			columnName := i.Name
			columnType := i.SQLType.Name
			fmt.Println("字段名称：",columnName)
			fmt.Println("字段类型：",columnType)
			fmt.Println("------------------------------")

			columnModel.TableName = tableName
			//columnModel.ColumnMap[columnName] = columnType
			m := make(map[string]string)
			m[columnName] = columnType
			columnModel.ColumnMap[index] = m
		}

		//tableModel.ColumnArray[index] = columnModel
		tableModel.ColumnArray = append(tableModel.ColumnArray,columnModel)
	}

	buildTableTemplate(tableModel)
}


func buildTableTemplate(tableModel TableModel)  {
	//fmt.Println(tableModel)
	//t := template.Must(template.New("").Parse(templates.TextMysqlTemplate))
	//
	//t.Execute(os.Stdout, tableModel.ColumnArray)

	//权限：os.O_RDWR 可读可写 | os.O_APPEND 尾部追加 | os.O_CREATE 文件不存在时创建 | os.O_TRUNC 打开时，清空文件
	//os.ModePerm：覆盖所有Unix权限位（用于通过&获取类型位）
	filePath := "./mysql2Struct.txt"
	f, err := os.OpenFile(filePath,os.O_RDWR  | os.O_CREATE | os.O_TRUNC,os.ModePerm)
	if err != nil {
		fmt.Println("create file: ", err)
		return
	}

	defer f.Close()

	for _,t := range tableModel.ColumnArray{
		tableName := t.TableName
		f.WriteString("// "+tableName+"...\n")
		f.WriteString("type "+camelString(tableName)+" struct {\n")

		var keys []int
		for k,_ := range t.ColumnMap{
			keys = append(keys, k)

			//f.WriteString(fmt.Sprint("    "+camelString(k)+" "+mysqlType2GoType(v)+" `json:"+k+"`\n"))
		}

		sort.Ints(keys)
		for _,kInt := range keys{
			m := t.ColumnMap[kInt]
			for k,v := range m{
				f.WriteString(fmt.Sprint("    "+camelString(k)+" "+mysqlType2GoType(v)+" `json:"+k+"`\n"))
			}

		}

		f.WriteString("}\n")
		f.WriteString("\n")
		f.WriteString("\n")
		f.WriteString("\n")
	}

	//TODO  调用方法，可以返回
	buf,_ := ioutil.ReadFile(filePath)
	fmt.Println("：：：：：：：：：：：：：")
	fmt.Println("：：：：：：：：：：：：：")
	fmt.Println("：：：：：：：：：：：：：")
	fmt.Println(string(buf))


}

func mysqlType2GoType(s string) string {
	newStr := "string"
	switch s {
		case "INT":
			newStr = "int"
			break

		case "VARCHAR":
			newStr = "string"
			break

		case "DATETIME":
			newStr = "time.Time"
			break
	}

	return newStr
}

/**
	蛇形转驼峰，首字母大写
 */
func camelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}


func main() {

	mysqlModel := MysqlModel{
		IP: "47.103.115.252",
		Root: "root",
		Password: "h5s/X_7FLkzj",
		DabaseName: "das",
	}
	mysqlModel.initMysql()

	findMysqlTableInfo()
}
