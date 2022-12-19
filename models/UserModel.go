package models

import "github.com/astaxie/beego/orm"

type UserModel struct {
	UserId   int
	UserKey  string
	UserName string
	AuthStr  string
	Password string
	IsAdmin  int8
}

func (m *UserModel) TableName() string {
	return "user"
}

func UserStruct() []*UserModel {
	query := orm.NewOrm().QueryTable("user")
	data := make([]*UserModel, 0)
	query.OrderBy("-user_id").All(&data)
	return data
}

func UserList(pageSize, page int) ([]*UserModel, int64) {
	query := orm.NewOrm().QueryTable("user")
	total, _ := query.Count()
	offset := (page - 1) * pageSize
	data := make([]*UserModel, 0)
	query.OrderBy("-user_id").Limit(pageSize, offset).All(&data)
	return data, total
}

func GetUserByName(userkey string) UserModel {
	o := orm.NewOrm()
	user := UserModel{UserKey: userkey}
	o.Read(&user, "user_key")
	return user
}
