package controllers

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
}

func (c *BaseController) Prepare() {
	c.controllerName, c.actionName = c.GetControllerAndAction()
	beego.Informational(c.controllerName, c.actionName)
}

func (c *BaseController) setTol(template ...string) {

}
