package sysinit

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func initDB() {
	dbAlias := beego.AppConfig.String("db_alias")
	dbName := beego.AppConfig.String("db_name")
	dbUser := beego.AppConfig.String("db_user")
	dbPwd := beego.AppConfig.String("db_pwd")
	dbHost := beego.AppConfig.String("db_host")
	dbPort := beego.AppConfig.String("db_port")
	dbCharset := beego.AppConfig.String("db_charset")

	orm.RegisterDataBase(dbAlias, "mysql", dbUser+":"+dbPwd+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset="+dbCharset, 30)
	fmt.Println(dbUser + ":" + dbPwd + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=" + dbCharset)
	isDev := (beego.AppConfig.String("runmode") == "dev")

	// orm.RunSyncdb("default", false, isDev)
	if isDev {
		orm.Debug = isDev
	}

}
