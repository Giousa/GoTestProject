package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"
	template2 "testProject/04-模板/templates"
	"time"
	"unicode"
)

/**
	TODO 输入相关名称，自动生成java的 Controller Service  Impl Mapper xml等文件 以及对应的方法 加上自定义备注(后期)
 */

type JavaModel struct {
	Package string
	NameMap []string
}

type ClassTemplate struct {
	Package string
	Name string
	NameHumpLower string //驼峰，首字母小写
	NameSeparate string //蛇形，-分隔
	Author string //作者
	DateTime string //时间
}

func (j *JavaModel) Run() {
	packageName := j.Package
	fmt.Println("包名：",packageName)
	path,_ := os.Getwd()
	fmt.Println("根目录：",path)
	//os.Mkdir(packageName,os.FileMode())
	newPackage := strings.Replace(packageName,".","/",-1)
	fmt.Println("包名转文件目录："+newPackage)
	fmt.Println("开始在根目录下创建包名对应文件夹：")
	newPath := path+"/"+newPackage
	os.MkdirAll(newPath,os.ModePerm)

	fmt.Println("开始创建：controller service service/impl文件夹")
	os.MkdirAll(newPath+"/controller",os.ModePerm)
	os.MkdirAll(newPath+"/service/impl",os.ModePerm)

	pathController := newPath+"/controller/"
	pathService := newPath+"/service/"
	pathImpl := newPath+"/service/impl/"

	fmt.Println(pathController)
	fmt.Println(pathService)
	fmt.Println(pathImpl)
	for _,v := range j.NameMap{
		buildControllerFile(pathController,packageName,v)
		buildServiceFile(pathService,packageName,v)
		buildServiceImplFile(pathImpl,packageName,v)
	}

}

func buildControllerFile(path string, packageName string, name string)  {
	className := name+"Controller.java"
	path = path+className

	file, err := os.Create(path)
	if err != nil {
		log.Println(err.Error())
	}
	defer file.Close()
	//全名：包含路径
	fmt.Println("新建Controller文件：",file.Name())

	t := template.Must(template.New("").Parse(template2.TextController))

	classTemplate := ClassTemplate{
		Package: packageName,
		Name: name,
		NameHumpLower: Lcfirst(name),
		NameSeparate: separateString(name),
		DateTime:time.Now().Format("2006-01-02"),
		Author: "不笑猫",
	}

	err = t.Execute(file,classTemplate)
	if err != nil {
		log.Print("execute: ", err)
		return
	}

}

func buildServiceFile(path string, packageName string, name string)  {
	className := name+"Service.java"
	path = path+className

	file, err := os.Create(path)
	if err != nil {
		log.Println(err.Error())
	}
	defer file.Close()
	//全名：包含路径
	fmt.Println("新建Service文件：",file.Name())

	t := template.Must(template.New("").Parse(template2.TextService))

	classTemplate := ClassTemplate{
		Package: packageName,
		Name: name,
		NameHumpLower: Lcfirst(name),
		NameSeparate: separateString(name),
		DateTime:time.Now().Format("2006-01-02"),
		Author: "不笑猫",
	}

	err = t.Execute(file,classTemplate)
	if err != nil {
		log.Print("execute: ", err)
		return
	}

}

func buildServiceImplFile(path string, packageName string, name string)  {
	className := name+"ServiceImpl.java"
	path = path+className

	file, err := os.Create(path)
	if err != nil {
		log.Println(err.Error())
	}
	defer file.Close()
	//全名：包含路径
	fmt.Println("新建Impl文件：",file.Name())

	t := template.Must(template.New("").Parse(template2.TextServiceImpl))

	classTemplate := ClassTemplate{
		Package: packageName,
		Name: name,
		NameHumpLower: Lcfirst(name),
		NameSeparate: separateString(name),
		DateTime:time.Now().Format("2006-01-02"),
		Author: "不笑猫",
	}

	err = t.Execute(file,classTemplate)
	if err != nil {
		log.Print("execute: ", err)
		return
	}

}

/**
	首字母小写
 */
func Lcfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

/**
	驼峰转蛇形：XxYy to xx-yy , XxYY to xx-y-y
 */
func separateString(s string) string {
	data := make([]byte, 0, len(s)*2)
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		// or通过ASCII码进行大小写的转化
		// 65-90（A-Z），97-122（a-z）
		//判断如果字母为大写的A-Z就在前面拼接一个_
		if i > 0 && d >= 'A' && d <= 'Z' {
			data = append(data, '-')
		}
		data = append(data, d)
	}
	//ToLower把大写字母统一转小写
	return strings.ToLower(string(data[:]))
}


//定义命令行参数
var packageParam = flag.String("package", "", "input your package")
var classNameParams = flag.String("className", "", "input your classNameList(用-隔开)")

func main() {

	//解析命令行参数
	flag.Parse()
	//输出命令行参数
	fmt.Printf("【packageParam=%s】【classNameParams=%s】",*packageParam,*classNameParams)

	nameMap := make([]string,0)
	paramList := strings.Split(*classNameParams,"-")
	for _,v := range paramList{
		nameMap = append(nameMap, v)
	}

	javaModel := JavaModel{
		Package: *packageParam,
		NameMap: nameMap,
	}

	if len(javaModel.NameMap) == 0 || javaModel.Package == ""{
		fmt.Println("package或className缺失")
		return
	}

	javaModel.Run()



	//nameMap := make([]string,0)
	//nameMap = append(nameMap, "Login")
	//nameMap = append(nameMap, "Register")
	//nameMap = append(nameMap, "EducInfo")
	//
	//javaModel := JavaModel{
	//	Package: "com.giousa.wx",
	//	NameMap: nameMap,
	//}
	//
	//javaModel.Run()

	
}
