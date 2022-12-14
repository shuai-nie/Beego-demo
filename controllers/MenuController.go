package controllers

type MenuController struct {
	BaseController
}

func (c *MenuController) Index() {
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "menu/footerjs.html"
	c.setTol("menu/index.html")
}

func (c *MenuController) List() {}

func (c *MenuController) Add() {}

func (c *MenuController) AddDo() {}

func (c *MenuController) Edit() {}

func (c *MenuController) EditDo() {}
