package models

import "github.com/astaxie/beego/orm"

func init() {
	orm.RegisterModel(new(MenuModel))
	orm.RegisterModel(new(UserModel))
	//orm.RegisterModel(new(DataModel))
}
