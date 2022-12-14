package controllers

import (
	"Beego-demo/consts"
	"Beego-demo/models"
)

type MenuController struct {
	BaseController
}

func (c *MenuController) Index() {
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "menu/footerjs.html"
	c.setTpl("menu/index.html")
}

func (c *MenuController) List() {
	data, total := models.MenuList()
	type MenuEx struct {
		models.MenuModel
		ParentName string
	}
	var menu = make(map[int]string)
	for _, v := range data {
		menu[v.Mid] = v.Name
	}
	var dataEx []MenuEx
	for _, v := range data {
		dataEx = append(dataEx, MenuEx{*v, menu[v.Parent]})
	}
	c.listJsonResult(consts.JRCodeSucc, "OK", total, dataEx)
}

func (c *MenuController) Add() {}

func (c *MenuController) AddDo() {}

func (c *MenuController) Edit() {}

func (c *MenuController) EditDo() {}
