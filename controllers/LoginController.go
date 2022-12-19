package controllers

import (
	"Beego-demo/models"
	"fmt"
	"github.com/astaxie/beego"
	"strings"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Index() {
	method := c.Ctx.Request.Method
	fmt.Println(method)
	if c.Ctx.Request.Method == "POST" {
		userkey := strings.TrimSpace(c.GetString("userkey"))
		password := strings.TrimSpace(c.GetString("password"))

		if len(userkey) > 0 && len(password) > 0 {
			user := models.GetUserByName(userkey)
			if password == user.Password {
				c.SetSession("user", user)
				c.Redirect("/", 302)
				c.StopRun()
			}
		}
	}
	c.TplName = "login/index.html"
}

func (c *LoginController) Logout() {
	method := c.Ctx.Request.Method
	fmt.Println(method)
	c.DelSession("user")
	c.Redirect("/", 302)
}
