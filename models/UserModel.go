package models

import "github.com/astaxie/beego/orm"



type UserModel struct {
	UserId int `orm:"pk;auto"`
	UserKey string `orm:"size(64);unique"`
	UserName string `orm:"size(64)"`
	AuthStr string `orm:"size(512)"`
	PassWord string `orm:"size(128)"`
	IsAdmin int8 `orm:"default(0)"`
}

func (m *UserModel) TableName() string {
	return "bee_user"
}

func UserStruct() []*UserModel {
	query := orm.NewOrm().QueryTable("bee_user")
	data := make([]*UserModel, 0)
	query.OrderBy("-user_id").All(&data)
	return data
}



func UserList(pageSize, page int) ([]*UserModel, int64) {
	query := orm.NewOrm().QueryTable("bee_user")
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
