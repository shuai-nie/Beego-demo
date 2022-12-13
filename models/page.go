package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Page struct {
	Id      int
	Website string
	Email   string
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root@135246@tcp(127.0.0.1:3306)")
	orm.RegisterModel(new(Page))
}

func GetPage() Page {

	o := orm.NewOrm()
	p := Page{Id: 1}
	err := o.Read(&p)
	if err != nil {
		fmt.Println(err)
	}
	return p
}

func UpdatePage() {
	p := Page{Id: 1, Email: "my@qq.com"}
	o := orm.NewOrm()
	o.Update(&p, "Email")
}
