package controllers

import (
	"Beego-demo/consts"
	"Beego-demo/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"strings"
)

type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
	IsLogin        bool
	Loginuser      interface{}
}

func (c *BaseController) Prepare() {
	c.controllerName, c.actionName = c.GetControllerAndAction()
	beego.Informational(c.controllerName, c.actionName)
	fmt.Println("beego:perpare" + c.controllerName + "," + c.actionName)
	//user := 0

	c.Data["Menu"] = models.MenuTreeStruct()
}

func (c *BaseController) setTpl(template ...string) {
	var tplName string
	layout := "common/layout.html"
	switch {
	case len(template) == 1:
		tplName = template[0]
	case len(template) == 2:
		tplName = template[0]
		layout = template[1]
	default:
		ctrlName := strings.ToLower(c.controllerName[0 : len(c.controllerName)-10])
		actionName := strings.ToLower(c.actionName)
		tplName = ctrlName + "/" + actionName + ".html"
	}

	_, found := c.Data["Footer"]
	if !found {
		c.Data["Footer"] = "menu/footerjs.html"
	}
	c.Layout = layout
	c.TplName = tplName
}

func (c *BaseController) jsonResult(code consts.JsonResultCode, msg string, obj interface{}) {
	r := &models.JsonResult{code, msg, obj}
	c.Data["json"] = r
	c.ServeJSON()
	c.StopRun()
}

func (c *BaseController) listJsonResult(code consts.JsonResultCode, msg string, count int64, obj interface{}) {
	r := &models.ListJsonResult{code, msg, count, obj}
	c.Data["json"] = r
	c.ServeJSON()
	c.StopRun()
}

func (c *BaseController) auth() models.UserModel {
	user := c.GetSession("user")
	fmt.Println("user", user)
	if user == nil {
		c.Redirect("/login", 302)
		c.StopRun()
		return models.UserModel{}
	}else{
		fmt.Println("get user:", user.(models.UserModel))
		return user.(models.UserModel)
	}
}

func (c *BaseController) checkToken() interface{} {
	oauth_token := c.GetSession("oauth_token")
	fmt.Println("oauth_token", oauth_token)
	if oauth_token != nil {
		req, err := http.NewRequest(http.MethodGet, "http://localhost:9096/test", nil)
		req.Header.Set("Content-Type", "application/x-www-urlencoded")
		req.Header.Set("Authorizetion", "Bearer "+oauth_token.(string))
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return ""
		}
		defer resp.Body.Close()
		bs, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			return ""
		}
		fmt.Println("resp str:", string(bs))
		m := make(map[string]interface{})
		err = json.Unmarshal(bs, &m)
		if err != nil {
			fmt.Println("Umarshal failed:", err)
			c.Redirect("/login", 302)
			c.StopRun()
		}
		fmt.Println("map:", m)
		return m
	}else{
		c.Redirect("/login", 302)
		c.StopRun()
		return ""
	}
}
